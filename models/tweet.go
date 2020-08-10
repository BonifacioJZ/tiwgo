package models

//Caṕutuara el body del tweet
type Tweet struct {
	Message string `bson:"message" json:"message" valid:"required, stringlength(0|36)"`
}
