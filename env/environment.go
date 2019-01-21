package env

import (
	"os"
)

func GetVar(name, defaultValue string) string {
	v := os.Getenv(name)
	if v == "" {
		v = defaultValue
	}
	return v
}

func SetVar(name, value string) string {
	v := os.Getenv(name)
	if err := os.Setenv(name, value); err != nil {
		panic(err)
	}
	return v
}
