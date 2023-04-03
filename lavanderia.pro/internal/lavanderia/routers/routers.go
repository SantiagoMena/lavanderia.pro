package routers

import "go.uber.org/fx"

var Module = fx.Invoke(
	NewGetLaundriesRouter,
	NewPostLaundrysRouter,
	NewDeleteLaundrysRouter,
	NewUpdateLaundryRouter,
	NewGetLaundryRouter,
	NewPingRouter,
)
