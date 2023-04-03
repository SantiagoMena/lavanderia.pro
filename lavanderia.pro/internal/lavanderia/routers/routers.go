package routers

import "go.uber.org/fx"

var Module = fx.Invoke(
	NewGetAllBusinessRouter,
	NewPostBusinesssRouter,
	NewDeleteBusinesssRouter,
	NewUpdateBusinessRouter,
	NewGetBusinessRouter,
	NewPingRouter,
)
