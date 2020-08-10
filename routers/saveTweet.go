package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"

	"github.com/BonifacioJZ/tiwgo/db"
	"github.com/BonifacioJZ/tiwgo/models"
)

//Guardar la informacion en la base de datos
func SaveTweet(w http.ResponseWriter, r *http.Request) {

	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		http.Error(w, "Error al decodificar los datos "+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = govalidator.ValidateStruct(message)

	if err != nil {
		http.Error(w, "Ocurrio un error con los datos "+err.Error(), http.StatusBadRequest)
		return
	}

	register := models.SaveTweet{
		UserID:  IdUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, err = govalidator.ValidateStruct(register)
	if err != nil {
		http.Error(w, "Ocurri un error en la incersion del los datos "+err.Error(), http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertTweet(register)

	if err != nil {
		http.Error(w, "Ocuccio un error al insertar los datos "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se insetaron los datos ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
