package databases

import (
	"context"
	"errors"

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
	FindOne(collection string, filter bson.D) (bson.M, error)
	FindAllFilter(collection string, filter bson.D) (*mongo.Cursor, error)
	FindAllFilterSort(collection string, filter bson.D, sort bson.D) (*mongo.Cursor, error)
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
	businessDb := db.mongo.Collection(collection)

	result, err := businessDb.Find(context.Background(), bson.D{})

	if err != nil {
		log.Panic(err)
	}

	return result, err
}

func (db database) Create(collection string, object bson.D) (*mongo.InsertOneResult, error) {
	businessDb := db.mongo.Collection(collection)

	// result, err := businessDb.Find(context.Background(), bson.D{})
	result, err := businessDb.InsertOne(context.TODO(), object)

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

	businessDb := db.mongo.Collection(collection)

	opts := options.FindOneAndUpdate().SetUpsert(true)
	var object bson.M
	err := businessDb.FindOneAndUpdate(
		context.TODO(),
		filter,
		update,
		opts,
	).Decode(&object)

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		// if err == mongo.ErrNoDocuments {
		// 	return bson.M{}, err
		// }

		// return bson.M{}, err
		return bson.M{}, errors.New("document to update not found")
	}

	// if err != nil {
	// 	return bson.M{}, err
	// }

	// defer func() {
	// 	if err := db.client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	return object, err
}

func (db database) FindOne(collection string, filter bson.D) (bson.M, error) {
	businessDb := db.mongo.Collection(collection)

	var result bson.M

	err := businessDb.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}

func (db database) FindAllFilter(collection string, filter bson.D) (*mongo.Cursor, error) {
	businessDb := db.mongo.Collection(collection)

	result, err := businessDb.Find(context.Background(), filter)

	if err != nil {
		log.Panic(err)
	}

	return result, err
}

func (db database) FindAllFilterSort(collection string, filter bson.D, sort bson.D) (*mongo.Cursor, error) {
	businessDb := db.mongo.Collection(collection)
	findOptions := options.Find()
	findOptions.SetSort(sort)

	result, err := businessDb.Find(context.Background(), filter, findOptions)

	if err != nil {
		log.Panic(err)
	}

	return result, err
}
