package delivery

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type UpdateDeliveryProfileHandler struct {
	repository *repositories.DeliveryRepository
}

func NewUpdateDeliveryProfileHandler(deliveryRepository *repositories.DeliveryRepository) *UpdateDeliveryProfileHandler {
	return &UpdateDeliveryProfileHandler{
		repository: deliveryRepository,
	}
}

func (ch *UpdateDeliveryProfileHandler) Handle(delivery *types.Delivery) (types.Delivery, error) {
	// find delivery
	deliveryFound, errorFind := ch.repository.GetDeliveryByAuth(delivery)
	if errorFind != nil {
		return types.Delivery{}, errorFind
	}

	deliveryFound.Name = delivery.Name

	// update delivery
	deliveryUpdated, errorUpdate := ch.repository.Update(&deliveryFound)

	return deliveryUpdated, errorUpdate
}
