package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/BonifacioJZ/tiwgo/models"

	"go.mongodb.org/mongo-driver/bson"
)

//Funcion para insertar el tweet
func InsertTweet(twit models.SaveTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tiwgo")
	col := db.Collection("twits")

	register := bson.M{
		"usserid": twit.UserID,
		"mensaje": twit.Message,
		"fecha":   twit.Date,
	}

	result, err := col.InsertOne(ctx, register)

	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil

}
