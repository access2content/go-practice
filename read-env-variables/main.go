package main

import (
	"fmt"
	"os"

	"github.com/access2content/go-practice/read-env-variables/util"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("HTTP_PORT")
	fmt.Println("Port = ", port)

	value := util.GetVal()
	fmt.Println("Value = ", value)
}
