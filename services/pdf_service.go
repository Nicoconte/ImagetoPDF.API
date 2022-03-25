package services

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func GeneratePDF(foldername string, outputFilename string) (string, error) {
	cmdStr := buildPdfCpuCommand(foldername, outputFilename)

	log.Println(cmdStr)

	cmd := exec.Command("powershell", cmdStr)
	err := cmd.Run()

	if err != nil {
		log.Println(fmt.Printf("Cannot proccess command: Reason: %s", err.Error()))
		return "", err
	}

	outputFilePath := Config.StoragePath + "/" + foldername + "/output/" + outputFilename

	return outputFilePath, nil
}

func buildPdfCpuCommand(inputFolder string, outputFilename string) string {
	path := Config.StoragePath + inputFolder + "/"
	imagesFromStorage, err := os.ReadDir(path)

	filters := ""

	if err != nil {
		log.Printf("Cannot files from dir => %s .Reason: %s", path, err.Error())

		return ""
	}

	imageNames := ""

	for _, image := range imagesFromStorage {

		imageNameParts := strings.Split(image.Name(), ".")

		extension := imageNameParts[len(imageNameParts)-1]

		if Config.AllowedExtensions[extension] {
			imageNames += fmt.Sprintf("%s ", path+image.Name())
		}
	}

	outputPath := path + "output"

	cmd := fmt.Sprintf("pdfcpu import %s %s/%s %s", filters, outputPath, outputFilename, imageNames)

	return cmd
}

func DeletePDFFromStorage(pdfPath string) error {
	err := os.RemoveAll(pdfPath)

	if err != nil {
		return err
	}

	return nil
}
