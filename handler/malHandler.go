package handler

import (
	"anime-server-go/model"
	"anime-server-go/service"
	"github.com/gin-gonic/gin"
)

func MalHandler(r *gin.Engine) {
	r.GET("/mal/getById", getAnimeByMalId)
}

func getAnimeByMalId(c *gin.Context) {
	type getAnimeByMalIdParameters struct {
		MalId int `form:"malId" binding:"required"`
	}
	var parameters getAnimeByMalIdParameters

	if err := c.BindQuery(&parameters); err != nil {
		c.AbortWithError(400, err)
		return
	}

	anime, err := service.GetById(parameters.MalId)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, model.FromMalAnime(anime))
}
