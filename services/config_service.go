package services

import (
	"fmt"
	"imagetopdf/models"
	"log"

	"encoding/json"
	"os"
)

var Config models.ConfigModel = GetConfig()

func GetConfig() models.ConfigModel {

	dir, err := os.Getwd()

	path := fmt.Sprintf("%s/configs/config.json", dir)

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
