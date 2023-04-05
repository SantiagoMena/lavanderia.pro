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

func TestFindOne(t *testing.T) {
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

	insertedIdTest := insertedId.(primitive.ObjectID).Hex()

	filter := bson.D{{"_id", mongoInsertResult.InsertedID}}

	objectFound, err := mongo.FindOne(collection, filter)

	type TestType struct {
		Id   string `json:"id" bson:"_id"`
		Test string `json:"test" bson:"test"`
	}

	var foundObjectTest TestType

	// convert m to s
	objectUpdt, errMarshal := bson.Marshal(objectFound)
	bson.Unmarshal(objectUpdt, &foundObjectTest)

	assert.Equal(t, err, nil, "Error on found object")
	assert.Equal(t, errMarshal, nil, "Error on marshal object")
	assert.NotEmpty(t, foundObjectTest, "FindOne() not found any")
	assert.Equal(t, insertedIdTest, foundObjectTest.Id, "FindOne() object ID is different from created")
}
