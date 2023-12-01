package controllers

import (
	"context"
	"fmt"
	"golang_image_cloudinary/pkg/images"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type CloudSvc interface {
	// upload have 3 params.
	// @Param File refer to file buffer
	// @Param pathDestination refer to target directory/bucket in cloud provider
	Upload(ctx context.Context, file interface{}, pathDestination string) (uri string, err error)
	// Remove(ctx context.Context, path string) (err error)
}

type Services struct {
	cloud CloudSvc
}

func UploadToCloudinary(c *gin.Context) {

	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load .env file: %v", err))
	}

	cloudName := os.Getenv("CLOUDINARY_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cloudProvider := images.NewCloudinary(cloudName, apiKey, apiSecret)

	if cloudProvider.IsError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": cloudProvider.IsError.Error(),
		})
		return
	}

	svc := Services{
		cloud: cloudProvider,
	}

	typeFile := c.PostForm("type")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	fileReader, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer fileReader.Close()

	url, err := svc.cloud.Upload(context.Background(), fileReader, typeFile)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": url,
		})
	}

}
