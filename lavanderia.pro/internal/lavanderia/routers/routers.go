package routers

import "go.uber.org/fx"

var Module = fx.Invoke(
	NewGetAllBusinessRouter,
	// NewPostBusinessRouter,
	NewDeleteBusinessRouter,
	NewUpdateBusinessRouter,
	NewGetBusinessRouter,
	NewPostRegisterBusinessRouter,
	NewPostLoginRouter,
	NewPingRouter,
)
