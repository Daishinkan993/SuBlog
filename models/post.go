package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Title     string             `bson:"title" json:"title" validate:"required,min=3,max=80"`
	Content   string             `bson:"content" json:"content" validate:"required,min=10"`
	Author    primitive.ObjectID `bson:"author" json:"author"`
	Upvotes   int                `bson:"upvotes" json:"upvotes"`
	Downvotes int                `bson:"downvotes" json:"downvotes"`
	ViewCount int                `bson:"viewCount" json:"viewCount"`
	Comments  []string           `bson:"comments" json:"comments"`
	Timestamp int64              `bson:"timestamp" json:"timestamp"`
}
