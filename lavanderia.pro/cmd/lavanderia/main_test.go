package main

import (
	"encoding/json"
	// "fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	bodySb, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("Error reading body: %v\n", err)
	}

	var decodedResponse interface{}
	err = json.Unmarshal(bodySb, &decodedResponse)
	if err != nil {
		t.Fatalf("Cannot decode response <%p> from server. Err: %v", bodySb, err)
	}

	assert.Equal(t, map[string]interface{}{"status": "ok"}, decodedResponse, "Should return status:ok")
}
