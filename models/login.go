package models

type Login struct {
	Email    string `valid:"email,required"`
	Password string `valid:"required"`
}
