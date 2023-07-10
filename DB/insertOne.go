package DB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOne(collection mongo.Collection, m bson.M) error {
	_, err := collection.InsertOne(context.Background(), m)
	if err != nil {
		return err
	}

	return nil
}
