package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/sony/sonyflake"
)

// Generate random identifier
func GenUUID() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, _ := flake.NextID()
	return fmt.Sprintf("%d", id)
}

// To sha1
func Sha1(origin string) string {
	o := sha1.New()
	o.Write([]byte(origin))
	return fmt.Sprintf("%x", o.Sum(nil))
}

// To md5
func Md5(origin string) string {
	o := md5.New()
	o.Write([]byte(origin))
	return fmt.Sprintf("%x", o.Sum(nil))
}

// To sha256
func Sha256(origin string) string {
	o := sha256.New()
	o.Write([]byte(origin))
	return fmt.Sprintf("%x", o.Sum(nil))
}

// To base64
func Base64(origin string) string {
	o := base64.URLEncoding.EncodeToString([]byte(origin))
	return o
}

// Ascii to base64
func AtoB(origin string) string {
	return Base64(origin)
}

// Base64 to ascii
func BtoA(origin string) string {
	o, _ := base64.URLEncoding.DecodeString(origin)
	return string(o)
}
