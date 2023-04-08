package delivery

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type GetDeliveryHandler struct {
	repositoryDelivery *repositories.DeliveryRepository
}

func NewGetDeliveryHandler(repositoryDelivery *repositories.DeliveryRepository) *GetDeliveryHandler {
	return &GetDeliveryHandler{
		repositoryDelivery: repositoryDelivery,
	}
}

func (ch GetDeliveryHandler) Handle(delivery *types.Delivery) (types.Delivery, error) {
	deliveryDb, err := ch.repositoryDelivery.GetDeliveryByAuth(delivery)

	return deliveryDb, err
}
