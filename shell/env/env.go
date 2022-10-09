package env

import (
	"strings"
	"syscall"
)

var (
	env map[string]string = map[string]string{}
)

func LoadEnv() {
	for _, pair := range syscall.Environ() {
		kv := strings.SplitN(pair, "=", 2)
		key, val := kv[0], kv[1]

		env[key] = val
	}
}

func Get(key string) string {
	v, ok := env[key]
	if !ok {
		return ""
	}
	return v
}

func Set(key, value string) {
	env[key] = value
}
