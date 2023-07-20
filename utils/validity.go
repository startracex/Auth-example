package utils

import (
	"regexp"
	"strings"
)

func CheckEmail(origin string) bool {
	reg := regexp.MustCompile(`^.+@.+\..+$`)
	return isMatch(reg, origin)
}

func CheckPassword(origin string) bool {
	origin = strings.ToLower(origin)
	reg := regexp.MustCompile(`^[0-9a-z]+$`)
	return isMatch(reg, origin)
}

func isMatch(reg *regexp.Regexp, origin string) bool {
	return reg.MatchString(strings.ToLower(origin))
}
