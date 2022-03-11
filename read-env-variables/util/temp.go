package util

import (
	"os"
	// _ "github.com/joho/godotenv/autoload"
)

func GetVal() string {
	something := os.Getenv("SOMETHING")
	return something
}
