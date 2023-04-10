package order

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type ProcessOrderHandler struct {
	repository *repositories.OrderRepository
}

func NewProcessOrderHandler(orderRepository *repositories.OrderRepository) *ProcessOrderHandler {
	return &ProcessOrderHandler{
		repository: orderRepository,
	}
}

func (ch *ProcessOrderHandler) Handle(order *types.Order) (types.Order, error) {
	orderAssignedPickUp, err := ch.repository.Process(order)

	return orderAssignedPickUp, err
}
