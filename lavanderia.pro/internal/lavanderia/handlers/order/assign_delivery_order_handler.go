package order

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type AssignDeliveryOrderHandler struct {
	repository *repositories.OrderRepository
}

func NewAssignDeliveryOrderHandler(orderRepository *repositories.OrderRepository) *AssignDeliveryOrderHandler {
	return &AssignDeliveryOrderHandler{
		repository: orderRepository,
	}
}

func (ch *AssignDeliveryOrderHandler) Handle(order *types.Order) (types.Order, error) {
	orderAssignedDelivery, err := ch.repository.AssignDelivery(order)

	return orderAssignedDelivery, err
}
