package db

import (
	"github.com/BonifacioJZ/tiwgo/models"
	"golang.org/x/crypto/bcrypt"
)

/*funcion para realizar el login*/
func Login(email string, password string) (models.User, bool) {

	user, find, _ := CheckUser(email)

	if find == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return user, false
	}
	return user, true

}
