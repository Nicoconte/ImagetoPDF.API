package services

import (
	"errors"
	"fmt"
	"imagetopdf/helpers"
	"mime/multipart"
	"strings"
)

//Temp VAR
var BaseStorageRoute string = "C:/GIT/me/ImagetoPDF/ImagetoPDF.Storage/"

var AllowedExtensions = map[string]bool{
	"jpg":  true,
	"png":  true,
	"jpeg": true,
}

func SaveImagesIntoStorage(files []*multipart.FileHeader) (bool, error) {

	for _, file := range files {
		output := BaseStorageRoute + file.Filename

		extensionParts := strings.Split(file.Filename, ".")

		extension := extensionParts[len(extensionParts)-1]

		if !AllowedExtensions[extension] {
			return false, errors.New(fmt.Sprintf("Extension %s is not allowed", extension))
		}

		if err := helpers.CreateFileFromRequestHeader(file, output); err != nil {
			return false, err
		}
	}

	return true, nil
}
