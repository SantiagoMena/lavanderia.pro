package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/order"
)

type OrderController struct {
	PostOrderHandler *order.PostOrderHandler
}

func NewOrderController(
	PostOrderHandler *order.PostOrderHandler,
) *OrderController {
	return &OrderController{
		PostOrderHandler: PostOrderHandler,
	}
}

func (controller OrderController) PostOrder(order *types.Order) (types.Order, error) {
	orderDb, err := controller.PostOrderHandler.Handle(order)

	if err != nil {
		return types.Order{}, err
	}

	return orderDb, err
}
