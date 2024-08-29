package user

import (
	"log"

	"github.com/LucasDGS/go-pancake-swp/db"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *User) (*User, error)
	GetUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id int) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) (UserRepository, error) {
	if db == nil {
		log.Println(gorm.ErrInvalidDB.Error())
	}

	return UserRepository{
		db: db,
	}, nil
}

func (ur UserRepository) CreateUser(user *User) (*User, error) {
	result := ur.db.Table("users").Create(user)
	err := db.HandleResult(result)
	if err != nil {
		log.Println(err.Error())
		return &User{}, err
	}

	return user, nil
}

func (ur UserRepository) GetUser(user *User) (*User, error) {
	result := ur.db.Table("users").First(user)
	err := db.HandleResult(result)
	if err != nil {
		log.Println(err.Error())
		return &User{}, err
	}

	return user, nil
}

func (ur UserRepository) UpdateUser(user *User) (*User, error) {
	result := ur.db.Table("users").Updates(user)
	err := db.HandleResult(result)
	if err != nil {
		log.Println(err.Error())
		return &User{}, err
	}

	return user, nil
}

func (cr UserRepository) DeleteUser(userId int) error {
	result := cr.db.Model(&User{}).Delete("id = ?", userId)
	err := db.HandleResult(result)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
