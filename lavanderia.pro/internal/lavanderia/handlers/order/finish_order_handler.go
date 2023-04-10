package order

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type FinishOrderHandler struct {
	repository *repositories.OrderRepository
}

func NewFinishOrderHandler(orderRepository *repositories.OrderRepository) *FinishOrderHandler {
	return &FinishOrderHandler{
		repository: orderRepository,
	}
}

func (ch *FinishOrderHandler) Handle(order *types.Order) (types.Order, error) {
	orderAssignedPickUp, err := ch.repository.Finish(order)

	return orderAssignedPickUp, err
}
