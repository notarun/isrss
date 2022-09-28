package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsHandler(c *gin.Context) {
	result, err := GetResults()
	if err != nil {
		c.Error(err)
		c.String(http.StatusBadGateway, err.Error())
		return
	}
	c.XML(http.StatusOK, NewRSS(result))
}
