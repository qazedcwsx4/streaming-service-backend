package handler

import (
	"anime-server-go/database"
	"github.com/gin-gonic/gin"
)

func AnimeHandler(r *gin.Engine) {
	entries := database.GetAllEntries()

	r.GET("/anime/getAll", func(c *gin.Context) {
		c.JSON(200, entries)
	})

	r.GET("/")
}