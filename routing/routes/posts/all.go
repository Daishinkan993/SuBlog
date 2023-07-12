package posts

import (
	"SuBlog/models"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

type PageData struct {
	Posts []models.Post
}

func All(w http.ResponseWriter, r *http.Request, database *mongo.Database) {
	collection := database.Collection("posts")

	pageData := PageData{}

	data, err := collection.Find(context.Background(), bson.M{}, options.Find().SetLimit(10).SetProjection(bson.M{"content": 0}).SetSort(bson.M{"timestamp": -1}))
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close(context.TODO())

	for data.Next(context.Background()) {
		var post models.Post
		err := data.Decode(&post)
		if err != nil {
			log.Fatal(err)
		}
		pageData.Posts = append(pageData.Posts, post)
	}

	result, err := json.Marshal(pageData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(result))
}
