package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/delivery"
)

type DeliveryController struct {
	PostDeliveryHandler *delivery.PostDeliveryHandler
}

func NewDeliveryController(
	PostDeliveryHandler *delivery.PostDeliveryHandler,
) *DeliveryController {
	return &DeliveryController{
		PostDeliveryHandler: PostDeliveryHandler,
	}
}

func (controller DeliveryController) PostDelivery(delivery *types.Delivery) (types.Delivery, error) {
	deliveryDb, err := controller.PostDeliveryHandler.Handle(delivery)

	if err != nil {
		return types.Delivery{}, err
	}

	return deliveryDb, err
}
