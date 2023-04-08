module lavanderia.pro/cmd/lavanderia

go 1.19

replace lavanderia.pro/internal/lavanderia/controllers => ../../internal/lavanderia/controllers

require (
	github.com/gin-gonic/gin v1.9.0
	github.com/joho/godotenv v1.5.1
	github.com/stretchr/testify v1.8.2
	go.uber.org/fx v1.19.2
	lavanderia.pro/internal/lavanderia/config v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/controllers v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/databases v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/handlers/address v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/handlers/auth v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/handlers/business v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/handlers/client v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/handlers/delivery v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/handlers/product v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/middlewares v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/repositories v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/routers v0.0.0-00010101000000-000000000000
)

require (
	github.com/bytedance/sonic v1.8.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.0.0-rc.2 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.9 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.mongodb.org/mongo-driver v1.11.3 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.16.1 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	lavanderia.pro/api/types v0.0.0-00010101000000-000000000000 // indirect
)

replace lavanderia.pro/internal/lavanderia/repositories => ../../internal/lavanderia/repositories

replace lavanderia.pro/internal/lavanderia/databases => ../../internal/lavanderia/databases

replace lavanderia.pro/api/types => ../../api/types

replace lavanderia.pro/internal/lavanderia/routers => ../../internal/lavanderia/routers

replace lavanderia.pro/internal/lavanderia/config => ../../internal/lavanderia/config

replace lavanderia.pro/internal/lavanderia/middlewares => ../../internal/lavanderia/middlewares

replace lavanderia.pro/internal/lavanderia/handlers/auth => ../../internal/lavanderia/handlers/auth

replace lavanderia.pro/internal/lavanderia/handlers/business => ../../internal/lavanderia/handlers/business

replace lavanderia.pro/internal/lavanderia/handlers/delivery => ../../internal/lavanderia/handlers/delivery

replace lavanderia.pro/internal/lavanderia/handlers/client => ../../internal/lavanderia/handlers/client

replace lavanderia.pro/internal/lavanderia/handlers/product => ../../internal/lavanderia/handlers/product

replace lavanderia.pro/internal/lavanderia/handlers/address => ../../internal/lavanderia/handlers/address
