package model

type Discount struct {
	ID            int32   `json:"_id,omitempty" bson:"_id,omitempty"`
	Code          string  `json:"code,omitempty" bson:"code,omitempty"`
	DiscountValue float64 `json:"discount_value,omitempty" bson:"discount_value,omitempty"`
	Description   string  `json:"description,omitempty" bson:"description,omitempty"`
	Condition     int32   `json:"condition,omitempty" bson:"condition,omitempty"`
	Quantity      int32   `json:"quantity,omitempty" bson:"quantity,omitempty"`
	StartTime     string  `json:"start_time,omitempty" bson:"start_time,omitempty"`
	EndTime       string  `json:"end_time,omitempty" bson:"end_time,omitempty"`
}

type GetDiscountRequest struct {
	Code string `json:"code"`
}
