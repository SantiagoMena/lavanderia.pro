package types

import "time"

type Address struct {
	ID        string     `json:"id" bson:"_id,omitempty" uri:"id"`
	Client    string     `json:"client,omitempty" bson:"client,omitempty"`
	Name      string     `json:"name" bson:"name"`
	Position  Geometry   `json:"position" bson:"position"`
	Address   string     `json:"address" bson:"address"`
	Phone     string     `json:"phone" bson:"phone"`
	Extra     string     `json:"extra" bson:"extra"`
	CreatedAt *time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at"`
}
