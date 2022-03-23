package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello world! Time %s", time.Now().String())})
	})

	http.ListenAndServe("localhost:8080", router)
}
