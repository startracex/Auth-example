package auth

import (
	"main/db"
	"go.mongodb.org/mongo-driver/bson"
)

func (a Auth) UserExist() bool {
	// TODO IF EXISTS
	return len(db.RegisterDatabase.Find(bson.M{"email": a.Email})) != 0
}

func (a Auth) UserRegister() {
	// TODO TO REGISTER
	db.RegisterDatabase.Insert(bson.M{"email": a.Email, "password": a.Password})
}

func (a Auth) UserLogin() map[string]any {
	// TODO TO LOGIN
	return db.LoginDatabase.Find(bson.M{"email": a.Email, "password": a.Password})
}
