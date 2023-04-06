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
)
