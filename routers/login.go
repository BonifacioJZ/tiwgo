package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BonifacioJZ/tiwgo/db"
	"github.com/BonifacioJZ/tiwgo/jwt"
	"github.com/BonifacioJZ/tiwgo/models"
	"github.com/asaskevich/govalidator"
)

//funcion para ejecutar el login
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")

	var user models.Login

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {

		http.Error(w, "Email y/o Contraseña invalidos"+err.Error(), 400)
		return
	}

	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		http.Error(w, "Ocurrio un error al procesar los datos"+err.Error(), 400)
		return
	}
	document, exist := db.Login(user.Email, user.Password)
	if exist == false {
		http.Error(w, "El usuario y/o Contraseña son Invalidos", 400)
		return
	}

	jwtKey, err := jwt.GeneraJWT(document)
	if err != nil {
		http.Error(w, "Ocurrio un error al generar al token"+err.Error(), 400)
		return
	}
	resp := models.RespLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
