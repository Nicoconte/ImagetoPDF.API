package controllers

import (
	"fmt"
	"imagetopdf/contracts/responses"
	"imagetopdf/data"
	"imagetopdf/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func StartSession(ctx *gin.Context) {
	session, err := services.CreateSession()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Success: false,
			Reason:  fmt.Sprintf("Cannot start a new session. Reason %s", err.Error()),
		})

		return
	}

	//Create session folder
	folderFullpath := data.Config.StoragePath + session

	err = os.Mkdir(folderFullpath, 0755)
	err = os.Mkdir(folderFullpath+"/output", 0755)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Success: false,
			Reason:  fmt.Sprintf("Cannot start a new session. Reason %s", err.Error()),
		})

		return
	}

	ctx.JSON(http.StatusOK, responses.StartSessionResponse{
		SessionID: session,
	})
}

func EndSession(ctx *gin.Context) {
	var id = ctx.Param("id")

	if !services.SessionExists(id) {
		ctx.JSON(http.StatusNotFound, responses.ErrorResponse{
			Success: false,
			Reason:  fmt.Sprintf("Session %s does not exist", id),
		})

		return
	}

	err := services.DeleteSession(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Success: false,
			Reason:  "Cannot end session request by user",
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.OkResponse{
		Success: true,
		Message: "Session ended",
	})
}

func GetSessionStatus(ctx *gin.Context) {
	var id = ctx.Param("id")

	if !services.SessionExists(id) {
		ctx.JSON(http.StatusNotFound, responses.ErrorResponse{
			Success: false,
			Reason:  fmt.Sprintf("Session %s does not exist", id),
		})

		return
	}

	isActive := services.CheckIfSessionIsActive(id)

	err := services.UpdateSessionTime(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Success: false,
			Reason:  fmt.Sprintf("Unexpected error. Reason: %s", err.Error()),
		})

		return
	}

	var sessionStatus string

	if isActive {
		sessionStatus = "OK"
	} else {
		sessionStatus = "DOWN"
	}

	ctx.JSON(http.StatusOK, responses.SessionStatusResponse{
		Success: true,
		Status:  isActive,
		Message: fmt.Sprintf("Session status: %s", sessionStatus),
	})
}
