package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// mongodb connection.
type Book struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Author    string             `json:"author" bson:"author,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// DeleteRequest struct is used to parse Delete Requests for Books
type DeleteRequest struct {
	ID string `json:"id"`
}
