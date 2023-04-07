package client

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRegisterClientHandler),
	fx.Provide(NewGetClientHandler),
	fx.Provide(NewPostClientHandler),
	fx.Provide(NewUpdateClientProfileHandler),
)
