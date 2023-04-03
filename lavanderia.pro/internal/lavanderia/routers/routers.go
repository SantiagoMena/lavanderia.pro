package routers

import "go.uber.org/fx"

var Module = fx.Invoke(
	NewGetLaundriesRouter,
	NewPostLaundrysRouter,
	NewPingRouter,
)
