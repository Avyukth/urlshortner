package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BaseUrl(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Users",
	})
}
