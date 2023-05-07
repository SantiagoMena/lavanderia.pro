package order

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type PostOrderHandler struct {
	repository        *repositories.OrderRepository
	productRepository *repositories.ProductRepository
}

func NewPostOrderHandler(orderRepository *repositories.OrderRepository, productRepository *repositories.ProductRepository) *PostOrderHandler {
	return &PostOrderHandler{
		repository:        orderRepository,
		productRepository: productRepository,
	}
}

func (ch *PostOrderHandler) Handle(order *types.Order) (types.Order, error) {
	// Find products
	for i := 0; i < len(order.Products); i++ {
		product, errorGetProduct := ch.productRepository.Get(&order.Products[i].Product)

		if errorGetProduct != nil {
			return types.Order{}, errorGetProduct
		}

		order.Products[i].Product = product
	}

	orderPosted, err := ch.repository.Create(order)

	return orderPosted, err
}
