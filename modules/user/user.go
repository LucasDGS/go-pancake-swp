package user

import (
	"errors"
	"os"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Login struct {
	Email    string `json:"email" validate:"required,min=6,max=64"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type User struct {
	gorm.Model `diff:"-" swaggerignore:"true"`
	Email      string `json:"email" validate:"required,min=6,max=64"`
	Password   string `json:"password" validate:"required,min=6,max=64"`
}

func (u *User) ValidateUser() error {

	if u.Email == "" {
		return errors.New("email cannot be empty")
	}

	if !isValidEmail(u.Email) {
		return errors.New("email is not valid")
	}

	if u.Password == "" {
		return errors.New("password cannot be empty")
	}

	return nil
}

// isValidEmail validates an email address using a regex pattern
func isValidEmail(email string) bool {
	regexPattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`
	matched, _ := regexp.MatchString(regexPattern, email)
	return matched
}

func (u *User) GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["email"] = u.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
