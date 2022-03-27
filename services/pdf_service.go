package services

import (
	"fmt"
	"imagetopdf/data"
	"log"
	"os"
	"os/exec"
	"strings"
)

var storageBasePath string = data.Config.StoragePath

func DeletePDFFromStorage(pdfPath string) error {
	err := os.RemoveAll(pdfPath)

	if err != nil {
		return err
	}

	return nil
}

func GeneratePDF(foldername string, outputFilename string) (string, error) {
	cmd, err := buildPdfCpuCommand(foldername, outputFilename)

	if err != nil {
		log.Println(fmt.Printf("Cannot build command: Reason: %s", err.Error()))
		return "", err
	}

	fmt.Printf("Exec: %s \n", data.Config.CommandExecutor)

	err = exec.Command(data.Config.CommandExecutor, cmd).Run()

	if err != nil {
		log.Println(fmt.Printf("Cannot proccess command: Reason: %s", err.Error()))
		return "", err
	}

	outputFilePath := storageBasePath + "/" + foldername + "/output/" + outputFilename

	return outputFilePath, nil
}

func buildPdfCpuCommand(inputFolder string, outputFilename string) (string, error) {
	path := storageBasePath + inputFolder + "/"

	imagesFromStorage, err := getImagesPathFromStorage(path)

	if err != nil {
		log.Printf("Cannot get images from storage. Reason: %s", err.Error())
		return "", err
	}

	imageNames := strings.Join(imagesFromStorage, " ")

	outputPath := path + "output"

	cmd := fmt.Sprintf("pdfcpu import %s/%s %s", outputPath, outputFilename, imageNames)

	return cmd, nil
}

func getImagesPathFromStorage(basePath string) ([]string, error) {
	imagesFromStorage, err := os.ReadDir(basePath)

	if err != nil {
		log.Printf("Cannot files from dir => %s .Reason: %s", basePath, err.Error())
		return nil, err
	}

	var imagesPath []string

	for _, image := range imagesFromStorage {

		imageNameParts := strings.Split(image.Name(), ".")

		extension := imageNameParts[len(imageNameParts)-1]

		if data.Config.AllowedExtensions[extension] {
			imagesPath = append(imagesPath, fmt.Sprintf("%s ", basePath+"/"+image.Name()))
		}
	}

	return imagesPath, nil
}

func DeleteAllImagesAfterDownload(foldername string) {
	paths, err := getImagesPathFromStorage(storageBasePath + foldername)

	if err != nil {
		log.Fatalf("Cannot delete image. Error: %s", err.Error())
	}

	for _, path := range paths {
		fmt.Println(path)
		os.Remove(path)
	}

}
