package main

import (
	"anime-server-go/handler"
	"anime-server-go/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	handler.AnimeHandler(r)
	handler.MalHandler(r)
	service.CreateAnime()
	r.Run() // listen and serve on 0.0.0.0:8080
}
