package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/trungnghia250/shoes_store/db"
	"github.com/trungnghia250/shoes_store/model"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func GetComment(c *fiber.Ctx, id int32) ([]*model.Comment, error) {
	var comment []*model.Comment
	cursor, err := db.DB.Comment.Find(c.Context(), bson.M{
		"product_id": id,
	})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(c.Context(), &comment); err != nil {
		return nil, err
	}

	return comment, nil
}

func CreateComment(c *fiber.Ctx, data *model.Comment) (*model.Comment, error) {
	if data == nil {
		return nil, errors.Errorf("data is nil")
	}

	id, err := getSequenceNextValue(c, "comment_id")
	if err != nil {
		return nil, err
	}

	data.ID = id
	data.CreatAt = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC).String()
	_, err = db.DB.Comment.InsertOne(c.Context(), data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
