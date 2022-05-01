package controllers

import (
	"fmt"
	"imagetopdf/contracts/responses"
	"imagetopdf/data"
	"imagetopdf/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteImage(ctx *gin.Context) {
	//TODO. Maybe use it into a middleware
	sessionId := ctx.Request.Header.Get("session-key")

	if !services.SessionExists(sessionId) {
		ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Success: false,
			Reason:  "Unauthorized session",
		})

		return
	}

	imageName := ctx.Param("image-name")

	deleted, err := services.DeleteImageFromStorage(imageName, sessionId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Success: false,
			Reason:  err.Error(),
		})

		return
	}

	err = services.UpdateSessionTime(sessionId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Success: false,
			Reason:  err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, responses.OkResponse{
		Success: deleted,
		Message: fmt.Sprintf("Image '%s' was deleted successfully", imageName),
	})
}

func UploadImage(ctx *gin.Context) {
	sessionId := ctx.Request.Header.Get("session-key")

	if !services.SessionExists(sessionId) {
		ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Success: false,
			Reason:  "Unauthorized session",
		})

		return
	}

	form, err := ctx.MultipartForm()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Success: false,
			Reason:  err.Error(),
		})

		return
	}

	images := form.File["images"]

	uploaded, err := services.SaveImagesIntoStorage(images, sessionId)

	if err != nil {

		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Success: false,
			Reason:  err.Error(),
		})

		return
	}

	err = services.UpdateSessionTime(sessionId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Success: false,
			Reason:  err.Error(),
		})

		return
	}

	var uploadImageResponse responses.UploadImageResponse

	baseHost := fmt.Sprintf("%s:%s", data.Config.Host, data.Config.Port)

	storeUrl := fmt.Sprintf("%s/store/%s", baseHost, sessionId)

	for _, image := range images {
		uploadImageResponse.ImagesName = append(uploadImageResponse.ImagesName, fmt.Sprintf("%s/%s", storeUrl, image.Filename))
	}

	uploadImageResponse.Success = uploaded

	ctx.JSON(http.StatusOK, gin.H{
		"success":  uploadImageResponse.Success,
		"filename": uploadImageResponse.ImagesName,
	})
}
