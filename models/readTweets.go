package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//modelos de read tweets
type ReadTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userId" json:"userId,omitempty"`
	Message string             `bson:"mensaje" json:"mensaje,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}
