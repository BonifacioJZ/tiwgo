package main

import (
	"log"

	"github.com/BonifacioJZ/tiwgo/db"
	"github.com/BonifacioJZ/tiwgo/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la DB")
	}
	handlers.Handlers()
}
