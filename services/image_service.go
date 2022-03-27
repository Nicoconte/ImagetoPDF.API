package services

import (
	"errors"
	"fmt"
	"imagetopdf/data"
	"imagetopdf/helpers"
	"mime/multipart"
	"os"
	"strings"
)

var BaseStorageRoute string = data.Config.StoragePath

var AllowedExtensions map[string]bool = data.Config.AllowedExtensions

func SaveImagesIntoStorage(files []*multipart.FileHeader, foldername string) (bool, error) {

	for _, file := range files {
		output := BaseStorageRoute + foldername + "/" + file.Filename

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

func DeleteImageFromStorage(filename string, foldername string) (bool, error) {
	path := BaseStorageRoute + foldername + "/" + filename

	err := os.Remove(path)

	if err != nil {
		return false, err
	}

	return true, nil
}
