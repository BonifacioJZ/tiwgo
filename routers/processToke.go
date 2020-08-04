package routers

import (
	"errors"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/BonifacioJZ/tiwgo/db"
	"github.com/BonifacioJZ/tiwgo/models"
)

var Email string

var IdUser string
//procesar el token
func ProcessToken(tkn string) (*models.Claim,bool,string,error) {
	myClave := []byte("MasterdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	sliptToken := strings.Split(tkn,"Bearer")
	if len(sliptToken) != 2{
		return claims,false,string(""),errors.New("formato de token invalido")
	}
	
	tkn = strings.TrimSpace(sliptToken[1])

	tk,err := jwt.ParseWithClaims(tkn,claims,func(token *jwt.Token)(interface{},error){
		return myClave,nil
	})
	if err == nil{
		_,find, _:= db.CheckUser(claims.Email)
		if find == true{
			Email = claims.Email
			IdUser = claims.ID.Hex()
		}
		return claims,find,IdUser,nil
	}
	if !tk.Valid{
		return claims,false,string(""),errors.New("token Invalido")
	}
	return claims, false,string(""),err
}