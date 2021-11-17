package model

type Customer struct {
	ID       int32  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty"`
	Dob      string `json:"dob,omitempty" bson:"dob,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
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
	Email    string `json:"email"`
}