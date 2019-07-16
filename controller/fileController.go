package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	. "urlshorner/migration"
	. "urlshorner/model"
)

type Urls struct {
	Urls []string `json:"urls"`
}

func FileUpload(c *gin.Context) {

	jsonPath := c.PostForm("jsonPath")
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var urls Urls
	json.Unmarshal(byteValue, &urls)
	var wg sync.WaitGroup
	for i := 0; i < len(urls.Urls); i++ {
		wg.Add(1)
		go func(url string) {
			urlHashId := generateHash(url)
			shortenUrl := "http://rzp.com/" + urlHashId
			if !checkHashExist(urlHashId) {
				data := UrlModel{UrlHashId: urlHashId, Url: url, Shorten: shortenUrl}
				Db.Debug().Create(&data)
			}
			wg.Done()
		}(urls.Urls[i])
	}
	wg.Wait()

	c.HTML(http.StatusOK, "allSuccess.html", gin.H{
		"title": "All Success",
	})

}

func FileUploadForm(c *gin.Context) {

	c.HTML(http.StatusOK, "file.html", gin.H{
		"title": "Users",
	})
}
