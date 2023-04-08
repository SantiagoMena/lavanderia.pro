package delivery

import (
	"errors"

	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type PostDeliveryHandler struct {
	repository *repositories.DeliveryRepository
}

func NewPostDeliveryHandler(deliveryRepository *repositories.DeliveryRepository) *PostDeliveryHandler {
	return &PostDeliveryHandler{
		repository: deliveryRepository,
	}
}

func (ch *PostDeliveryHandler) Handle(delivery *types.Delivery) (types.Delivery, error) {

	// find delivery by auth
	deliveryFound, _ := ch.repository.GetDeliveryByAuth(delivery)
	emptyDelivery := types.Delivery{}

	// if exists error
	if deliveryFound != emptyDelivery {
		return types.Delivery{}, errors.New("delivery already registered")
	}

	// if not create
	deliveryPosted, err := ch.repository.Create(delivery)

	return deliveryPosted, err
}
