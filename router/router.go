package router

import (
	"log"

	"github.com/LucasDGS/go-pancake-swp/middlewares"
	"github.com/LucasDGS/go-pancake-swp/modules/user/user_controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Router struct{}

func (r Router) SetupRouter(app *fiber.App) error {
	err := SetupV1Routes(app)
	if err != nil {
		return err
	}

	return nil
}

func SetupV1Routes(app *fiber.App) error {
	log.Println("setting up v1 routes...")

	userController, err := user_controller.NewUserController()
	if err != nil {
		return err
	}

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/v1/login", userController.Login)
	app.Post("/v1/register", userController.CreateUser)

	//users
	api := app.Group("/v1", middlewares.AuthRequired())
	api.Get("/users/:id", userController.GetUser)

	defer log.Println("successfully started v1 routes!")
	return nil
}
