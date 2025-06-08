package main

// 安装：go install github.com/gin-gonic/gin@latest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func uploadSingleFile(r *gin.Engine) {
	//  curl -X POST  http://localhost:8080/upload1 -F "file=@go.mod"
	r.POST("/upload1", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		// c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, "%s uploaded!", file.Filename)
	})
}

func uploadMultipleFiles(r *gin.Engine) {
	//  curl -X POST http://localhost:8080/upload2 \
	// -F "upload[]=@go.sum" \
	// -F "upload[]=@go.mod"

	r.POST("/upload2", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, "%d files uploaded!", len(files))
	})
}
