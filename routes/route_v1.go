package routes

import (
	"imagetopdf/controllers"

	"github.com/gin-gonic/gin"
)

var Router = HandleRoutes()

func HandleRoutes() *gin.Engine {

	r := gin.Default()

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
