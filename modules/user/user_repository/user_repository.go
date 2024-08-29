package user_repository

import (
	"log"

	"github.com/LucasDGS/go-pancake-swp/db"
	"github.com/LucasDGS/go-pancake-swp/modules/user/user_models"
	"gorm.io/gorm"
)

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

func (ur UserRepository) CreateUser(user *user_models.User) (*user_models.User, error) {
	result := ur.db.Table("users").Create(user)
	err := db.HandleResult(result)
	if err != nil {
		log.Println(err.Error())
		return &user_models.User{}, err
	}

	return user, nil
}

func (ur UserRepository) GetUser(user *user_models.User) (*user_models.User, error) {
	result := ur.db.Table("users").First(user)
	err := db.HandleResult(result)
	if err != nil {
		log.Println(err.Error())
		return &user_models.User{}, err
	}

	return user, nil
}

func (ur UserRepository) UpdateUser(user *user_models.User) (*user_models.User, error) {
	result := ur.db.Table("users").Updates(user)
	err := db.HandleResult(result)
	if err != nil {
		log.Println(err.Error())
		return &user_models.User{}, err
	}

	return user, nil
}

func (cr UserRepository) DeleteUser(userId int) error {
	result := cr.db.Model(&user_models.User{}).Delete("id = ?", userId)
	err := db.HandleResult(result)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
