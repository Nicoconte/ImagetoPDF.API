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

	pdfName := ctx.Param("pdf-name")

	if pdfName == "default" {
		pdfName = helpers.GetGuid()
	}

	pdfName = fmt.Sprintf("%s.pdf", pdfName)

	targetPath, err := services.GeneratePDF(sessionId, pdfName)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Success: false,
			Reason:  err.Error(),
		})

		return
	}

	services.UpdateSessionTime(sessionId)

	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+pdfName)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(targetPath)

	services.DeletePDFFromStorage(targetPath)
	services.DeleteAllImagesAfterDownload(sessionId)
}
