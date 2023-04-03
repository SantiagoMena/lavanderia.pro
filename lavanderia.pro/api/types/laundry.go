package types

import "time"

type Business struct {
	ID        string     `json:"id" bson:"_id,omitempty" uri:"id"`
	Name      string     `json:"name"`
	Lat       float64    `json:"lat"`
	Long      float64    `json:"long"`
	CreatedAt *time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at"`
}
