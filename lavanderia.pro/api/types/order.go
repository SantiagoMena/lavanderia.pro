package types

import "time"

type Order struct {
	ID                 string     `json:"id" bson:"_id,omitempty" uri:"id"`
	Business           Business   `json:"business" bson:"business,omitempty"`
	Client             Client     `json:"client,omitempty" bson:"client,omitempty"`
	Address            Address    `json:"address", bson:"address"`
	CreatedAt          *time.Time `json:"created_at,omitempty" bson:"created_at"`
	AcceptedAt         *time.Time `json:"created_at,omitempty" bson:"created_at"`
	RejectedAt         *time.Time `json:"created_at,omitempty" bson:"created_at"`
	AssignedPickUpAt   *time.Time `json:"created_at,omitempty" bson:"created_at"`
	PickUpClientAt     *time.Time `json:"created_at,omitempty" bson:"created_at"`
	ProcessingOrderAt  *time.Time `json:"created_at,omitempty" bson:"created_at"`
	FinishedOrderAt    *time.Time `json:"created_at,omitempty" bson:"created_at"`
	AssignedDeliveryAt *time.Time `json:"created_at,omitempty" bson:"created_at"`
	PickUpBusinessAt   *time.Time `json:"created_at,omitempty" bson:"created_at"`
	DeliveryClientAt   *time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty" bson:"deleted_at"`
}
