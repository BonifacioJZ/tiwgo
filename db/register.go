package db

import (
	"context"
	"time"
	"github.com/BonifacioJZ/tiwgo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Registrar usuario
func InsertUser(u models.User)(string, bool,error){
	
	ctx, cancel := context.WithTimeout(context.Background(),15*time.Second)
	
	defer cancel()
	
	db := MongoCN.Database("tiwgo")
	col := db.Collection("users")

	u.Password,_ :=EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx,u)
	if err != nil{
		return "", false, err
	}

	objID,_ := result.InsertedID.(primitive.ObjectID)
	return objID.String(),true,nil
}