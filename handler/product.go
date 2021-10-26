package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trungnghia250/shoes_store/model"
	"github.com/trungnghia250/shoes_store/service"
	"strconv"
)

func ListBrandProduct(c *fiber.Ctx) error {
	brandName := c.Params("brand_name")

	products, err := service.ListBrandProduct(c, brandName)
	if err != nil {
		return err
	}
	return c.JSON(products)
}

func ListProductBySize(c *fiber.Ctx) error {
	size := c.Params("size")
	sizeInt, _ := strconv.ParseInt(size,10, 32)

	products, err := service.ListProductBySize(c, int32(sizeInt))
	if err != nil {
		return err
	}
	return c.JSON(products)
}

func GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, _ := strconv.ParseInt(id,10, 32)

	product, err := service.GetProductByID(c, int32(idInt))
	if err != nil {
		return err
	}

	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	req := new(model.Product)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	customer, err := service.CreateProduct(c, req)
	if err != nil {
		return err
	}
	return c.JSON(customer)
}