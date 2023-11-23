package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const PATH = "public/upload"

func UploadImage(c *gin.Context) {

	typeFile := c.PostForm("type")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	location := fmt.Sprintf("%s/%s", PATH, typeFile)

	err = os.MkdirAll(location, os.ModePerm)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	//with gin
	c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", location, file.Filename))
	// err = c.Copy().SaveUploadedFile(file, fmt.Sprintf("%s/%s", location, file.Filename))
	// if err != nil {
	// 	c.JSON(500, gin.H{
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	//manual
	// dst, err := os.Create(fmt.Sprintf("%s/%s", location, file.Filename))
	// if err != nil {
	// 	c.JSON(500, gin.H{
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	// defer dst.Close()

	// f, err := file.Open()
	// if err != nil {
	// 	c.JSON(500, gin.H{
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	// defer f.Close()

	// if _, err := io.Copy(dst, f); err != nil {
	// 	c.JSON(500, gin.H{
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	c.JSON(200, gin.H{
		"message": "success upload image",
	})
}

func GetUploadImage(c *gin.Context) {
	typeFile := c.PostForm("type")
	fileName := c.PostForm("name")

	location := fmt.Sprintf("%s/%s/%s", PATH, typeFile, fileName)

	file, err := os.Open(location)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	defer file.Close()

	c.File(location)
}
