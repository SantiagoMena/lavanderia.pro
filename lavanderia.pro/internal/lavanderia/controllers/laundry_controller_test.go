package controllers

import (
	// "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	// "io/ioutil"
	"net/http"
	"net/http/httptest"
	// "lavanderia.pro/cmd/lavanderia"
	"testing"
)

func TestLaundries(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env file found")
	}

	router := gin.Default()
	router.GET("/laundries", Laundries)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/laundries", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// bodySb, err := ioutil.ReadAll(w.Body)
	// if err != nil {
	// 	t.Fatalf("Error reading body: %v\n", err)
	// }

	// var decodedResponse interface{}
	// err = json.Unmarshal(bodySb, &decodedResponse)
	// if err != nil {
	// 	t.Fatalf("Cannot decode response <%p> from server. Err: %v", bodySb, err)
	// }
}
