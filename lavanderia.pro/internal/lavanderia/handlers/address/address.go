package address

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewCreateAddressHandler),
	fx.Provide(NewGetAddressHandler),
	fx.Provide(NewUpdateAddressHandler),
	fx.Provide(NewGetAddressesHandler),
	fx.Provide(NewDeleteAddressHandler),
)
