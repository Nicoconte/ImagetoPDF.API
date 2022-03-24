package services

import (
	"errors"
	"fmt"
	"imagetopdf/helpers"
	"mime/multipart"
	"os"
	"strings"
)

var BaseStorageRoute string = Config.StoragePath

var AllowedExtensions map[string]bool = Config.AllowedExtensions

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

func DeleteImageFromStorage(filename string) (bool, error) {
	path := BaseStorageRoute + filename

	err := os.Remove(path)

	if err != nil {
		return false, err
	}

	return true, nil
}
