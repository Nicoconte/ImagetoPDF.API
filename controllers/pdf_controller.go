package controllers

import (
	"fmt"
	"imagetopdf/contracts/responses"
	"imagetopdf/helpers"
	"imagetopdf/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConvertImagesToPDF(ctx *gin.Context) {
	sessionId := ctx.Request.Header.Get("session-key")

	if !services.SessionExists(sessionId) {
		ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Success: false,
			Reason:  "Unauthorized session",
		})

		return
	}

	pdfName := ctx.Param("pdf-name")

	if pdfName == "default" {
		pdfName = helpers.GetGuid()
	}

	targetPath, err := services.GeneratePDF(sessionId, fmt.Sprintf("%s.pdf", pdfName))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Success: false,
			Reason:  err.Error(),
		})
	}

	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+pdfName+".pdf")
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(targetPath)

	services.DeletePDFFromStorage(targetPath)
}
