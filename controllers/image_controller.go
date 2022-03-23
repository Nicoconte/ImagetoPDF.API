package controllers

import (
	"imagetopdf/contracts/responses"
	"imagetopdf/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	ctx.JSON(http.StatusOK, gin.H{
		"success":  uploaded,
		"filename": images,
	})
}
