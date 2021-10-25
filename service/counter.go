package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trungnghia250/shoes_store/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getSequenceNextValue(c *fiber.Ctx, seqName string) (int32, error) {
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	var result bson.M
	if err := db.DB.Counter.FindOneAndUpdate(c.Context(), bson.M{
		"_id": seqName,
	}, bson.M{
		"$inc": bson.M{
			"seq": 1,
		},
	}, &opt).Decode(&result); err != nil {
		return -1, err
	}

	seq := result["seq"].(int32)

	return seq, nil
}
