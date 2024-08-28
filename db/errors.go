package db

import (
	"errors"
)

var (
	ErrDBNil = errors.New("the database is not instantiated")
)

var (
	StrGetDBFail = "database: failed to get a new database connection"
)
