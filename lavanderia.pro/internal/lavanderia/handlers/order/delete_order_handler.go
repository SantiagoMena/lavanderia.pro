package order

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type DeleteOrderHandler struct {
	repository *repositories.OrderRepository
}

func NewDeleteOrderHandler(orderRepository *repositories.OrderRepository) *DeleteOrderHandler {
	return &DeleteOrderHandler{
		repository: orderRepository,
	}
}

func (ch *DeleteOrderHandler) Handle(order *types.Order) (types.Order, error) {
	orderDeleted, err := ch.repository.Delete(order)

	return orderDeleted, err
}
