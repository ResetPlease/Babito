package tools

import (
	"fmt"
	"os"
)

func GetenvWithPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("empty value for key=%s", key))
	}
	return value
}
