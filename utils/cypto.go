package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/sony/sonyflake"
)

func GenUUID() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, _ := flake.NextID()
	return Md5(fmt.Sprintf("%d", id))
}

func Sha1(origin string) string {
	o := sha1.New()
	o.Write([]byte(origin))
	return hex.EncodeToString(o.Sum(nil))
}

func Md5(origin string) string {
	o := md5.New()
	o.Write([]byte(origin))
	return hex.EncodeToString(o.Sum(nil))
}
