package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllNews(c *gin.Context) {
	result, err := GetResults(ALL_NEWS)
	if err != nil {
		c.Error(err)
		c.String(http.StatusBadGateway, err.Error())
		return
	}
	c.XML(http.StatusOK, NewRSS(result, "Inshorts: All"))
}

func Trending(c *gin.Context) {
	result, err := GetResults(TRENDING_NEWS)
	if err != nil {
		c.Error(err)
		c.String(http.StatusBadGateway, err.Error())
		return
	}
	c.XML(http.StatusOK, NewRSS(result, "Inshorts: Trending"))
}

func TopStories(c *gin.Context) {
	result, err := GetResults(TOP_NEWS)
	if err != nil {
		c.Error(err)
		c.String(http.StatusBadGateway, err.Error())
		return
	}
	c.XML(http.StatusOK, NewRSS(result, "Inshorts: Top"))
}
