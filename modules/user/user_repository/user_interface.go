package user_repository

import (
	"github.com/LucasDGS/go-pancake-swp/modules/user/user_models"
)

type IUserRepository interface {
	CreateUser(user *user_models.User) (*user_models.User, error)
	GetUser(user *user_models.User) (*user_models.User, error)
	UpdateUser(user *user_models.User) (*user_models.User, error)
	DeleteUser(id int) error
}
