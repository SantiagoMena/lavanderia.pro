package main

import (
	"context"
	// "errors"
	// "fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/fx"

	// "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/controllers"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/repositories"
	"lavanderia.pro/internal/lavanderia/routers"
)

func TestMain(t *testing.T) {
	if err := godotenv.Load("../../.env.test"); err != nil {
		log.Println("No .env.test file found")
	}

	app := fx.New(
		fx.Provide(config.NewConfig),
		fx.Provide(databases.NewMongoDatabase),
		repositories.Module,
		controllers.Module,
		fx.Provide(provideGinEngine),
		routers.Module,
		fx.Invoke(
			startServer,
		),
	)

	// In a typical application, we could just use app.Run() here. Since we
	// don't want this example to run forever, we'll use the more-explicit Start
	// and Stop.
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	errStart := app.Start(startCtx)
	assert.Nil(t, errStart, "Server is no starting")

	// Normally, we'd block here with <-app.Done(). Instead, we'll make an HTTP
	// request to demonstrate that our server is running.
	_, errServe := http.Get("http://localhost:8080/")
	assert.Nil(t, errServe, "Server is not running")

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	errStop := app.Stop(stopCtx)

	assert.Nil(t, errStop, "Server is no stopping")

}
