package databases

import (
	// "encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	// "io/ioutil"
	// "net/http"
	// "net/http/httptest"
	"github.com/joho/godotenv"
	"testing"
)

func TestFindAll(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env file found")
	}

	cursor, err := FindAll("COLLECTION")

	fmt.Println(cursor)

	assert.Equal(t, err, nil)
}
