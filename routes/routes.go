package routes

import (
	"github.com/gin-gonic/gin"
	. "urlshorner/controller"
)

func AllRoutes() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	v1 := router.Group("/")
	{
		v1.GET("/", BaseUrl)
		v1.GET("/get", GetShortUrl)
		v1.POST("/long", RedirectShortUrl)
		v1.POST("/created", CreateShortUrl)
		v1.GET("/short", CreateUrlShorter)
		v1.GET("/upload", FileUploadForm)
		v1.POST("/fileUploadSuccess", FileUpload)
	}

	router.Run()

}
