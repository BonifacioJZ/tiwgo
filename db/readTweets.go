package db

import (
	"context"
	"log"
	"time"

	"github.com/BonifacioJZ/tiwgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Funcion para regresar los tweets
func ReadTweets(ID string, page int64) ([]*models.ReadTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("tiwgo")
	col := db.Collection("twits")

	var result []*models.ReadTweets

	condition := bson.M{
		"usserid": ID,
	}

	option := options.Find()
	option.SetLimit(20)
	option.SetSort(bson.D{{Key: "fecha", Value: -1}})
	option.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, option)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	for cursor.Next(context.TODO()) {

		var register models.ReadTweets
		err := cursor.Decode(&register)
		if err != nil {
			return result, false
		}
		result = append(result, &register)
	}
	return result, true

}
