//File: server/main.go
package main

import (
	"errors"
	"strconv"

	"net/http"

	"github.com/yageek/glmf-go-api/models"
	"github.com/yageek/glmf-go-api/store"
	"gopkg.in/gin-gonic/gin.v1"
)

var (
	newsStore         store.Store = store.NewMemoryStore()
	ErrNotImplemented             = errors.New("Not implemented yet.")
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
func DeleteNews(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := newsStore.DeleteNews(int(id)); err != nil {
		c.AbortWithError(http.StatusNotFound, ErrNotImplemented)
		return
	}

	c.Status(http.StatusOK)

}

func UpdateNews(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	news := models.News{}
	if err := c.BindJSON(&news); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedNews, err := newsStore.UpdateNews(int(id), news)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, ErrNotImplemented)
		return
	}

	c.JSON(http.StatusOK, updatedNews)
}

func GetNews(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	news, err := newsStore.GetNews(int(id))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
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
		c.AbortWithError(http.StatusNotFound, ErrNotImplemented)
		return
	}
	c.JSON(http.StatusCreated, dbNews)
}

func main() {
	r := gin.Default()
	r.GET("/news", GetAllNews)
	r.GET("/news/:id", GetNews)
	r.PUT("/news", CreateNews)
	r.POST("/news/:id", UpdateNews)
	r.DELETE("/news/:id", DeleteNews)
	r.Run()
}
