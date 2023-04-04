package controllers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewBusinessController),
	fx.Provide(NewAuthController),
	fx.Provide(NewPingController),
)
