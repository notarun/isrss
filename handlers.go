package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllNews(c *gin.Context) {
	result, err := GetResults("all_news")
	if err != nil {
		c.Error(err)
		c.String(http.StatusBadGateway, err.Error())
		return
	}
	c.XML(http.StatusOK, NewRSS(result, "Inshorts: All"))
}

func Trending(c *gin.Context) {
	result, err := GetResults(c.Query("trending"))
	if err != nil {
		c.Error(err)
		c.String(http.StatusBadGateway, err.Error())
		return
	}
	c.XML(http.StatusOK, NewRSS(result, "Inshorts: Trending"))
}

func TopStories(c *gin.Context) {
	result, err := GetResults(c.Query("top_stories"))
	if err != nil {
		c.Error(err)
		c.String(http.StatusBadGateway, err.Error())
		return
	}
	c.XML(http.StatusOK, NewRSS(result, "Inshorts: Top"))
}
