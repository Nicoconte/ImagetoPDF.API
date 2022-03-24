package routes

import (
	"imagetopdf/controllers"

	"github.com/gin-gonic/gin"
)

var Router = HandleRoutes()

func HandleRoutes() *gin.Engine {

	r := gin.Default()

	r.POST("/image", controllers.UploadImage)
	r.DELETE("/image/:image-name", controllers.DeleteImage)

	return r
}
