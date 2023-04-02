package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"github.com/gin-gonic/gin"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/controllers"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/repositories"
	"lavanderia.pro/internal/lavanderia/routers"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}

	// r := routers.SetupRouter()

	// r.Run()

	fx.New(
		fx.Provide(config.NewConfig),
		fx.Provide(databases.NewMongoDatabase),
		repositories.Module,
		controllers.Module,
		fx.Provide(provideGinEngine),
		routers.Module,
		fx.Invoke(
			startServer,
		),
	).Run()
}

// func registerService(ginEngine *gin.Engine, userSvcRouter usersvc.Router) {
// 	gGroup := ginEngine.Group("api/v1")
// 	userSvcRouter.Register(gGroup)
// }

func startServer(ginEngine *gin.Engine, lifecycle fx.Lifecycle) {
	port := "8080"
	server := http.Server{
		Addr:    ":" + port,
		Handler: ginEngine,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("run on port:", port)
			go func() {
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					fmt.Errorf("failed to listen and serve from server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}

func provideGinEngine() *gin.Engine {
	return gin.Default()
}
