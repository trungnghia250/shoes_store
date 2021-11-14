package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/trungnghia250/shoes_store/db"
	"github.com/trungnghia250/shoes_store/mail"
	"github.com/trungnghia250/shoes_store/model"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
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

func UpdatePasswordByEmail(c *fiber.Ctx, email, password string) error {
	_, err := db.DB.Customer.UpdateOne(c.Context(), bson.M{
		"email": email,
	}, bson.M{
		"$set": bson.M{"password": password},
	})
	if err != nil {
		return err
	}

	return nil
}

func Login(c *fiber.Ctx, req *model.LoginRequest) (*model.Customer, error) {
	cus := new(model.Customer)
	if err := db.DB.Customer.FindOne(c.Context(), bson.M{
		"email":    req.Email,
		"password": req.Password,
	}).Decode(&cus); err != nil {
		return nil, err
	}
	return cus, nil
}

func ForgetPassword(c *fiber.Ctx, req *model.ForgetRequest) error {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	newPass := make([]byte, 10)
	for i := range newPass {
		newPass[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	err := UpdatePasswordByEmail(c, req.Email, string(newPass))
	if err != nil {
		return err
	}
	opts := []mail.OptMessage{
		mail.WithMessageHTML("Amazing Shoes [Resend Password]",
			mail.WithAction([]mail.EmailButton{
				{
					Color:     "#0ba366",
					TextColor: "#ffffff",
					Text:      string(newPass),
				},
			}),
			mail.WithOuttros([]string{"Need help, or have questions? Just reply to this email, we'd love to help."}),
			mail.WithIntros([]string{"Thanks for choosing Amazing Store! Your password has been changed"})),

		mail.WithTo([]string{req.Email})}
	err = mail.Send(opts...)
	if err != nil {
		return err
	}
	return nil
}
