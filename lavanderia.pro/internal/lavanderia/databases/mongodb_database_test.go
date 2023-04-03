package databases

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func TestCreate(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()

	mongo := NewMongoDatabase(config)

	object := bson.D{
		{Key: "test", Value: "test"},
	}

	mongoInsertResult, err := mongo.Create("TESTCOLLECTION", object)

	fmt.Println(mongoInsertResult)

	assert.Equal(t, err, nil)
}

func TestUpdateOne(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	collection := "TESTCOLLECTION"
	config := config.NewConfig()

	mongo := NewMongoDatabase(config)

	object := bson.D{
		{Key: "test", Value: "test"},
	}

	mongoInsertResult, err := mongo.Create(collection, object)

	insertedId := mongoInsertResult.InsertedID

	filter := bson.D{{"_id", insertedId}}
	update := bson.D{{"$set", bson.D{{"test", "updated"}}}}

	objectUpdated, err := mongo.UpdateOne(collection, filter, update)

	fmt.Println(objectUpdated)

	type TestType struct {
		Id   string `json:"id" bson:"_id"`
		Test string `json:"test" bson:"test"`
	}

	var updatedObject TestType

	// convert m to s
	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &updatedObject)

	insertedIdTest := insertedId.(primitive.ObjectID).Hex()

	assert.Equal(t, err, nil)
	assert.Equal(t, updatedObject.Id, insertedIdTest, "Not equal id updated")
}
