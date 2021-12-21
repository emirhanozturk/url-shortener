package handlers

import (
	"encoding/json"
	"io/ioutil"
	"url-shortener/shortener"
	"url-shortener/store"

	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	body, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, &creationRequest)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	shortUrl := shortener.GenerateShortUrl(creationRequest.LongUrl)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl)
	host := "http://localhost:8080/"
	c.JSON(200, gin.H{
		"message":   "Short url created succesfully.",
		"short_url": host + shortUrl,
	})

}

func RedirectMainUrl(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
