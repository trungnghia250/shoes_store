package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trungnghia250/shoes_store/model"
	"github.com/trungnghia250/shoes_store/service"
)

func GetComment(c *fiber.Ctx) error {
	req := new(model.GetCommentByProductID)
	if err := c.QueryParser(req); err != nil {
		return err
	}

	orders, err := service.GetComment(c, req.ID)
	if err != nil {
		return err
	}

	return c.JSON(orders)
}

func CreateComment(c *fiber.Ctx) error {
	req := new(model.Comment)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	comment, err := service.CreateComment(c, req)
	if err != nil {
		return err
	}
	return c.JSON(comment)
}
