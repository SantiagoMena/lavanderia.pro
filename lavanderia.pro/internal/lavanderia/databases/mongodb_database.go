package databases

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	FindAll(collection string) (*mongo.Cursor, error)
}

type database struct {
	mongo  mongo.Database
	client mongo.Client
}

func NewMongoDatabase(uri string, db string) Database {
	mongoDb, mongoClient := Mongodb(uri, db)

	return database{
		mongo:  *mongoDb,
		client: *mongoClient,
	}
}

func Mongodb(uri string, database string) (*mongo.Database, *mongo.Client) {
	// uri := os.Getenv("MONGODB_URI")
	// database := os.Getenv("MONGODB_DB")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic(err)
	}
	databaseConnect := client.Database(database)

	return databaseConnect, client
}

func (db database) FindAll(collection string) (*mongo.Cursor, error) {
	// databaseConnect, client := Mongodb()

	laundryDb := db.mongo.Collection(collection)

	result, err := laundryDb.Find(context.Background(), bson.D{})

	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := db.client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	return result, err
}
