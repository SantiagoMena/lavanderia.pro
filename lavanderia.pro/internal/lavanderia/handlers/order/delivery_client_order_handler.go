package order

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type DeliveryClientOrderHandler struct {
	repository *repositories.OrderRepository
}

func NewDeliveryClientOrderHandler(orderRepository *repositories.OrderRepository) *DeliveryClientOrderHandler {
	return &DeliveryClientOrderHandler{
		repository: orderRepository,
	}
}

func (ch *DeliveryClientOrderHandler) Handle(order *types.Order) (types.Order, error) {
	orderDelivered, err := ch.repository.DeliveryClient(order)

	return orderDelivered, err
}
