package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/BonifacioJZ/tiwgo/models"
)

func UpdateUser(user models.User, Id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tiwgo")
	col := db.Collection("users")

	registro := make(map[string]interface{})
	if len(user.Nombre) > 0 {
		registro["nombre"] = user.Nombre
	}
	if len(user.Apellidos) > 0 {
		registro["apellidos"] = user.Apellidos
	}
	registro["nacimiento"] = user.FechaNacimiento
	if len(user.Avater) > 0 {
		registro["avatar"] = user.Avater
	}
	if len(user.Biografia) > 0 {
		registro["biographic"] = user.Biografia
	}
	if len(user.Ubicacion) > 0 {
		registro["ubicacion"] = user.Ubicacion
	}
	if len(user.SitioWeb) > 0 {
		registro["sitio"] = user.SitioWeb
	}

	updateString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(Id)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}
	return true, nil

}
