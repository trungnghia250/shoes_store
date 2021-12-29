package model

type Product struct {
	ID            int32   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string  `json:"name,omitempty" bson:"name,omitempty"`
	Brand         string  `json:"brand,omitempty" bson:"brand,omitempty"`
	OriginPrice   int32   `json:"origin_price,omitempty" bson:"origin_price,omitempty"`
	DiscountPrice int32   `json:"discount_price,omitempty" bson:"discount_price,omitempty"`
	Link          string  `json:"link,omitempty" bson:"link,omitempty"`
	Gender        string  `json:"gender,omitempty" bson:"gender,omitempty"`
	Size          int32   `json:"size,omitempty" bson:"size,omitempty"`
	Description   string  `json:"description,omitempty" bson:"description,omitempty"`
	Discount      int32   `json:"discount,omitempty" bson:"discount,omitempty"`
	Color         string  `json:"color,omitempty" bson:"color,omitempty"`
	Quantity      int32   `json:"quantity" bson:"quantity"`
	Rating        *Rating `json:"rating,omitempty" bson:"rating,omitempty"`
}

type Rating struct {
	TotalRating int32   `json:"total_rating,omitempty" bson:"total_rating,omitempty"`
	Value       float64 `json:"value,omitempty" bson:"value,omitempty"`
}

type UpdateRequest struct {
	Id int32 `json:"id"`
}

type RatingRequest struct {
	Value int32 `json:"value"`
}
