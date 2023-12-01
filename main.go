package main

import (
	"golang_image_cloudinary/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello, World!")
	})

	router.POST("/upload", controllers.UploadToCloudinary)
	router.GET("/image", controllers.GetUploadImage)

	router.Run()

}
