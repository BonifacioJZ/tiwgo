package routers

import (
	"encoding/json"
	"net/http"

	"github.com/BonifacioJZ/tiwgo/db"

	"github.com/BonifacioJZ/tiwgo/models"
)

//Modificar la informacion del usuario
func UpdateProfile(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Datos Incorrectos"+err.Error(), 400)
		return
	}

	var status bool
	status, err = db.UpdateUser(user, IdUser)

	if err != nil {
		http.Error(w, "Ocurrio un error al modificar el registro"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se Modificaron los datos"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
