package main

import (
	"log"

	"github.com/LucasDGS/go-pancake-swp/db"
	_ "github.com/LucasDGS/go-pancake-swp/docs"
	"github.com/LucasDGS/go-pancake-swp/server"
	_ "github.com/joho/godotenv/autoload"
)

// @title                      Go Pancake Swap API
// @version                    1.0
// @description                This is Go Pancake Swap API.
// @BasePath                   /v1
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization

func main() {
	if _, err := db.Connect(true); err != nil {
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
	}

	log.Println("Starting go-pancake-swp app server")

	server := server.NewServer()

	log.Fatalf("fatal error while running app server (%v)", server.Run())

}
