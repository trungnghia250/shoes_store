package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/trungnghia250/shoes_store/db"
	"github.com/trungnghia250/shoes_store/model"
	"go.mongodb.org/mongo-driver/bson"
)

func ListOrderByUserID(c *fiber.Ctx, userId int32) ([]*model.Order, error) {
	var orders []*model.Order
	cursor, err := db.DB.Order.Find(c.Context(), bson.M{
		"user_id": userId,
	})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(c.Context(), &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

func CreateOrder(c *fiber.Ctx, data *model.Order) (*model.Order, error) {
	if data == nil {
		return nil, errors.Errorf("data is nil")
	}

	id, err := getSequenceNextValue(c, "order_id")
	if err != nil {
		return nil, err
	}

	data.ID = id
	_, err = db.DB.Order.InsertOne(c.Context(), data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
