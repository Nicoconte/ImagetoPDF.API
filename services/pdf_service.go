package services

import (
	"fmt"
	"imagetopdf/data"
	"os"

	"github.com/signintech/gopdf"
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

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	path := storageBasePath + foldername + "/"

	imagesFromStorage, err := GetImagesFromStorage(path)

	if err != nil {
		fmt.Printf("Cannot generate pdf: Reason %s\n", err.Error())
		return "", err
	}

	for _, img := range imagesFromStorage {

		pageDimensions := &gopdf.Rect{
			W: img.Width,
			H: img.Height,
		}

		//Page with size base on img w and h
		pdf.AddPageWithOption(gopdf.PageOption{
			PageSize: pageDimensions,
		})

		pdf.Image(img.Path, 0.0, 0.0, pageDimensions)
	}

	outputPath := fmt.Sprintf("%soutput/%s", path, outputFilename)

	err = pdf.WritePdf(outputPath)

	if err != nil {
		fmt.Printf("Cannot write pdf: Reason %s\n", err.Error())
		return "", err
	}

	return outputPath, nil
}
