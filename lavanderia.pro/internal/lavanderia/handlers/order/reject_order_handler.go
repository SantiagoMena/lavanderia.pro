package order

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type RejectOrderHandler struct {
	repository *repositories.OrderRepository
}

func NewRejectOrderHandler(orderRepository *repositories.OrderRepository) *RejectOrderHandler {
	return &RejectOrderHandler{
		repository: orderRepository,
	}
}

func (ch *RejectOrderHandler) Handle(order *types.Order) (types.Order, error) {
	orderRejected, err := ch.repository.Reject(order)

	return orderRejected, err
}
