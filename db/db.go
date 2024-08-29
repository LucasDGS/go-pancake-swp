package db

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/LucasDGS/go-pancake-swp/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	l "gorm.io/gorm/logger"
)

type DatabaseCli struct {
	db           *gorm.DB
	readTimeout  time.Duration
	writeTimeout time.Duration
}

var Conn *DatabaseCli

func GetReadTimeout() time.Duration {
	if Conn == nil {
		return 0
	}

	return Conn.readTimeout
}

func GetWriteTimeout() time.Duration {
	if Conn == nil {
		return 0
	}

	return Conn.writeTimeout
}

func Ping() error {
	if Conn == nil || Conn.db == nil {
		return ErrDBNil
	}

	db, err := Conn.db.DB()
	if err != nil {
		return err
	}

	return db.Ping()
}

func Close() error {
	if Conn == nil || Conn.db == nil {
		return ErrDBNil
	}

	db, err := Conn.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func GetDB() (*gorm.DB, error) {
	if Conn == nil || Conn.db == nil {
		return nil, ErrDBNil
	}

	return Conn.db, nil
}

func Connect(disableLogs bool) (*DatabaseCli, error) {

	var err error

	host := utils.GetEnv("DB_HOST", "localhost")
	port := utils.GetEnv("DB_PORT", "5432")
	dbName := utils.GetEnv("DB_NAME", "go-pancake")
	username := utils.GetEnv("DB_USERNAME", "postgres")
	password := utils.GetEnv("DB_PASSWORD", "postgres")

	logMod := l.Info
	if disableLogs {
		logMod = l.Silent
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", host, username, password, dbName, port)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: l.Default.LogMode(logMod),
	})
	if err != nil {
		return nil, err
	}

	dbPostgres, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	maxOpenConn, err := strconv.Atoi(utils.GetEnv("DB_MAX_OPEN_CONNECTIONS", "200"))
	if err != nil {
		return nil, err
	}

	maxIdleConn, err := strconv.Atoi(utils.GetEnv("DB_MAX_IDLE_CONNECTIONS", "10"))
	if err != nil {
		return nil, err
	}

	dbPostgres.SetMaxOpenConns(maxOpenConn)
	dbPostgres.SetMaxIdleConns(maxIdleConn)

	read, write, err := GetDatabaseTimeouts(3, 3)
	if err != nil {
	}

	Conn = &DatabaseCli{
		db:           gormDB,
		readTimeout:  read,
		writeTimeout: write,
	}

	return Conn, nil
}

func GetDatabaseTimeouts(defaultReadTimeout, defaultWriteTimeout int64) (time.Duration, time.Duration, error) {
	readTimeout, ok := os.LookupEnv("DATABASE_READ_TIMEOUT")
	if !ok {
		readTimeout = fmt.Sprintf("%v", defaultReadTimeout)
	}
	parsedReadTimeout, err := strconv.Atoi(readTimeout)
	if err != nil {
		return time.Duration(0), time.Duration(0), err
	}

	writeTimeout, ok := os.LookupEnv("DATABASE_WRITE_TIMEOUT")
	if !ok {
		writeTimeout = fmt.Sprintf("%v", defaultWriteTimeout)
	}

	parsedWriteTimeout, err := strconv.Atoi(writeTimeout)
	if err != nil {
		return time.Duration(0), time.Duration(0), err
	}

	return time.Duration(parsedReadTimeout) * time.Second, time.Duration(parsedWriteTimeout) * time.Second, nil
}
