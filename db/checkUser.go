package db

import (
	"context"
	"time"

	"github.com/BonifacioJZ/tiwgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

//Confirmar si el usuario existe
func CheckUser(email string) (models.User, bool, string) {
	cxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tiwgo")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(cxt, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
