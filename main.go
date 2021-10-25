package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/trungnghia250/shoes_store/db"
	"github.com/trungnghia250/shoes_store/router"
	"log"
)

func main() {
	if err := db.ConnectDB(); err != nil {
		log.Fatal(err)
	}
	defer db.DB.Client.Disconnect(context.Background())

	app := fiber.New()
	router.Create(app)
	_ = app.Listen(":4000")
}
