package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/order"
)

type OrderController struct {
	PostOrderHandler *order.PostOrderHandler
	GetOrderHandler  *order.GetOrderHandler
}

func NewOrderController(
	PostOrderHandler *order.PostOrderHandler,
	GetOrderHandler *order.GetOrderHandler,
) *OrderController {
	return &OrderController{
		PostOrderHandler: PostOrderHandler,
		GetOrderHandler:  GetOrderHandler,
	}
}

func (controller OrderController) PostOrder(order *types.Order) (types.Order, error) {
	orderDb, err := controller.PostOrderHandler.Handle(order)

	if err != nil {
		return types.Order{}, err
	}

	return orderDb, err
}

func (controller OrderController) GetOrder(order *types.Order) (types.Order, error) {
	orderDb, err := controller.GetOrderHandler.Handle(order)

	if err != nil {
		return types.Order{}, err
	}

	return orderDb, err
}
