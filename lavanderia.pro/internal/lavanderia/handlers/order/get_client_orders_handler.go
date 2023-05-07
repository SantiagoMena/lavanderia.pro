package order

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type GetClientOrdersHandler struct {
	repository *repositories.OrderRepository
}

func NewGetClientOrdersHandler(orderRepository *repositories.OrderRepository) *GetClientOrdersHandler {
	return &GetClientOrdersHandler{
		repository: orderRepository,
	}
}

func (ch *GetClientOrdersHandler) Handle(client *types.Client) ([]types.Order, error) {
	orderGotten, err := ch.repository.GetAllOrdersByClient(client.ID)

	return orderGotten, err
}
