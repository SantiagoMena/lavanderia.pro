package databases

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFindAll(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	uri := os.Getenv("MONGODB_URI")
	database := os.Getenv("MONGODB_DB")

	mongo := NewMongoDatabase(uri, database)

	cursor, err := mongo.FindAll("COLLECTION")

	fmt.Println(cursor)

	assert.Equal(t, err, nil)
}
