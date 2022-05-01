package routes

import (
	"imagetopdf/controllers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router = HandleRoutes()

func HandleRoutes() *gin.Engine {

	r := gin.Default()

	r.Use(cors.Default())

	r.Static("/store", "./storage")

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello World"})
		return
	})

	//Image routes
	r.POST("api/v1/image", controllers.UploadImage)
	r.DELETE("api/v1/image/:image-name", controllers.DeleteImage)

	//Session routes
	r.POST("api/v1/session", controllers.StartSession)
	r.DELETE("api/v1/session/:id", controllers.EndSession)
	r.GET("api/v1/session/status/:id", controllers.GetSessionStatus)

	//PDF routes
	r.GET("api/v1/pdf/:pdf-name", controllers.ConvertImagesToPDF)

	return r
}
