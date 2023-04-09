package types

import "time"

type Order struct {
	ID                 string         `json:"id" bson:"_id,omitempty" uri:"id"`
	Business           Business       `json:"business" bson:"business,omitempty"`
	Client             Client         `json:"client,omitempty" bson:"client,omitempty"`
	Address            Address        `json:"address" bson:"address"`
	Delivery           Delivery       `json:"delivery" bson:"delivery"`
	Products           []OrderProduct `json:"products" bson:"products"`
	CreatedAt          *time.Time     `json:"created_at,omitempty" bson:"created_at"`
	AcceptedAt         *time.Time     `json:"accepted_at,omitempty" bson:"accepted_at"`
	RejectedAt         *time.Time     `json:"rejected_at,omitempty" bson:"rejected_at"`
	AssignedPickUpAt   *time.Time     `json:"assigned_pickup_at,omitempty" bson:"assigned_pickup_at"`
	PickUpClientAt     *time.Time     `json:"pickup_client_at,omitempty" bson:"pickup_client_at"`
	ProcessingOrderAt  *time.Time     `json:"processing_at,omitempty" bson:"processing_at"`
	FinishedOrderAt    *time.Time     `json:"finished_at,omitempty" bson:"finished_at"`
	AssignedDeliveryAt *time.Time     `json:"assigned_delivery_at,omitempty" bson:"assigned_delivery_at"`
	PickUpBusinessAt   *time.Time     `json:"pickup_business_at,omitempty" bson:"pickup_business_at"`
	DeliveredClientAt  *time.Time     `json:"delivered_client_at,omitempty" bson:"delivered_client_at"`
	DeletedAt          *time.Time     `json:"deleted_at,omitempty" bson:"deleted_at"`
}
