package server

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/LucasDGS/go-pancake-swp/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	App    *fiber.App
	Router router.Router
}

func NewServer() Server {
	serverOpts := make([]interface{}, 0)
	app := fiber.New()

	if os.Getenv("LOG_LEVEL") == "info" || os.Getenv("LOG_LEVEL") == "debug" {

		serverOpts = append(serverOpts, logger.New(logger.Config{
			Format:       "[${time}] ${status} - ${latency} -[${ip}]:${port} - ${path} ${queryParams} Request: ${body}  Response:  ${resBody}\n",
			TimeInterval: time.Millisecond,
		}))
	}

	if os.Getenv("ALLOW_CORS") == "true" {
		log.Println("allowing cors...")
		serverOpts = append(serverOpts, cors.New(cors.Config{
			AllowOrigins: "*",
			AllowMethods: strings.Join([]string{
				fiber.MethodGet,
				fiber.MethodPost,
				fiber.MethodHead,
				fiber.MethodPut,
				fiber.MethodDelete,
				fiber.MethodPatch,
			}, ","),
			AllowHeaders: "*"},
		))
	}

	app.Use(serverOpts...)

	return Server{
		App:    app,
		Router: router.Router{},
	}
}

func (s Server) Run() error {
	err := s.Router.SetupRouter(s.App)
	if err != nil {
		return err
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	return s.App.Listen(fmt.Sprintf(":%v", port))
}
