package routers

import "go.uber.org/fx"

var Module = fx.Invoke(
	NewLaundryRouter,
	NewPingRouter,
)
