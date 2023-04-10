package order

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewPostOrderHandler),
	fx.Provide(NewGetOrderHandler),
	fx.Provide(NewDeleteOrderHandler),
	fx.Provide(NewAcceptOrderHandler),
	fx.Provide(NewRejectOrderHandler),
	fx.Provide(NewAssignPickUpOrderHandler),
	fx.Provide(NewPickUpClientOrderHandler),
)
