package types

import "time"

type Laundry struct {
	ID        string     `json:"id" bson:"_id,omitempty"`
	Name      string     `json:"name"`
	Lat       float64    `json:"lat"`
	Long      float64    `json:"long"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
