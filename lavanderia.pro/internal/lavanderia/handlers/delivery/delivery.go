package delivery

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRegisterDeliveryHandler),
	fx.Provide(NewPostDeliveryHandler),
)
