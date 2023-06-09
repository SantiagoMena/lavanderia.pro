package business

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewGetAllBusinessHandler),
	fx.Provide(NewCreateBusinessHandler),
	fx.Provide(NewDeleteBusinessHandler),
	fx.Provide(NewUpdateBusinessHandler),
	fx.Provide(NewGetBusinessHandler),
	fx.Provide(NewRegisterBusinessHandler),
	fx.Provide(NewGetAllBusinessByAuthHandler),
	fx.Provide(NewRegisterBusinessDeliveryHandler),
	fx.Provide(NewSearchBusinessHandler),
)
