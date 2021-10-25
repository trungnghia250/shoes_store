package model

type Customer struct {
	ID   int32  `json:"_id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

type GetCusRequest struct {
	ID int32 `json:"id"`
}
