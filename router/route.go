package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trungnghia250/shoes_store/handler"
)

func Create(app fiber.Router) {
	var r fiber.Router

	r = app.Group("/homepage")
	r.Get("/", handler.GetHomePage)

	r = app.Group("/user")
	r.Post("/login", handler.Login)
	r.Post("/forget", handler.ForgetPassword)
	r.Get("/list", handler.ListUsers)

	r = app.Group("/customer")
	r.Post("/", handler.CreateCustomer)
	r.Get("/", handler.GetCustomerByID)
	r.Put("/", handler.UpdateCustomerInfo)
	r.Delete("/", handler.DeleteCustomer)

	r = app.Group("/brand")
	r.Get("/:brand_name", handler.ListBrandProduct)

	r = app.Group("/shoes")
	r.Get("/size/:size", handler.ListProductBySize)
	r.Get("/id/:id", handler.GetProductByID)
	r.Post("/", handler.CreateProduct)
	r.Get("/", handler.ListAllProduct)
	r.Put("/", handler.UpdateProduct)
	r.Delete("/", handler.DeleteProduct)
	r.Post("/rating", handler.Rating)

	r = app.Group("/order")
	r.Get("/", handler.ListOrderByUserID)
	r.Post("/", handler.CreateOrder)
	r.Get("/list", handler.ListAllOrders)
	r.Put("/", handler.UpdateOrder)
	r.Get("/one", handler.GetOrderByID)
	r.Delete("/", handler.DeleteOrder)

	r = app.Post("/schedule", handler.Schedule)
	r = app.Get("/schedule", handler.GetSchedule)

	r = app.Group("/comment")
	r.Put("/", handler.CreateComment)
	r.Get("/", handler.GetComment)

	r = app.Group("/discount")
	r.Put("/", handler.CreateDiscount)
	r.Post("/", handler.UpdateDiscount)
	r.Get("/", handler.GetDiscount)
	r.Get("/list", handler.ListDiscount)
	r.Delete("/", handler.DeleteDiscount)
}
