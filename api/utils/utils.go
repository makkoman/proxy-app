package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	err := godotenv.Load("/Users/makko/go/src/proxy-app/.env")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(os.Getenv("PORT"))
}