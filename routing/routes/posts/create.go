package posts

import (
	"SuBlog/models"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/nicklaw5/go-respond"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

func Create(w http.ResponseWriter, r *http.Request, database *mongo.Database) {
	collection := database.Collection("posts")

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	post := &models.Post{
		ID:        primitive.NewObjectID(),
		Title:     r.FormValue("title"),
		Content:   r.FormValue("content"),
		Author:    primitive.ObjectID{},
		Upvotes:   0,
		Downvotes: 0,
		Comments:  nil,
		Timestamp: time.Now().Unix(),
	}

	err = validator.New().Struct(post)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
		return
	}

	_, err = collection.InsertOne(context.Background(), post)
	if err != nil {
		return
	}

	respond.NewResponse(w).Ok(struct {
		ID string `json:"id"`
	}{
		ID: post.ID.Hex(),
	})
}
