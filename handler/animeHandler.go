package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

func AnimeHandler(r *gin.Engine) {
	r.POST("/anime/add", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for _, file := range files {
			log.Println(file.Filename)
		}

		c.Writer.WriteHeader(200)
	})
}
