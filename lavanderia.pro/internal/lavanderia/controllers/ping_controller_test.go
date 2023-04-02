package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPing(t *testing.T) {
	controller := NewPingController()

	ping, err := controller.Ping()

	pingExpected := status{
		Status: "ok",
	}

	assert.Equal(t, err, nil, "Error in Ping() ping_controller.go")
	assert.Equal(t, ping, pingExpected, "Ping() returns different result")
}
