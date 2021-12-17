package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trungnghia250/shoes_store/model"
	"github.com/trungnghia250/shoes_store/service"
)

func CreateDiscount(c *fiber.Ctx) error {
	req := new(model.Discount)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	discount, err := service.CreateDiscount(c, req)
	if err != nil {
		return err
	}
	return c.JSON(discount)
}

func ListDiscount(c *fiber.Ctx) error {
	discounts, err := service.ListAllDiscounts(c)
	if err != nil {
		return err
	}

	return c.JSON(discounts)
}

func UpdateDiscount(c *fiber.Ctx) error {
	req := new(model.UpdateRequest)
	if err := c.QueryParser(req); err != nil {
		return err
	}

	body := new(model.Discount)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	err := service.UpdateDiscount(c, req.Id, body)
	if err != nil {
		return err
	}

	return c.JSON(DefaultResponse{
		StatusCode: fiber.StatusOK,
	})
}

func GetDiscount(c *fiber.Ctx) error {
	req := new(model.GetDiscountRequest)
	if err := c.QueryParser(req); err != nil {
		return err
	}

	discount, err := service.GetDiscountByName(c, req.Code)
	if err != nil {
		return err
	}

	return c.JSON(discount)
}

func DeleteDiscount(c *fiber.Ctx) error {
	req := new(model.GetCusRequest)
	if err := c.QueryParser(req); err != nil {
		return err
	}
	err := service.DeleteDiscountByID(c, req.ID)
	if err != nil {
		return err
	}
	return c.JSON(DefaultResponse{StatusCode: fiber.StatusOK})
}
