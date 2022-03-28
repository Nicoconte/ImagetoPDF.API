package services

import (
	"errors"
	"fmt"
	Image "image"
	"imagetopdf/data"
	"imagetopdf/helpers"
	"imagetopdf/models"
	"io/fs"
	"log"
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

func GetImagesFromStorage(basePath string) ([]models.ImageModel, error) {
	imageEntries, err := os.ReadDir(basePath)

	if err != nil {
		log.Printf("Cannot files from dir => %s .Reason: %s", basePath, err.Error())
		return nil, err
	}

	var images []models.ImageModel

	for _, imageEntry := range imageEntries {

		if imageEntry.IsDir() {
			continue
		}

		img, err := GetImageInformation(imageEntry, basePath)

		if err != nil {
			fmt.Printf("Cannot get image information. Reason %s \n", err.Error())
			return nil, err
		}

		if data.Config.AllowedExtensions[img.Extension] {
			images = append(images, img)
		}
	}

	if len(images) == 0 {
		return nil, errors.New("Directory is empty")
	}

	return images, nil
}

func GetImageInformation(imageEntry fs.DirEntry, basePath string) (models.ImageModel, error) {
	img := models.ImageModel{}

	imgFullpath := fmt.Sprintf("%s/%s", basePath, imageEntry.Name())

	reader, err := os.Open(imgFullpath)

	if err != nil {
		fmt.Printf("Cannot get image information. Reason: %s\n", err.Error())
		return img, err
	}

	defer reader.Close()

	imgDecoded, _, err := Image.DecodeConfig(reader)

	if err != nil {
		fmt.Printf("Cannot decode image. Name: %s - Reason: %s\n", imageEntry.Name(), err.Error())
		return img, err
	}

	imgNameParts := strings.Split(imageEntry.Name(), ".")
	extension := imgNameParts[len(imgNameParts)-1]

	img.Extension = extension
	img.Name = imageEntry.Name()
	img.Width = float64(imgDecoded.Width)
	img.Height = float64(imgDecoded.Height)
	img.Path = imgFullpath

	return img, nil
}

func DeleteAllImagesAfterDownload(foldername string) {
	images, err := GetImagesFromStorage(storageBasePath + foldername)

	if err != nil {
		log.Fatalf("Cannot delete image. Error: %s", err.Error())
	}

	for _, img := range images {
		os.Remove(img.Path)
	}

}
