package databases

import (
	"context"
	// "go/types"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"lavanderia.pro/internal/lavanderia/config"
)

type Database interface {
	FindAll(collection string) (*mongo.Cursor, error)
	Create(collection string, object bson.D) (*mongo.InsertOneResult, error)
	UpdateOne(collection string, filter bson.D, update bson.D) (bson.M, error)
}

type database struct {
	mongo  mongo.Database
	client mongo.Client
}

func NewMongoDatabase(config *config.Config) Database {
	mongoDb, mongoClient := Mongodb(config.MongodbUri, config.MongodbDb)

	return database{
		mongo:  *mongoDb,
		client: *mongoClient,
	}
}

func Mongodb(uri string, database string) (*mongo.Database, *mongo.Client) {
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
	laundryDb := db.mongo.Collection(collection)

	result, err := laundryDb.Find(context.Background(), bson.D{})

	if err != nil {
		log.Panic(err)
	}

	// defer func() {
	// 	if err := db.client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	return result, err
}

func (db database) Create(collection string, object bson.D) (*mongo.InsertOneResult, error) {
	laundryDb := db.mongo.Collection(collection)

	// result, err := laundryDb.Find(context.Background(), bson.D{})
	result, err := laundryDb.InsertOne(context.TODO(), object)

	if err != nil {
		log.Panic(err)
	}

	// defer func() {
	// 	if err := db.client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	return result, err
}

func (db database) UpdateOne(collection string, filter bson.D, update bson.D) (bson.M, error) {

	laundryDb := db.mongo.Collection(collection)

	opts := options.FindOneAndUpdate().SetUpsert(true)
	var object bson.M
	err := laundryDb.FindOneAndUpdate(
		context.TODO(),
		filter,
		update,
		opts,
	).Decode(&object)

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Fatal(err)
	}

	if err != nil {
		log.Panic(err)
	}

	// defer func() {
	// 	if err := db.client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	return object, err
}
