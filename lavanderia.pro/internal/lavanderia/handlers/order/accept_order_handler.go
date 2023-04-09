package order

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type AcceptOrderHandler struct {
	repository *repositories.OrderRepository
}

func NewAcceptOrderHandler(orderRepository *repositories.OrderRepository) *AcceptOrderHandler {
	return &AcceptOrderHandler{
		repository: orderRepository,
	}
}

func (ch *AcceptOrderHandler) Handle(order *types.Order) (types.Order, error) {
	orderAccepted, err := ch.repository.Accept(order)

	return orderAccepted, err
}
