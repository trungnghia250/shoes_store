package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trungnghia250/shoes_store/model"
	"github.com/trungnghia250/shoes_store/service"
)

func ListOrderByUserID(c *fiber.Ctx) error {
	req := new(model.GetOrderByUserIDRequest)
	if err := c.QueryParser(req); err != nil {
		return err
	}

	orders, err := service.ListOrderByUserID(c, req.ID)
	if err != nil {
		return err
	}

	return c.JSON(orders)
}


func CreateOrder(c *fiber.Ctx) error {
	req := new(model.Order)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	customer, err := service.CreateOrder(c, req)
	if err != nil {
		return err
	}
	return c.JSON(customer)
}
