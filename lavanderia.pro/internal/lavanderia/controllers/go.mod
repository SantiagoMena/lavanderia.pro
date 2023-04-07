module lavanderia.pro/internal/lavanderia/controllers

go 1.19

replace lavanderia.pro/api/types => ../../../api/types

replace lavanderia.pro/internal/lavanderia/repositories => ../repositories

require (
	github.com/joho/godotenv v1.5.1
	github.com/stretchr/testify v1.8.2
	go.uber.org/fx v1.19.2
	golang.org/x/crypto v0.7.0
	lavanderia.pro/api/types v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/config v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/databases v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/handlers/address v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/handlers/auth v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/handlers/business v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/handlers/client v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/handlers/product v0.0.0-00010101000000-000000000000
	lavanderia.pro/internal/lavanderia/repositories v0.0.0-00010101000000-000000000000
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang-jwt/jwt/v5 v5.0.0-rc.2 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.5.5 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.mongodb.org/mongo-driver v1.11.3 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.16.1 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace lavanderia.pro/internal/lavanderia/routers => ../routers

replace lavanderia.pro/internal/lavanderia/databases => ../databases

replace lavanderia.pro/internal/lavanderia/config => ../config

replace lavanderia.pro/internal/lavanderia/handlers/business => ../handlers/business

replace lavanderia.pro/internal/lavanderia/handlers/auth => ../handlers/auth

replace lavanderia.pro/internal/lavanderia/handlers/product => ../handlers/product

replace lavanderia.pro/internal/lavanderia/handlers/client => ../handlers/client

replace lavanderia.pro/internal/lavanderia/handlers/address => ../handlers/address
