package middlewares

import (
	"imagetopdf/contracts/responses"
	"imagetopdf/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckSessionExists() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestEndpoint := ctx.Request.RequestURI

		if strings.Contains(requestEndpoint, "session") || strings.Contains(requestEndpoint, "store") {
			return
		}

		sessionId := ctx.Request.Header.Get("session-key")

		if !services.SessionExists(sessionId) {
			ctx.JSON(http.StatusUnauthorized, responses.ErrorResponse{
				Success: false,
				Reason:  "Unauthorized session",
			})

			return
		}
	}
}
