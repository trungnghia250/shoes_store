package model

type Product struct {
	ID          int32  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Brand       string `json:"brand,omitempty" bson:"brand,omitempty"`
	Price       int32  `json:"price,omitempty" bson:"price,omitempty"`
	Link        string `json:"link,omitempty" bson:"link,omitempty"`
	Gender      string `json:"gender,omitempty" bson:"gender,omitempty"`
	Size        int32  `json:"size,omitempty" bson:"size,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Discount    int32  `json:"discount,omitempty" bson:"discount,omitempty"`
}

type UpdateRequest struct {
	Id int32 `json:"id"`
}
