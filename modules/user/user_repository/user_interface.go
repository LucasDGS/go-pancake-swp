package user_repository

import (
	"github.com/LucasDGS/go-pancake-swp/modules/user/user_models"
	"github.com/LucasDGS/go-pancake-swp/utils"
)

type IUserRepository interface {
	CreateUser(user *user_models.CreateUser) (*user_models.CreateUser, error)
	GetUser(user *user_models.User) (*user_models.User, error)
	UpdateUser(user *user_models.User) (*user_models.User, error)
	DeleteUser(id int) error
	ListUsers(user *user_models.User, page, pageSize int32) (*utils.Pagination, error)
}
