package DB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindOne(collection mongo.Collection, m bson.M) (*mongo.SingleResult, error) {
	result := collection.FindOne(context.Background(), m)
	if result.Err() != nil {
		return nil, result.Err()
	}

	return result, nil
}
