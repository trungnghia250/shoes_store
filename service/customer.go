package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/trungnghia250/shoes_store/db"
	"github.com/trungnghia250/shoes_store/model"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateCustomer(c *fiber.Ctx, data *model.Customer) (*model.Customer, error) {
	if data == nil {
		return nil, errors.Errorf("data is nil")
	}

	id, err := getSequenceNextValue(c, "customer_id")
	if err != nil {
		return nil, err
	}

	data.ID = id
	_, err = db.DB.Customer.InsertOne(c.Context(), data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetCustomerByID(c *fiber.Ctx, id int32) (*model.Customer, error) {
	cus := new(model.Customer)
	if err := db.DB.Customer.FindOne(c.Context(), bson.M{
		"_id": id,
	}).Decode(&cus); err != nil {
		return nil, err
	}

	return cus, nil
}

func UpdateCustomerByID(c *fiber.Ctx, id int32, data *model.Customer) error {
	_, err := db.DB.Customer.UpdateOne(c.Context(), bson.M{
		"_id": id,
	}, bson.M{
		"$set": data,
	})
	if err != nil {
		return err
	}

	return nil
}