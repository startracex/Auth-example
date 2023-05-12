package auth

import (
	"main/db"

	"go.mongodb.org/mongo-driver/bson"
)

type BindAuth struct {
	Email    string `json:"email" form:"email"`
	E_mail   string `json:"e-mail" form:"e-mail"`
	Passwd   string `json:"passwd" form:"passwd"`
	Password string `json:"password" form:"password"`
}
type Auth struct {
	Email    string
	Password string
}

// 将BindAuth转换为Auth
func (ba *BindAuth) ToAuth() Auth {
	// 将非空的字段赋值给Auth
	var auth Auth
	if ba.E_mail != "" {
		auth.Email = ba.E_mail
	}
	if ba.Email != "" {
		auth.Email = ba.Email
	}
	if ba.Password != "" {
		auth.Password = ba.Password
	}
	if ba.Passwd != "" {
		auth.Password = ba.Passwd
	}
	return auth
}

// 返回Auth的账户_id,如果没有则返回空
func (a Auth) GetID() string {
	return db.LoginDatabase.GetID(bson.M{"email": a.Email, "password": a.Password})
}

// 返回Auth的存在性
func (a Auth) UserExist() bool {
	return len(db.RegisterDatabase.Find(bson.M{"email": a.Email})) != 0
}

// 注册Auth到数据库
func (a Auth) UserRegister() {
	db.RegisterDatabase.Insert(bson.M{"email": a.Email, "password": a.Password})
}

// 返回Auth的账户信息
func (a Auth) UserLogin() map[string]any {
	return db.LoginDatabase.Find(bson.M{"email": a.Email, "password": a.Password})
}
