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

	dir, err := os.Getwd()

	env := strings.ToLower(os.Getenv("ImagetopdfEnv"))

	var filename string = ""

	if env == "docker" {
		filename = "config.docker"
	} else if env == "local" {
		filename = "config.local"
	} else {
		panic(1)
	}

	path := fmt.Sprintf("%s/configs/%s.json", dir, filename)

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
