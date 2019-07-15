package main
//
//import (
//	"crypto/md5"
//	"encoding/hex"
//	"encoding/json"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"io/ioutil"
//	"net/http"
//	"os"
//	"strings"
//	//"database/sql"
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jinzhu/gorm"
//	."urlshorner/migration"
//	."urlshorner/model"
//)
//
//var db *gorm.DB
////var url = "www.google.com"
//
////var bleveIdx bleve.Index
////
////// Bleve connect or create the index persistence
////func Bleve(indexPath string) (bleve.Index, error) {
////
////	// with bleveIdx isn't set...
////	if bleveIdx == nil {
////		var err error
////		// try to open de persistence file...
////		bleveIdx, err = bleve.Open(indexPath)
////		// if doesn't exists or something goes wrong...
////		if err != nil {
////			// create a new mapping file and create a new index
////			mapping := bleve.NewIndexMapping()
////			bleveIdx, err = bleve.New(indexPath, mapping)
////			if err != nil {
////				return nil, err
////			}
////		}
////	}
////
////	// return de index
////	return bleveIdx, nil
////}
//
//func main() {
//
//	db=InitDB(db)
//	fmt.Println(db)
//	router := gin.Default()
//	router.LoadHTMLGlob("views/*")
//	v1 := router.Group("/")
//	{
//		v1.GET("/",baseUrl)
//		v1.GET("/{id}", baseUrl)
//		v1.GET("/get", getShortUrl)
//		v1.POST("/long", redirectShortUrl)
//		v1.POST("/created", cretaeShortUrl)
//		v1.GET("/short", createUrlShorter)
//		v1.GET("/upload", fileUploadForm)
//		v1.POST("/fileUploadSuccess", fileUpload)
//	}
//
//	router.Run()
//
//}
//
//type Urls struct {
//	Urls []string `json:"urls"`
//}
//func fileUploadForm(c *gin.Context){
//
//	c.HTML(http.StatusOK, "file.html", gin.H{
//		"title": "Users",
//	})
//}
//
//func fileUpload(c *gin.Context) {
//	fmt.Println("Hello gvHHbhewfbdhbcws hdenfjknn")
//
//	jsonPath := c.PostForm("jsonPath")
//	jsonFile, err := os.Open(jsonPath)
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer jsonFile.Close()
//	byteValue, _ := ioutil.ReadAll(jsonFile)
//	var urls Urls
//	json.Unmarshal(byteValue, &urls)
//	for i := 0; i < len(urls.Urls); i++{
//		urlHashId := generateHash(urls.Urls[i])
//		shortenUrl := "http://rzp.com/" + urlHashId
//		if (!chackHashExist(urlHashId)){
//			data := UrlModel{UrlHashId: urlHashId, Url: urls.Urls[i], Shorten: shortenUrl}
//			db.Debug().Create(&data)
//		}
//	}
//	c.HTML(http.StatusOK, "allSuccess.html", gin.H{
//		"title": "All Success",
//	})
//
//}
//
//func generateHash(urlData string) string {
//	h := md5.New()
//	h.Write([]byte(urlData))
//	urlHashId := hex.EncodeToString(h.Sum(nil))[:6]
//	return urlHashId
//}
//
//func chackHashExist(urlHashId string) bool {
//	var todo UrlModel
//	fmt.Println(todo.UrlHashId)
//	db.Where("url_hash_id = ?", urlHashId).First(&todo)
//	fmt.Println("urlHashId2")
//
//	if todo.ID == 0 {
//		return false
//	}
//	fmt.Println("urlHashId3")
//	return true
//
//}
//
//func cretaeShortUrl(c *gin.Context) {
//	urlData := c.PostForm("url")
//	urlHashId := generateHash(urlData)
//	shortenUrl := "http://rzp.com/" + urlHashId
//	fmt.Println(shortenUrl)
//	//fmt.Println(shortenUrl)
//
//	if (!chackHashExist(urlHashId)){
//		data := UrlModel{UrlHashId: urlHashId, Url: urlData, Shorten: shortenUrl}
//		fmt.Println(data)
//		db.Debug().Create(&data)
//		fmt.Println("Hello Data")
//
//	}
//	c.HTML(http.StatusOK, "success.html", gin.H{
//		"shortU": shortenUrl,
//	})
//}
//func getShortUrl(c *gin.Context) {
//
//	c.HTML(http.StatusOK, "shortForms.html", gin.H{
//		"title": "Users", "type": "Short",
//	})
//}
//func redirectShortUrl(c *gin.Context) {
//	urlData := c.PostForm("url")
//	hashID := strings.Split(urlData, "http://rzp:3030/")
//	var todo UrlModel
//	UrlHashId := hashID[1]
//
//	//db.First(&todo, UrlHashId)
//	db.Where("url_hash_id = ?", UrlHashId).First(&todo)
//	if todo.ID == 0 {
//		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No URL found!"})
//		return
//	}
//	//c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todo.Url})
//	c.HTML(http.StatusOK, "success.html", gin.H{
//		"shortU": todo.Url, "type": "Long",
//	})
//
//}
//func baseUrl(c *gin.Context) {
//	c.HTML(http.StatusOK, "index.html", gin.H{
//		"title": "Users",
//	})
//}
//func createUrlShorter(c *gin.Context) {
//	fmt.Println("here")
//	c.HTML(http.StatusOK, "forms.html", gin.H{
//		"title": "Users",
//	})
//
//}
