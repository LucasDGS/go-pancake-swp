package user_controller

import (
	"strconv"

	"github.com/LucasDGS/go-pancake-swp/db"
	"github.com/LucasDGS/go-pancake-swp/modules/user/user_models"
	"github.com/LucasDGS/go-pancake-swp/modules/user/user_repository"
	"github.com/LucasDGS/go-pancake-swp/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserController struct {
	userRepository user_repository.IUserRepository
}

func NewUserController() (UserController, error) {
	db, err := db.GetDB()
	if err != nil {
		return UserController{}, err
	}

	userRepository, err := user_repository.NewUserRepository(db)
	if err != nil {
		return UserController{}, err
	}

	return UserController{
		userRepository: userRepository,
	}, nil
}

// Login godoc
//
// @Summary Log in a user
// @Tags Users
// @Accept  json
// @Produce json
// @Param   body             body user_models.Login true "Request Body"
// @Success 200 {object}     user_models.LoginResponse
// @Failure 400 {object}     fiber.Map "Invalid request format"
// @Failure 401 {object}     fiber.Map "Unauthorized (Invalid email or password)"
// @Failure 404 {object}     fiber.Map "User not found"
// @Failure default {object} controller_common.SingleErrorMessage "An unexpected error response."
// @Router  /login [post]
func (uc UserController) Login(c *fiber.Ctx) error {
	loginData := user_models.Login{}

	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request format"})
	}

	user, err := uc.userRepository.GetUser(&user_models.User{Email: loginData.Email})
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	if !utils.CheckPasswordHash(loginData.Password, user.Password) {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid password"})
	}

	token, err := user.GenerateJWT()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}

// CreateUser godoc
//
// @Summary Create a new user
// @Tags Users
// @Accept  json
// @Produce json
// @Param   body             body user_models.CreateUser true "Request Body"
// @Success 201 {object}     user_models.User "Created user"
// @Failure 400 {object}     fiber.Map "Invalid request format or validation error"
// @Failure 500 {object}     fiber.Map "Failed to hash password or create user"
// @Failure default {object} controller_common.SingleErrorMessage "An unexpected error response."
// @Router  /register [post]
func (uc UserController) CreateUser(c *fiber.Ctx) error {
	user := &user_models.CreateUser{}

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if err := user.ValidateUser(); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	user.Password = hashedPassword

	createdUser, err := uc.userRepository.CreateUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(201).JSON(createdUser)
}

// GetUser godoc
//
// @Summary Get user by ID
// @Tags Users
// @Accept  json
// @Produce json
// @Param   id     path int true "User ID"
// @Success 200 {object}     user_models.User "User details"
// @Failure 400 {object}     fiber.Map "User ID is required"
// @Failure 404 {object}     fiber.Map "User not found"
// @Failure default {object} controller_common.SingleErrorMessage "An unexpected error response."
// @Router  /users/{id} [get]
// @Security BearerAuth
func (uc UserController) GetUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	if userID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "User ID is required"})
	}

	id, _ := strconv.Atoi(userID)

	foundUser, err := uc.userRepository.GetUser(&user_models.User{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(foundUser)
}

// ListUsers godoc
//
// @Summary List users
// @Tags Users
// @Accept  json
// @Produce json
// @Param   discount_percentage query int false "Filter users by discount percentage"
// @Param   page query int false "Page number for pagination" default(1)
// @Param   limit query int false "Limit of results per page" default(10)
// @Success 200 {object}     utils.Pagination "Paginated list of users"
// @Failure 500 {object}     fiber.Map "Failed to fetch users"
// @Failure default {object} controller_common.SingleErrorMessage "An unexpected error response."
// @Router  /users [get]
// @Security BearerAuth
func (uc UserController) ListUsers(c *fiber.Ctx) error {
	var userFilter user_models.User
	discountPercentage, _ := strconv.Atoi(c.Query("discount_percentage", "0"))
	userFilter.DiscountPercentage = int32(discountPercentage)

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	pagination, err := uc.userRepository.ListUsers(&userFilter, int32(page), int32(limit))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	return c.JSON(pagination)
}
