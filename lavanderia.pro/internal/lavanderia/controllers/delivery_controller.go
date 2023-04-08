package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/delivery"
)

type DeliveryController struct {
	PostDeliveryHandler *delivery.PostDeliveryHandler
	GetDeliveryHandler  *delivery.GetDeliveryHandler
}

func NewDeliveryController(
	PostDeliveryHandler *delivery.PostDeliveryHandler,
	GetDeliveryHandler *delivery.GetDeliveryHandler,
) *DeliveryController {
	return &DeliveryController{
		PostDeliveryHandler: PostDeliveryHandler,
		GetDeliveryHandler:  GetDeliveryHandler,
	}
}

func (controller DeliveryController) PostDelivery(delivery *types.Delivery) (types.Delivery, error) {
	deliveryDb, err := controller.PostDeliveryHandler.Handle(delivery)

	if err != nil {
		return types.Delivery{}, err
	}

	return deliveryDb, err
}

func (controller DeliveryController) GetDeliveryByAuth(delivery *types.Delivery) (types.Delivery, error) {
	deliveryDb, err := controller.GetDeliveryHandler.Handle(delivery)

	if err != nil {
		return types.Delivery{}, err
	}

	return deliveryDb, err
}
