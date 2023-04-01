package databases

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mongodb() (*mongo.Database, *mongo.Client) {
	uri := os.Getenv("MONGODB_URI")
	database := os.Getenv("MONGODB_DB")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	databaseConnect := client.Database(database)

	return databaseConnect, client
}

func FindAll(collection string) (*mongo.Cursor, error) {
	databaseConnect, client := Mongodb()

	laundryDb := databaseConnect.Collection(collection)

	result, err := laundryDb.Find(context.Background(), bson.D{})

	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	return result, err
}
