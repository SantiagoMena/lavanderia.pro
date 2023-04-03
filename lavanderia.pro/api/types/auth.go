package types

import "time"

type Auth struct {
	ID         string     `json:"id" bson:"_id,omitempty" uri:"id"`
	Email      string     `json:"email,omitempty" bson:"email,omitempty"`
	Password   []byte     `json:"-" bson:"-"`
	FacebookId string     `json:"facebook_id,omitempty" bson:"facebook_id,omitempty"`
	GoogleId   string     `json:"google_id,omitempty" bson:"google_id,omitempty"`
	AppleId    string     `json:"apple_id,omitempty" bson:"apple_id,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" bson:"deleted_at"`
}
