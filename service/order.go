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

func ListAllOrders(c *fiber.Ctx) ([]*model.Order, error) {
	var orders []*model.Order
	cursor, err := db.DB.Order.Find(c.Context(), bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(c.Context(), &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

func UpdateOrder(c *fiber.Ctx, id int32, data *model.Order) error {
	_, err := db.DB.Order.UpdateOne(c.Context(), bson.M{
		"_id": id,
	}, bson.M{
		"$set": data,
	})
	if err != nil {
		return err
	}

	return nil
}
func GetOrderByID(c *fiber.Ctx, id int32) (*model.Order, error) {
	var order *model.Order
	if err := db.DB.Order.FindOne(c.Context(), bson.M{
		"_id": id,
	}).Decode(&order); err != nil {
		return nil, err
	}
	return order, nil
}

func DeleteOrderByID(c *fiber.Ctx, id int32) error {
	_, err := db.DB.Order.DeleteOne(c.Context(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
