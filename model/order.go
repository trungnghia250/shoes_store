package model

type Order struct {
	ID        int32  `json:"_id,omitempty" bson:"_id,omitempty"`
	State     string `json:"state,omitempty" bson:"state,omitempty"`
	UserId    int32  `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Detail    string `json:"detail,omitempty" bson:"detail,omitempty"`
	Total     int32  `json:"total,omitempty" bson:"total,omitempty"`
	OrderDate string `json:"order_date,omitempty" bson:"order_date,omitempty"'`
}

type GetOrderByUserIDRequest struct {
	ID int32 `json:"id"`
}
