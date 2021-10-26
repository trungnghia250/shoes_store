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
	r.Put("/", handler.UpdateCustomerInfo)

	r = app.Group("/brand")
	r.Get("/:brand_name", handler.ListBrandProduct)

	r = app.Group("/shoes")
	r.Get("/size/:size", handler.ListProductBySize)
	r.Get("/id/:id", handler.GetProductByID)
	r.Post("/", handler.CreateProduct)

	r = app.Group("/order")
	r.Get("/", handler.ListOrderByUserID)
	r.Post("/", handler.CreateOrder)
}
