package order

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type GetOrderHandler struct {
	repository *repositories.OrderRepository
}

func NewGetOrderHandler(orderRepository *repositories.OrderRepository) *GetOrderHandler {
	return &GetOrderHandler{
		repository: orderRepository,
	}
}

func (ch *GetOrderHandler) Handle(order *types.Order) (types.Order, error) {
	orderGotten, err := ch.repository.Get(order)

	return orderGotten, err
}
