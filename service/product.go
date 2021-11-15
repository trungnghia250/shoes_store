package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/trungnghia250/shoes_store/db"
	"github.com/trungnghia250/shoes_store/model"
	"go.mongodb.org/mongo-driver/bson"
)

func ListBrandProduct(c *fiber.Ctx, brandName string) ([]*model.Product, error) {
	var products []*model.Product
	cursor, err := db.DB.Product.Find(c.Context(), bson.M{
		"brand": brandName,
	})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(c.Context(), &products); err != nil {
		return nil, err
	}

	return products, nil
}

func ListProductBySize(c *fiber.Ctx, size int32) ([]*model.Product, error) {
	var products []*model.Product
	cursor, err := db.DB.Product.Find(c.Context(), bson.M{
		"size": size,
	})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(c.Context(), &products); err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductByID(c *fiber.Ctx, id int32) (*model.Product, error) {
	var product *model.Product
	if err := db.DB.Product.FindOne(c.Context(), bson.M{
		"_id": id,
	}).Decode(&product); err != nil {
		return nil, err
	}
	return product, nil
}

func CreateProduct(c *fiber.Ctx, data *model.Product) (*model.Product, error) {
	if data == nil {
		return nil, errors.Errorf("data is nil")
	}

	id, err := getSequenceNextValue(c, "product_id")
	if err != nil {
		return nil, err
	}

	data.ID = id
	_, err = db.DB.Product.InsertOne(c.Context(), data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ListAllProduct(c *fiber.Ctx) ([]*model.Product, error) {
	var products []*model.Product
	cursor, err := db.DB.Product.Find(c.Context(), bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(c.Context(), &products); err != nil {
		return nil, err
	}

	return products, nil
}

func UpdateProduct(c *fiber.Ctx, id int32, data *model.Product) error {
	_, err := db.DB.Product.UpdateOne(c.Context(), bson.M{
		"_id": id,
	}, bson.M{
		"$set": data,
	})
	if err != nil {
		return err
	}

	return nil
}

func DeleteProductByID(c *fiber.Ctx, id int32) error {
	_, err := db.DB.Product.DeleteOne(c.Context(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
