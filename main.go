package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", Index)
	r.GET("/all", AllNews)
	r.GET("/top", TopStories)
	r.GET("/trending", Trending)
	r.Run()
}
