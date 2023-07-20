package jwt

import (
	"main/conf"
	"time"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var host = conf.Global.Host
var name = conf.Global.Name

type WithPaylod struct {
	Object_ID primitive.ObjectID `json:"_id"`
	jwt.StandardClaims
}

const ext = time.Hour * 24 * 30

func JWTstd() jwt.StandardClaims {
	return jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ext).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    host,
		Subject:   name,
	}
}

func JWTmap() jwt.MapClaims {
	return jwt.MapClaims{
		"exp": time.Now().Add(ext).Unix(),
		"iat": time.Now().Unix(),
		"iss": host,
		"sub": name,
	}
}

func MapToSTD(claims jwt.MapClaims) jwt.StandardClaims {
	std := JWTstd()
	std.ExpiresAt = claims["exp"].(int64)
	std.IssuedAt = claims["iat"].(int64)
	std.Issuer = claims["iss"].(string)
	std.Subject = claims["sub"].(string)
	return std
}

func STDToMap(std jwt.StandardClaims) jwt.MapClaims {
	claims := JWTmap()
	claims["exp"] = std.ExpiresAt
	claims["iat"] = std.IssuedAt
	claims["iss"] = std.Issuer
	claims["sub"] = std.Subject
	return claims
}

func From(customClaims WithPaylod) jwt.Claims {
	std := JWTstd()
	customClaims.StandardClaims = std
	return customClaims
}

func Map(param map[string]interface{}) jwt.Claims {
	claims := JWTmap()
	for k, v := range param {
		claims[k] = v
	}
	return claims
}

func Parse(toparse string, key string) (*WithPaylod, *jwt.Token, error) {
	p := &WithPaylod{}
	t, e := jwt.ParseWithClaims(toparse, p, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	return p, t, e
}

func Sign256(claims jwt.Claims, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

func Sign385(claims jwt.Claims, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	return token.SignedString([]byte(key))
}

func Sign512(claims jwt.Claims, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(key))
}
