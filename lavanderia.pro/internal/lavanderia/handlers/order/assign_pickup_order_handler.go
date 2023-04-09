package order

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type AssignPickUpOrderHandler struct {
	repository *repositories.OrderRepository
}

func NewAssignPickUpOrderHandler(orderRepository *repositories.OrderRepository) *AssignPickUpOrderHandler {
	return &AssignPickUpOrderHandler{
		repository: orderRepository,
	}
}

func (ch *AssignPickUpOrderHandler) Handle(order *types.Order) (types.Order, error) {
	orderAssignedPickUp, err := ch.repository.AssignPickUp(order)

	return orderAssignedPickUp, err
}
