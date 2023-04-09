package order

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewPostOrderHandler),
	fx.Provide(NewGetOrderHandler),
	fx.Provide(NewDeleteOrderHandler),
)
