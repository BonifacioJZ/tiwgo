package db

import (
	"context"
	"fmt"
	"time"

	"github.com/BonifacioJZ/tiwgo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Funcion para buscar perfil
func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("tiwgo")
	col := db.Collection("users")

	var profile models.User
	ojbId, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": ojbId,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("registro no encontrado" + err.Error())
		return profile, err
	}
	return profile, nil

}
