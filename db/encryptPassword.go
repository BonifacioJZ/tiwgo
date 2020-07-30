package db

import "golang.org/x/crypto/bcrypt"

//funcion para encriptar el password
func EncryptPassword(pass string) (string, error) {

	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
