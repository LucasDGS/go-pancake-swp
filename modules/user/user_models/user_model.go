package user_models

import (
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type CreateUser struct {
	gorm.Model         `diff:"-" swaggerignore:"true"`
	FirstName          string `json:"firstName" validate:"required,min=3,max=64"`
	LastName           string `json:"lastName" validate:"required,min=3,max=64"`
	Email              string `json:"email" validate:"required,min=6,max=64"`
	Phone              string `json:"phone" validate:"required,min=6,max=64"`
	Password           string `json:"password" validate:"required,min=6"`
	Address            string `json:"address" validate:"required,min=6"`
	DiscountPercentage int32  `json:"discountPercentage"`
	IsAdmin            bool   `json:"isAdmin"`
}

type Login struct {
	Email    string `json:"email" validate:"required,min=6,max=64"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
type User struct {
	gorm.Model         `diff:"-" swaggerignore:"true"`
	FirstName          string `json:"firstName" validate:"required,min=3,max=64"`
	LastName           string `json:"lastName" validate:"required,min=3,max=64"`
	Email              string `json:"email" validate:"required,min=6,max=64"`
	Phone              string `json:"phone" validate:"required,min=6,max=64"`
	Password           string `json:"password" validate:"required,min=6,max=64"`
	Address            string `json:"address" validate:"required,min=6,max=64"`
	DiscountPercentage int32  `json:"discountPercentage"`
	IsAdmin            bool   `json:"isAdmin"`
}

func (u *CreateUser) ValidateUser() error {
	if u.FirstName == "" {
		return errors.New("first_name cannot be empty")
	}

	if u.LastName == "" {
		return errors.New("last_name cannot be empty")
	}

	if u.Email == "" {
		return errors.New("email cannot be empty")
	}

	if !isValidEmail(u.Email) {
		return errors.New("email is not valid")
	}

	if u.Phone == "" {
		return errors.New("phone cannot be empty")
	}

	if u.Password == "" {
		return errors.New("password cannot be empty")
	}

	if u.DiscountPercentage < 0 {
		return errors.New("discount_percentage cannot be lower than 0")
	}

	return nil
}

// isValidEmail validates an email address using a regex pattern
func isValidEmail(email string) bool {
	regexPattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`
	matched, _ := regexp.MatchString(regexPattern, email)
	return matched
}

const jwtSecret = "your_jwt_secret" // Altere isso para sua chave secreta JWT

func (u *User) GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["email"] = u.Email
	claims["is_admin"] = u.IsAdmin
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
