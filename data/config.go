package data

import (
	"fmt"
	"imagetopdf/models"
	"log"
	"strings"

	"encoding/json"
	"os"
)

var Config models.ConfigModel = GetConfig()

func GetConfig() models.ConfigModel {
	configuration := models.ConfigModel{}

	dir, err := os.Getwd()

	env := strings.ToLower(os.Getenv("APP_ENV"))

	var filename string = ""

	switch env {
	case "dev":
		filename = "config.dev"

	case "prod":
		configuration.Host = fmt.Sprintf(":%s", os.Getenv("PORT"))
		configuration.RedisUrl = os.Getenv("REDIS_URL")
		configuration.StoragePath = os.Getenv("STORAGE_PATH")

	default:
		panic(1)
	}

	fmt.Printf("Env: %s - Filename: %s \n", env, filename)

	path := fmt.Sprintf("%s/configs/%s.json", dir, filename)

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	file, _ := os.Open(path)
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&configuration)

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	return configuration
}
