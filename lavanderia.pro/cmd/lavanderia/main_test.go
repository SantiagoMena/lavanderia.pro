package main

import (
	"context"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	if err := godotenv.Load("../../.env.test"); err != nil {
		log.Println("No .env.test file found")
	}

	app := MakeApp()
	expectedApp := fx.New()
	assert.IsType(t, expectedApp, app, "MakeApp() is not returning *fx.App")

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
