package model

type Product struct {
	ID          int32  `json:"_id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Brand       string `json:"brand" bson:"brand"`
	Price       int32  `json:"price" bson:"price"`
	Link        string `json:"link" bson:"link"`
	Gender      string `json:"gender" bson:"gender"`
	Size        int32  `json:"size" bson:"size"`
	Description string `json:"description" bson:"description"`
	Discount    int32  `json:"discount" bson:"discount"`
}
