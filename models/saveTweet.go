package models

import (
	"time"
)

//Estructuta para guardar tweet
type SaveTweet struct {
	UserID  string    `bson:"userid" json:"userid,omitempty" valid:"required"`
	Message string    `bson:"message" json:"message,omitempty" valid:"required,stringlength(0|36)"`
	Date    time.Time `bson:"date" json:"date,omitempty" valid:"required"`
}
