package routers

import "go.uber.org/fx"

var Module = fx.Invoke(
	NewPingRouter,
	NewGetAllBusinessRouter,
	NewPostBusinessRouter,
	NewDeleteBusinessRouter,
	NewUpdateBusinessRouter,
	NewGetBusinessRouter,
	NewPostRegisterBusinessRouter,
	NewPostLoginRouter,
	NewPostRefreshTokenRouter,
	NewPostProductRouter,
	NewGetProductsByBusinessRouter,
	NewDeleteProductRouter,
	NewGetProductRouter,
	NewUpdateProductRouter,
	NewPostRegisterClientRouter,
	NewGetClientRouter,
	NewPostClientRouter,
	NewPutClientRouter,
	NewGetAddressRouter,
	NewPostAddressRouter,
	NewUpdateAddressRouter,
	NewGetAddressesRouter,
	NewDeleteAddressRouter,
	NewPostRegisterDeliveryRouter,
)
