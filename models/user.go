package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Modelo de usuario
type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"name" json:"name,omitempty" valid:"alphanum,required"`
	Apellidos       string             `bson:"last_name" json:"last_name,omitempty" valid:"alphanum,required" `
	FechaNacimiento time.Time          `bson:"nacimiento" json:"nacimienti,omitempty" valid:"required"`
	Email           string             `bson:"email" json:"email" valid:"email,required"`
	Password        string             `bson:"password" json:"password,omitempty" valid:"length(6|32),required" `
	Avater          string             `bson:"avatar" json:"avatar,omitempty" valid:"optional"`
	Banner          string             `bson:"banner" json:"banner,omitempty" valid:"optional"`
	Biografia       string             `bson:"biographic" json:"biographic,omitempty" valid:"optional"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omitempty" valid:"optional"`
	SitioWeb        string             `bson:"sitio_web" json:"sitio_web,omitempty valid:"optional"`
}
