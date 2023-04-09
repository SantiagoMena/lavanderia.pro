package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/order"
)

type OrderController struct {
	PostOrderHandler         *order.PostOrderHandler
	GetOrderHandler          *order.GetOrderHandler
	DeleteOrderHandler       *order.DeleteOrderHandler
	AcceptOrderHandler       *order.AcceptOrderHandler
	RejectOrderHandler       *order.RejectOrderHandler
	AssignPickUpOrderHandler *order.AssignPickUpOrderHandler
}

func NewOrderController(
	PostOrderHandler *order.PostOrderHandler,
	GetOrderHandler *order.GetOrderHandler,
	DeleteOrderHandler *order.DeleteOrderHandler,
	AcceptOrderHandler *order.AcceptOrderHandler,
	RejectOrderHandler *order.RejectOrderHandler,
	AssignPickUpOrderHandler *order.AssignPickUpOrderHandler,
) *OrderController {
	return &OrderController{
		PostOrderHandler:         PostOrderHandler,
		GetOrderHandler:          GetOrderHandler,
		DeleteOrderHandler:       DeleteOrderHandler,
		AcceptOrderHandler:       AcceptOrderHandler,
		RejectOrderHandler:       RejectOrderHandler,
		AssignPickUpOrderHandler: AssignPickUpOrderHandler,
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

func (controller OrderController) DeleteOrder(order *types.Order) (types.Order, error) {
	orderDb, err := controller.DeleteOrderHandler.Handle(order)

	if err != nil {
		return types.Order{}, err
	}

	return orderDb, err
}

func (controller OrderController) AcceptOrder(order *types.Order) (types.Order, error) {
	orderDb, err := controller.AcceptOrderHandler.Handle(order)

	if err != nil {
		return types.Order{}, err
	}

	return orderDb, err
}

func (controller OrderController) RejectOrder(order *types.Order) (types.Order, error) {
	orderDb, err := controller.RejectOrderHandler.Handle(order)

	if err != nil {
		return types.Order{}, err
	}

	return orderDb, err
}

func (controller OrderController) AssignPickUpOrder(order *types.Order) (types.Order, error) {
	orderDb, err := controller.AssignPickUpOrderHandler.Handle(order)

	if err != nil {
		return types.Order{}, err
	}

	return orderDb, err
}
