package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/trungnghia250/shoes_store/db"
	"github.com/trungnghia250/shoes_store/router"
	"log"
	"os"
)

func main() {
	if err := db.ConnectDB(); err != nil {
		log.Fatal(err)
	}
	defer db.DB.Client.Disconnect(context.Background())
	if err := db.ConnectAws(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Use(cors.New())
	router.Create(app)
	port := os.Getenv("PORT")

	_ = app.Listen(":"+port)
}
