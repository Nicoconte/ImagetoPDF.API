package services

import (
	"errors"
	"fmt"
	"imagetopdf/helpers"
	"imagetopdf/models"
	"mime/multipart"
	"strings"
)

var config models.ConfigModel = GetConfig()

var BaseStorageRoute string = config.StoragePath

var AllowedExtensions = config.AllowedExtensions

func SaveImagesIntoStorage(files []*multipart.FileHeader) (bool, error) {

	for _, file := range files {
		output := BaseStorageRoute + file.Filename

		fileParts := strings.Split(file.Filename, ".")

		extension := fileParts[len(fileParts)-1]

		if !AllowedExtensions[extension] {
			return false, errors.New(fmt.Sprintf("Extension .%s is not allowed", extension))
		}

		if err := helpers.CreateFileFromRequestHeader(file, output); err != nil {
			return false, err
		}
	}

	return true, nil
}
