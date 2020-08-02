package models

//Respuesta del login
type RespLogin struct {
	Token string `json:"token,omitempty"`
}
