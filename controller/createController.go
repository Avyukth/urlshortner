package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	."urlshorner/migration"
	."urlshorner/model"
)
//var db = GetDb()
func CreateUrlShorter(c *gin.Context) {
	fmt.Println("here")
	c.HTML(http.StatusOK, "forms.html", gin.H{
		"title": "Users",
	})

}

func CretaeShortUrl(c *gin.Context) {
	urlData := c.PostForm("url")
	urlHashId := generateHash(urlData)
	shortenUrl := "http://rzp.com/" + urlHashId
	fmt.Println(shortenUrl)
	//fmt.Println(shortenUrl)

	if (!chackHashExist(urlHashId)){
		data := UrlModel{UrlHashId: urlHashId, Url: urlData, Shorten: shortenUrl}
		fmt.Println(data)
		Db.Debug().Create(&data)
		fmt.Println("Hello Data")

	}
	c.HTML(http.StatusOK, "success.html", gin.H{
		"shortU": shortenUrl,
	})
}

func chackHashExist(urlHashId string) bool {
	var todo UrlModel
	fmt.Println(todo.UrlHashId)
	Db.Where("url_hash_id = ?", urlHashId).First(&todo)
	fmt.Println("urlHashId2")

	if todo.ID == 0 {
		return false
	}
	fmt.Println("urlHashId3")
	return true

}

func generateHash(urlData string) string {
	h := md5.New()
	h.Write([]byte(urlData))
	urlHashId := hex.EncodeToString(h.Sum(nil))[:6]
	return urlHashId
}
