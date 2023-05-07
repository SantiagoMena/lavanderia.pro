package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/order"
)

type OrderController struct {
	PostOrderHandler           *order.PostOrderHandler
	GetOrderHandler            *order.GetOrderHandler
	DeleteOrderHandler         *order.DeleteOrderHandler
	AcceptOrderHandler         *order.AcceptOrderHandler
	RejectOrderHandler         *order.RejectOrderHandler
	AssignPickUpOrderHandler   *order.AssignPickUpOrderHandler
	PickUpClientOrderHandler   *order.PickUpClientOrderHandler
	ProcessOrderHandler        *order.ProcessOrderHandler
	FinishOrderHandler         *order.FinishOrderHandler
	AssignDeliveryOrderHandler *order.AssignDeliveryOrderHandler
	PickUpBusinessOrderHandler *order.PickUpBusinessOrderHandler
	DeliveryClientOrderHandler *order.DeliveryClientOrderHandler
	GetClientOrdersHandler     *order.GetClientOrdersHandler
}

func NewOrderController(
	PostOrderHandler *order.PostOrderHandler,
	GetOrderHandler *order.GetOrderHandler,
	DeleteOrderHandler *order.DeleteOrderHandler,
	AcceptOrderHandler *order.AcceptOrderHandler,
	RejectOrderHandler *order.RejectOrderHandler,
	AssignPickUpOrderHandler *order.AssignPickUpOrderHandler,
	PickUpClientOrderHandler *order.PickUpClientOrderHandler,
	ProcessOrderHandler *order.ProcessOrderHandler,
	FinishOrderHandler *order.FinishOrderHandler,
	AssignDeliveryOrderHandler *order.AssignDeliveryOrderHandler,
	PickUpBusinessOrderHandler *order.PickUpBusinessOrderHandler,
	DeliveryClientOrderHandler *order.DeliveryClientOrderHandler,
	GetClientOrdersHandler *order.GetClientOrdersHandler,
) *OrderController {
	return &OrderController{
		PostOrderHandler:           PostOrderHandler,
		GetOrderHandler:            GetOrderHandler,
		DeleteOrderHandler:         DeleteOrderHandler,
		AcceptOrderHandler:         AcceptOrderHandler,
		RejectOrderHandler:         RejectOrderHandler,
		AssignPickUpOrderHandler:   AssignPickUpOrderHandler,
		PickUpClientOrderHandler:   PickUpClientOrderHandler,
		ProcessOrderHandler:        ProcessOrderHandler,
		FinishOrderHandler:         FinishOrderHandler,
		AssignDeliveryOrderHandler: AssignDeliveryOrderHandler,
		PickUpBusinessOrderHandler: PickUpBusinessOrderHandler,
		DeliveryClientOrderHandler: DeliveryClientOrderHandler,
		GetClientOrdersHandler:     GetClientOrdersHandler,
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

func (controller OrderController) PickUpClientOrder(order *types.Order) (types.Order, error) {
	orderDb, err := controller.PickUpClientOrderHandler.Handle(order)

	if err != nil {
		return types.Order{}, err
	}

	return orderDb, err
}

func (controller OrderController) ProcessOrder(order *types.Order) (types.Order, error) {
	orderDb, err := controller.ProcessOrderHandler.Handle(order)

	if err != nil {
		return types.Order{}, err
	}

	return orderDb, err
}

func (controller OrderController) FinishOrder(order *types.Order) (types.Order, error) {
	orderDb, err := controller.FinishOrderHandler.Handle(order)

	if err != nil {
		return types.Order{}, err
	}

	return orderDb, err
}

func (controller OrderController) AssignDeliveryOrder(order *types.Order) (types.Order, error) {
	orderDb, err := controller.AssignDeliveryOrderHandler.Handle(order)

	if err != nil {
		return types.Order{}, err
	}

	return orderDb, err
}

func (controller OrderController) PickUpBusinessOrder(order *types.Order) (types.Order, error) {
	orderDb, err := controller.PickUpBusinessOrderHandler.Handle(order)

	if err != nil {
		return types.Order{}, err
	}

	return orderDb, err
}

func (controller OrderController) DeliveryClientOrder(order *types.Order) (types.Order, error) {
	orderDb, err := controller.DeliveryClientOrderHandler.Handle(order)

	if err != nil {
		return types.Order{}, err
	}

	return orderDb, err
}

func (controller OrderController) GetClientOrders(client *types.Client) ([]types.Order, error) {
	orderDb, err := controller.GetClientOrdersHandler.Handle(client)

	if err != nil {
		return []types.Order{}, err
	}

	return orderDb, err
}
