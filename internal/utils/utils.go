package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(logName string) string {
	err := godotenv.Load("../configs/.env")
	if err != nil {
		err = godotenv.Load("configs/.env")
	}

	if err != nil {
		result := fmt.Sprintf("Error loading .env file to get %s", logName)
		log.Fatal(result)
	}
	return os.Getenv(logName)
}
