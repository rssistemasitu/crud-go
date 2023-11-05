package utils

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "crud-go"

func GetEnvVariable(logName string) string {
	// Garantir que busque o .env independente de onde chamar
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		result := fmt.Sprintf("Error loading .env file to get %s", logName)
		log.Fatal(result)
	}

	return os.Getenv(logName)
}
