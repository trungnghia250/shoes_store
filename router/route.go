package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trungnghia250/shoes_store/handler"
)

func Create(app fiber.Router) {
	var r fiber.Router

	r = app.Group("/homepage")
	r.Get("/", handler.GetHomePage)

	r = app.Group("/customer")
	r.Post("/", handler.CreateCustomer)
	r.Get("/", handler.GetCustomerByID)
}
