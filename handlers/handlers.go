package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/BonifacioJZ/tiwgo/middlew"
	"github.com/BonifacioJZ/tiwgo/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Handler seteo de puertos
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
