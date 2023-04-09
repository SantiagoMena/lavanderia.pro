package types

type OrderProduct struct {
	Product Product `json:"product" bson:"product"`
	Amount  int64   `json:"amount" bson:"amount"`
}
