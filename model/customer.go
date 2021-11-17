package model

type Customer struct {
	ID           int32  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	Phone        string `json:"phone,omitempty" bson:"phone,omitempty"`
	Dob          string `json:"dob,omitempty" bson:"dob,omitempty"`
	Email        string `json:"email,omitempty" bson:"email,omitempty"`
	Password     string `json:"password,omitempty" bson:"password,omitempty"`
	DeliveryInfo string `json:"delivery_info,omitempty" bson:"delivery_info,omitempty"`
}

type GetCusRequest struct {
	ID int32 `json:"id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ForgetRequest struct {
	Email string `json:"email"`
}

type Schedule struct {
	Id           int32  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	Phone        string `json:"phone,omitempty" bson:"phone,omitempty"`
	Email        string `json:"email,omitempty" bson:"email,omitempty"`
	NumOfPair    int32  `json:"num_of_pair,omitempty" bson:"num_of_pair,omitempty"`
	Pack         string `json:"pack,omitempty" bson:"pack,omitempty"`
	SendTime     string `json:"send_time,omitempty" bson:"send_time,omitempty"`
	ReceivedTime string `json:"received_time,omitempty" bson:"received_time,omitempty"`
	TypeReceived string `json:"type_received,omitempty" bson:"type_received,omitempty"`
	Address      string `json:"address,omitempty" bson:"address,omitempty"`
	Total        int32  `json:"total,omitempty" bson:"total,omitempty"`
}
