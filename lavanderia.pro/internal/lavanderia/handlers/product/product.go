package product

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewCreateProductHandler),
	fx.Provide(NewGetAllProductsByBusinessHandler),
	fx.Provide(NewDeleteProductHandler),
	fx.Provide(NewGetProductHandler),
)
