package databases

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
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
