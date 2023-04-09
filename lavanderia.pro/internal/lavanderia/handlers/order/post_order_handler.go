package order

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type PostOrderHandler struct {
	repository *repositories.OrderRepository
}

func NewPostOrderHandler(orderRepository *repositories.OrderRepository) *PostOrderHandler {
	return &PostOrderHandler{
		repository: orderRepository,
	}
}

func (ch *PostOrderHandler) Handle(order *types.Order) (types.Order, error) {
	orderPosted, err := ch.repository.Create(order)

	return orderPosted, err
}
