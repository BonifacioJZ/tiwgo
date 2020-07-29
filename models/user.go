package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Modelo de usuario
type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"name" json:"name,omitempty"`
	Apellidos       string             `bson:"last_name" json:"last_name,omitempty"`
	FechaNacimiento time.Time          `bson:"nacimiento" json:"nacimienti,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avater          string             `bson:"avatar" json:"avatar,omitempty"`
	Banner          string             `bson:"banner" json:"banner,omitempty"`
	Biografia       string             `bson:"biographic" json:"biographic,omitempty"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb        string             `bson:"sitio_web" json:"sitio_web,omitempty"`
}
