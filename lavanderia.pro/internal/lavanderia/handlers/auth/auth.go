package auth

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewCreateJWTHandler),
	fx.Provide(NewLoginHandler),
)
