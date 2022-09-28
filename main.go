package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/all.xml", AllNews)
	r.GET("/top.xml", TopStories)
	r.GET("/trending.xml", Trending)
	r.Run()
}
