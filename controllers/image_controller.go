package controllers

import (
	"fmt"
	"imagetopdf/contracts/responses"
	"imagetopdf/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteImage(ctx *gin.Context) {

	imageName := ctx.Param("image-name")

	deleted, err := services.DeleteImageFromStorage(imageName)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Success: false,
			Reason:  err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusBadRequest, responses.OkResponse{
		Success: deleted,
		Message: fmt.Sprintf("Image deleted %s successfully", imageName),
	})
}

func UploadImage(ctx *gin.Context) {
	form, err := ctx.MultipartForm()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Success: false,
			Reason:  err.Error(),
		})

		return
	}

	images := form.File["images"]

	uploaded, err := services.SaveImagesIntoStorage(images)

	if err != nil {

		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Success: false,
			Reason:  err.Error(),
		})

		return
	}

	var uploadImageResponse responses.UploadImageResponse

	for _, image := range images {
		uploadImageResponse.ImagesName = append(uploadImageResponse.ImagesName, image.Filename)
	}

	uploadImageResponse.Success = uploaded

	ctx.JSON(http.StatusOK, gin.H{
		"success":  uploadImageResponse.Success,
		"filename": uploadImageResponse.ImagesName,
	})
}
