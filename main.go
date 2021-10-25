package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
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

	app := fiber.New()
	router.Create(app)
	port := os.Getenv("PORT")

	_ = app.Listen(":"+port)
}
