//File: server/main.go
package main

import (
	"strconv"

	"github.com/yageek/glmf-go-api/models"

	"net/http"

	"github.com/yageek/glmf-go-api/store"
	"gopkg.in/gin-gonic/gin.v1"
)

var (
	newsStore store.Store
	newsIDKey = "news_request_key"
)

// API News
func GetAllNews(c *gin.Context) {
	news, err := newsStore.GetAllNews()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, news)
}

type CreateRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreateNews(c *gin.Context) {
	news := CreateRequest{}
	if err := c.BindJSON(&news); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dbNews, err := newsStore.CreateNews(news.Title, news.Content)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusCreated, dbNews)
}
func DeleteNews(c *gin.Context) {
	id, _ := c.Get(newsIDKey)
	if err := newsStore.DeleteNews(id.(int)); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, nil)

}
func UpdateNews(c *gin.Context) {
	id, _ := c.Get(newsIDKey)
	news := models.News{}
	if err := c.BindJSON(&news); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	updatedNews, err := newsStore.UpdateNews(id.(int), news)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, updatedNews)
}

func GetNews(c *gin.Context) {

	id, _ := c.Get(newsIDKey)
	news, err := newsStore.GetNews(id.(int))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, news)
}
func newsIDMiddleware(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Set(newsIDKey, int(id))
	c.Next()
}

func main() {

	store, err := store.NewMariaDBStore("root:password@/glmf_go?parseTime=true")
	if err != nil {
		panic(err)
	}
	newsStore = store

	r := gin.Default()
	r.GET("/news", GetAllNews)
	r.PUT("/news", CreateNews)

	r.GET("/news/:id", newsIDMiddleware, GetNews)
	r.POST("/news/:id", newsIDMiddleware, UpdateNews)
	r.DELETE("/news/:id", newsIDMiddleware, DeleteNews)
	r.Run()
}
