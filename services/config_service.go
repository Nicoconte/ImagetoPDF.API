package services

import (
	"fmt"
	"imagetopdf/models"
	"log"

	"encoding/json"
	"os"
)

func GetConfig() models.ConfigModel {

	dir, err := os.Getwd()

	path := fmt.Sprintf("%s/config.json", dir)

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	file, _ := os.Open(path)
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := models.ConfigModel{}

	err = decoder.Decode(&configuration)

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	return configuration
}
