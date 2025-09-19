// main.go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/misbahkun/test-fullstack-2025/tugas2-fiber-login/config"
	"github.com/misbahkun/test-fullstack-2025/tugas2-fiber-login/handler"
)

func main() {
	redisConnection, err := config.ConnectToRedis()
	if err != nil {
		log.Fatalf("Gagal terhubung ke Redis: %v", err)
	}

	config.SeedDatabaseWithDummyUser(redisConnection)

	fiberApp := fiber.New()

	authenticationHandler := handler.NewAuthHandler(redisConnection)

	apiGroup := fiberApp.Group("/api/v1")
	apiGroup.Post("/login", authenticationHandler.HandleLogin)

	log.Println("Server akan dimulai di http://localhost:3000")
	err = fiberApp.Listen(":3000")
	if err != nil {
		log.Fatalf("Server gagal berjalan: %v", err)
	}
}