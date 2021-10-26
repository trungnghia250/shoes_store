package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trungnghia250/shoes_store/model"
	"github.com/trungnghia250/shoes_store/service"
)

func GetCustomerByID(c *fiber.Ctx) error {
	req := new(model.GetCusRequest)
	if err := c.QueryParser(req); err != nil {
		return err
	}
	customer, err := service.GetCustomerByID(c, req.ID)
	if err != nil {
		return err
	}
	return c.JSON(customer)
}

func CreateCustomer(c *fiber.Ctx) error {
	req := new(model.Customer)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	customer, err := service.CreateCustomer(c, req)
	if err != nil {
		return err
	}
	return c.JSON(customer)
}

func UpdateCustomerInfo(c *fiber.Ctx) error {
	req := new(model.GetCusRequest)
	if err := c.QueryParser(req); err != nil {
		return err
	}
	body := new(model.Customer)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	err := service.UpdateCustomerByID(c, req.ID, body)
	if err!= nil {
		return err
	}

	customer, err := service.GetCustomerByID(c, req.ID)
	if err != nil {
		return err
	}
	return c.JSON(customer)
}