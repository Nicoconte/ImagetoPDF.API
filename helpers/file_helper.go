package helpers

import (
	"io"
	"mime/multipart"
	"os"
)

func CreateFileFromRequestHeader(fileHeader *multipart.FileHeader, dst string) error {
	file, err := fileHeader.Open()

	if err != nil {
		return err
	}

	defer file.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)

	return err
}
