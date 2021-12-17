package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/trungnghia250/shoes_store/db"
	"github.com/trungnghia250/shoes_store/model"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateDiscount(c *fiber.Ctx, data *model.Discount) (*model.Discount, error) {
	if data == nil {
		return nil, errors.Errorf("data is nil")
	}

	id, err := getSequenceNextValue(c, "discount_id")
	if err != nil {
		return nil, err
	}

	data.ID = id
	_, err = db.DB.Discount.InsertOne(c.Context(), data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ListAllDiscounts(c *fiber.Ctx) ([]*model.Discount, error) {
	var discounts []*model.Discount
	cursor, err := db.DB.Discount.Find(c.Context(), bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(c.Context(), &discounts); err != nil {
		return nil, err
	}

	return discounts, nil
}

func UpdateDiscount(c *fiber.Ctx, id int32, data *model.Discount) error {
	_, err := db.DB.Discount.UpdateOne(c.Context(), bson.M{
		"_id": id,
	}, bson.M{
		"$set": data,
	})
	if err != nil {
		return err
	}

	return nil
}
func GetDiscountByName(c *fiber.Ctx, code string) (*model.Discount, error) {
	var discount *model.Discount
	if err := db.DB.Discount.FindOne(c.Context(), bson.M{
		"code": code,
	}).Decode(&discount); err != nil {
		return nil, err
	}
	return discount, nil
}

func DeleteDiscountByID(c *fiber.Ctx, id int32) error {
	_, err := db.DB.Discount.DeleteOne(c.Context(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
