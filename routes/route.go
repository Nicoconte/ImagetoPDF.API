package routes

import (
	"imagetopdf/controllers"

	"github.com/gin-gonic/gin"
)

var Router = HandleRoutes()

func HandleRoutes() *gin.Engine {

	r := gin.Default()

	//Image routes
	r.POST("/image", controllers.UploadImage)
	r.DELETE("/image/:image-name", controllers.DeleteImage)

	//Session routes
	r.POST("/session", controllers.StartSession)
	r.DELETE("/session/:id", controllers.EndSession)
	r.GET("session/status/:id", controllers.GetSessionStatus)

	//PDF routes
	r.GET("/pdf/:pdf-name", controllers.ConvertImagesToPDF)

	return r
}
