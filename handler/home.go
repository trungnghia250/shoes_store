package handler

import "github.com/gofiber/fiber/v2"

func GetHomePage(c *fiber.Ctx) error {
	return c.JSON("Hello world")
}