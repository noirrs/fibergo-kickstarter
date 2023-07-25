package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connect to database and return a connection
func Connect() *mongo.Client {

	url := LoadPreferences().ConnURL

	client, err := mongo.NewClient(options.Client().ApplyURI(url))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client

}

// global client instance
var DB *mongo.Client = Connect()

func GetCollection(collectionName string) *mongo.Collection {
	name := LoadPreferences().Dbname
	collection := DB.Database(name).Collection(collectionName)
	return collection
}
