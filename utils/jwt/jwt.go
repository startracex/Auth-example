package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/startracex/Auth-example/conf"
	"reflect"
	"time"
)

var host = conf.Global.Host
var name = conf.Global.Name

var Exp = 24 * 30 * time.Hour

type MapClaims jwt.MapClaims

func Get(claims jwt.Claims, query string) (any, bool) {
	c := claims.(jwt.MapClaims)
	v, y := c[query]
	return v, y
}

func GetString(claims jwt.Claims, query string) (string, bool) {
	v, y := Get(claims, query)
	return v.(string), y
}

func ToJWT(structer any) jwt.MapClaims {
	v := reflect.ValueOf(structer)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	mapper := make(map[string]any)
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		name := fi.Name
		if tag := fi.Tag.Get("json"); tag != "" {
			name = tag
		}
		mapper[name] = v.Field(i)
	}
	return JWT(mapper)
}

func JWT(mapper map[string]any) jwt.MapClaims {
	c := jwt.MapClaims{}
	for k, v := range mapper {
		c[k] = v
	}
	if _, y := c["iss"]; !y {
		c["iss"] = "127.0.0.1"
	}
	if _, y := c["iad"]; !y {
		c["iad"] = time.Now().Unix()
	}
	if _, y := c["exp"]; !y {
		c["exp"] = time.Now().Add(Exp).Unix()
	}
	return c
}

func Parse(tokenString string, key any) (*jwt.Token, error) {
	if str, y := key.(string); y {
		key = []byte(str)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func Sign(token *jwt.Token, key any) (string, error) {
	if str, y := key.(string); y {
		key = []byte(str)
	}
	return token.SignedString(key)
}

func Sign256(claims jwt.MapClaims, key any) (string, error) {
	method := jwt.SigningMethodHS256
	token := &jwt.Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": method.Alg(),
		},
		Claims: claims,
		Method: method,
	}
	return Sign(token, key)
}

func Sign384(claims jwt.MapClaims, key any) (string, error) {
	method := jwt.SigningMethodHS384
	token := &jwt.Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": method.Alg(),
		},
		Claims: claims,
		Method: method,
	}
	return Sign(token, key)
}

func Sign512(claims jwt.MapClaims, key any) (string, error) {
	method := jwt.SigningMethodHS512
	token := &jwt.Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": method.Alg(),
		},
		Claims: claims,
		Method: method,
	}
	return Sign(token, key)
}
