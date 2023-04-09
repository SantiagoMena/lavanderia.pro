package types

import "time"

type Product struct {
	ID        string     `json:"id" bson:"_id,omitempty" uri:"id"`
	Name      string     `json:"name"`
	Price     float64    `json:"price"`
	Business  string     `json:"business,omitempty" bson:"business,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at"`
}
