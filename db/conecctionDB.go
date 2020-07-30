package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el objeto de la conexion ala base de datos
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://reb_jz:fOsAtOTHcC8NXEqW@reb.iefiq.azure.mongodb.net/tiwgo?retryWrites=true&w=majority"))

/*Es una funcion para conectar a la base de datos*/
func ConnectDB() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Println("Conexion Existosa con la BD")
	return client

}

/*Es un  check a la base de datos*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1
}
