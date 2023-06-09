package types

import "time"

type Business struct {
	ID        string     `json:"id" bson:"_id,omitempty" uri:"id"`
	Name      string     `json:"name"`
	Position  Geometry   `json:"position" bson:"position"`
	Auth      string     `json:"auth,omitempty" bson:"auth,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at"`
}
