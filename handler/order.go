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

func ListAllOrders(c *fiber.Ctx) error {
	orders, err := service.ListAllOrders(c)
	if err != nil {
		return err
	}

	return c.JSON(orders)
}

func UpdateOrder(c *fiber.Ctx) error {
	req := new(model.UpdateRequest)
	if err := c.QueryParser(req); err != nil {
		return err
	}

	body := new(model.Order)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	err := service.UpdateOrder(c, req.Id, body)
	if err != nil {
		return err
	}

	return c.JSON(DefaultResponse{
		StatusCode: fiber.StatusOK,
	})
}

func GetOrderByID(c *fiber.Ctx) error {
	req := new(model.UpdateRequest)
	if err := c.QueryParser(req); err != nil {
		return err
	}

	order, err := service.GetOrderByID(c, req.Id)
	if err != nil {
		return err
	}

	return c.JSON(order)
}

func DeleteOrder(c *fiber.Ctx) error {
	req := new(model.GetCusRequest)
	if err := c.QueryParser(req); err != nil {
		return err
	}
	err := service.DeleteOrderByID(c, req.ID)
	if err != nil {
		return err
	}
	return c.JSON(DefaultResponse{StatusCode: fiber.StatusOK})
}
