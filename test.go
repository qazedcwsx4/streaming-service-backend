package main

import (
	"anime-server-go/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	handler.AnimeHandler(r)
	handler.MalHandler(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
