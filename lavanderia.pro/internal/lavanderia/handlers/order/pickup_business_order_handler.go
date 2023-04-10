package order

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type PickUpBusinessOrderHandler struct {
	repository *repositories.OrderRepository
}

func NewPickUpBusinessOrderHandler(orderRepository *repositories.OrderRepository) *PickUpBusinessOrderHandler {
	return &PickUpBusinessOrderHandler{
		repository: orderRepository,
	}
}

func (ch *PickUpBusinessOrderHandler) Handle(order *types.Order) (types.Order, error) {
	orderPicketUp, err := ch.repository.PickUpBusiness(order)

	return orderPicketUp, err
}
