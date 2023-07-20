package test

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"testing"
)

func Sha1(origin string) string {
	o := sha1.New()
	o.Write([]byte(origin))
	return fmt.Sprintf("%x", o.Sum(nil))

}

func Md5(origin string) string {
	o := md5.New()
	o.Write([]byte(origin))
	return fmt.Sprintf("%x", o.Sum(nil))

}

func Sha256(origin string) string {
	o := sha256.New()
	o.Write([]byte(origin))
	return fmt.Sprintf("%x", o.Sum(nil))
}

func Sha512(origin string) string {
	o := sha512.New()
	o.Write([]byte(origin))
	return fmt.Sprintf("%x", o.Sum(nil))
}
func Base64(origin string) string {
	o := base64.URLEncoding.EncodeToString([]byte(origin))
	return o
}
func Base64Decode(origin string) string {
	o, _ := base64.URLEncoding.DecodeString(origin)
	return string(o)
}
func TestXxx(t *testing.T) {
	fmt.Println(Sha1("123456"))
	fmt.Println(Md5("123456"))
	fmt.Println(Sha256("123456"))
	fmt.Println(Sha512("123456"))
	fmt.Println(Base64("123456"))
	fmt.Println(Base64Decode("MTIzNDU2"))

}
func TestMap(t *testing.T) {
	m := map[string]interface{}{"a": 1, "b": 2}
	fmt.Println(m["c"])
}
func Testu(t *testing.T) {

}
