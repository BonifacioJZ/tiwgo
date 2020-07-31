package routers

import (
	"encoding/json"
	"net/http"

	"github.com/BonifacioJZ/tiwgo/db"
	"github.com/BonifacioJZ/tiwgo/models"
	"github.com/asaskevich/govalidator"
)

//Registro para la creacion de usuario
func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error en los datos"+err.Error(), 400)
		return
	}

	_, e := govalidator.ValidateStruct(user)

	if e != nil {
		http.Error(w, "Error en la validacion"+e.Error(), 400)
		return
	}

	_, find, _ := db.CheckUser(user.Email)

	if find == true {
		http.Error(w, "Ya existe un usuario registrado con el email", 400)
	}

	_, status, err := db.InsertUser(user)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar"+err.Error(), 400)
	}
	if status == false {
		http.Error(w, "No se ha logrado la insersion", 400)
	}

	w.WriteHeader(http.StatusCreated)

}
