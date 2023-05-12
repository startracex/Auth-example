package config

import (
	"os"
	"strings"
)

type Argp struct {
	args   []string
	origin []string
}

// Equal sign
var Eq = "="

// Attach sign
var At = "--"

// Set the args and origin from origin and return the Argp setted
func (X *Argp) From(origin []string) *Argp {
	X.origin = origin
	X.args = origin
	return &Argp{
		args:   origin,
		origin: origin,
	}
}

// Set the args and origin from os.Args and return the Argp setted
func (X *Argp) New() *Argp {
	return X.From(os.Args)
}

// Query finds in the args and return it's existence
// The default value must be false
func (X *Argp) Bool(finds ...string) bool {
	for _, find := range finds {
		for i, v := range X.args {
			if v == find {
				X.args = append(X.args[:i], X.args[i+1:]...)
				return true
			}
		}
	}
	return false
}

// Query finds in the args and set the exist result to value
func (X *Argp) BoolVar(value *bool, finds ...string) {
	for _, find := range finds {
		*value = X.Bool(find)
	}
}

// Query finds in the args and return the result and true or "" and false
func (X *Argp) String(finds ...string) (value string, exist bool) {
	for _, find := range finds {
		for i, v := range X.args {
			if v == find && i+1 < len(X.args) {
				// 删除相同值和之后的值
				value = X.args[i+1]
				if value == Eq && i+2 < len(X.args) {
					value = X.args[i+2]
					X.args = append(X.args[:i], X.args[i+1:]...)
				}
				X.args = append(X.args[:i], X.args[i+2:]...)
				return value, true
			} else if strings.HasPrefix(v, find+Eq) && len(v) > len(find) {
				value = v[len(find)+1:]
				X.args = append(X.args[:i], X.args[i+1:]...)
				return value, true
			}
		}
	}
	return "", false
}

// Query finds in the args and set the parsed result to value
func (X *Argp) StringVar(value *string, finds ...string) {
	for _, find := range finds {
		v, e := X.String(find)
		if e {
			*value = v
		}
	}
}

// Query the args which start with find and return the result and true or "" and false
func (X *Argp) Start(find string) (string, bool) {
	for i, v := range X.args {
		if strings.HasPrefix(v, find) && len(v) > len(find) {
			X.args = append(X.args[:i], X.args[i+1:]...)
			return v[strings.Index(v, find):], true
		}
	}
	return "", false
}

// Return the args after the "--"
func (X *Argp) Attach() []string {
	for i, v := range X.args {
		if v == At {
			value := X.args[i+1:]
			X.args = X.args[:i]
			return value
		}
	}
	return []string{}
}
