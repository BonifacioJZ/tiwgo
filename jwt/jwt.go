package jwt

import (
	"time"

	"github.com/BonifacioJZ/tiwgo/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//Funcion para generar el JWT
func GeneraJWT(user models.User) (string, error) {

	myClave := []byte("MasterdelDesarrollo_grupodeFacebook")

	payload := jwt.MapClaims{
		"email":            user.Email,
		"nombre":           user.Nombre,
		"fecha_nacimeinto": user.FechaNacimiento,
		"biografia":        user.Biografia,
		"ubicacion":        user.Ubicacion,
		"sitioweb":         user.SitioWeb,
		"_id":              user.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
