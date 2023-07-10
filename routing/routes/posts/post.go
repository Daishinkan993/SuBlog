package posts

import (
	"SuBlog/models"
	"context"
	"github.com/gorilla/mux"
	"github.com/nicklaw5/go-respond"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request, database *mongo.Database) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	post := models.Post{}

	id := mux.Vars(r)["id"]
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		respond.NewResponse(w).BadRequest(struct {
			Message string
		}{
			Message: "Неверный айди",
		})
		return
	}

	collection := database.Collection("posts")

	result := collection.FindOne(context.Background(), bson.M{"_id": objectID})
	if result.Err() != nil {
		respond.NewResponse(w).NotFound(struct {
			Message string
		}{
			Message: "Айди не найден",
		})
		return
	}

	err = result.Decode(&post)
	if err != nil {
		log.Fatal(err)
	}

	respond.NewResponse(w).Ok(post)
}
