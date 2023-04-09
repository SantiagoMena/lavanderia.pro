package order

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type PickUpClientOrderHandler struct {
	repository *repositories.OrderRepository
}

func NewPickUpClientOrderHandler(orderRepository *repositories.OrderRepository) *PickUpClientOrderHandler {
	return &PickUpClientOrderHandler{
		repository: orderRepository,
	}
}

func (ch *PickUpClientOrderHandler) Handle(order *types.Order) (types.Order, error) {
	orderAssignedPickUp, err := ch.repository.PickUpClient(order)

	return orderAssignedPickUp, err
}
