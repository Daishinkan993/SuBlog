package DB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Init(url string, databaseName string) (*mongo.Database, error) {
	connect, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	database := connect.Database(databaseName)

	return database, err
}
