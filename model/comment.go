package model

type Comment struct {
	ID        int32  `json:"_id,omitempty" bson:"_id,omitempty"`
	Owner     string `json:"owner,omitempty" bson:"owner,omitempty"`
	ProductID int32  `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Content   string `json:"content,omitempty" bson:"content,omitempty"`
	CreatAt   string `json:"creat_at,omitempty" bson:"creat_at,omitempty"`
}

type GetCommentRequest struct {
	ProductID int32 `json:"product_id"`
}
