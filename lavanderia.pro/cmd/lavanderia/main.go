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
	"lavanderia.pro/internal/lavanderia/handlers/auth"
	"lavanderia.pro/internal/lavanderia/handlers/business"
	"lavanderia.pro/internal/lavanderia/middlewares"
	"lavanderia.pro/internal/lavanderia/repositories"
	"lavanderia.pro/internal/lavanderia/routers"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}

	RunServer()
}

func RunServer() {
	app := MakeApp()

	app.Run()
}

func MakeApp() *fx.App {
	return fx.New(
		fx.Provide(config.NewConfig),
		fx.Provide(databases.NewMongoDatabase),
		repositories.Module,
		controllers.Module,
		fx.Provide(provideGinEngine),
		middlewares.Module,
		routers.Module,
		business.Module,
		auth.Module,
		fx.Invoke(
			startServer,
		),
	)
}

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
