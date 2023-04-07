package repositories

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewBusinessRepository),
	fx.Provide(NewAuthRepository),
	fx.Provide(NewProductRepository),
	fx.Provide(NewClientRepository),
	fx.Provide(NewAddressRepository),
)
