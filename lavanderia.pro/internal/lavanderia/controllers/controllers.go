package controllers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewLaundryController),
	fx.Provide(NewPingController),
)
