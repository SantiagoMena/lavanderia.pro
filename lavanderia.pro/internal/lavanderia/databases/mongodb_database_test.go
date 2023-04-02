package databases

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"lavanderia.pro/internal/lavanderia/config"
	"testing"
)

func TestFindAll(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()

	mongo := NewMongoDatabase(config)

	cursor, err := mongo.FindAll("COLLECTION")

	fmt.Println(cursor)

	assert.Equal(t, err, nil)
}
