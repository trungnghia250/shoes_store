package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/matcornic/hermes/v2"
	"github.com/pkg/errors"
	"github.com/trungnghia250/shoes_store/db"
	"github.com/trungnghia250/shoes_store/mail"
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

	user, _ := GetCustomerByID(c, data.UserId)
	opts := []mail.OptMessage{
		mail.WithMessageHTML("Amazing Shoes Confirmed Order",
			mail.WithTable(hermes.Table{
				Data: [][]hermes.Entry{
					[]hermes.Entry{
						hermes.Entry{
							Key:   "Customer Name",
							Value: user.Name,
						},
						hermes.Entry{
							Key:   "Item",
							Value: data.Items[0].Name,
						},
						hermes.Entry{
							Key:   "Information",
							Value: "Description",
						},
						hermes.Entry{
							Key:   "Payment Method",
							Value: data.PaymentMethod,
						},
						hermes.Entry{
							Key:   "Time Order",
							Value: data.OrderDate,
						},
						hermes.Entry{
							Key:   "Total",
							Value: fmt.Sprintf("%d VND", data.Total),
						},
					},
				},
			}),
			mail.WithOuttros([]string{"Need help, or have questions? Just reply to this email, we'd love to help."}),
			mail.WithIntros([]string{"Thanks for choosing Amazing Store! We have received your order information. All about information that:"})),

		mail.WithTo([]string{user.Email})}
	err = mail.Send(opts...)
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

func Schedule(c *fiber.Ctx, data *model.Schedule) error {
	id, err := getSequenceNextValue(c, "schedule_id")
	if err != nil {
		return err
	}

	data.Id = id
	_, err = db.DB.Schedule.InsertOne(c.Context(), data)
	if err != nil {
		return err
	}

	opts := []mail.OptMessage{
		mail.WithMessageHTML("Amazing Shoes Confirmed Email",
			mail.WithTable(hermes.Table{
				Data: [][]hermes.Entry{
					[]hermes.Entry{
						hermes.Entry{
							Key:   "Customer name",
							Value: data.Name,
						},
						hermes.Entry{
							Key:   "Email",
							Value: data.Email,
						},
						hermes.Entry{
							Key:   "Address",
							Value: data.Address,
						},
						hermes.Entry{
							Key:   "Pack",
							Value: data.Pack,
						},
						hermes.Entry{
							Key:   "Number of Pairs",
							Value: fmt.Sprintf("%d pairs", data.NumOfPair),
						},
						hermes.Entry{
							Key:   "Time Send",
							Value: data.SendTime,
						},
						hermes.Entry{
							Key:   "Time want to received",
							Value: data.ReceivedTime,
						},
						hermes.Entry{
							Key:   "Total",
							Value: fmt.Sprintf("%d VND", data.Total),
						},
					},
				},
			}),
			mail.WithOuttros([]string{"Need help, or have questions? Just reply to this email, we'd love to help."}),
			mail.WithIntros([]string{"Thanks for choosing Amazing Store! We have received your booking cleaning shoes service. All about information that:"})),

		mail.WithTo([]string{data.Email})}
	err = mail.Send(opts...)
	if err != nil {
		return err
	}
	return nil
}
