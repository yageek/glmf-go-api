package main

import (
	"errors"

	"net/http"

	"github.com/yageek/glmf-go-api/store"
	"gopkg.in/gin-gonic/gin.v1"
)

var (
	newsStore         store.Store
	ErrNotImplemented = errors.New("Not implemented yet.")
)

// API News
func GetAllNews(c *gin.Context) {
	c.AbortWithError(http.StatusNotFound, ErrNotImplemented)
}
func DeleteNews(c *gin.Context) {
	c.AbortWithError(http.StatusNotFound, ErrNotImplemented)
}
func UpdateNews(c *gin.Context) {
	c.AbortWithError(http.StatusNotFound, ErrNotImplemented)
}

func GetNews(c *gin.Context) {
	c.AbortWithError(http.StatusNotFound, ErrNotImplemented)
}

func CreateNews(c *gin.Context) {
	c.AbortWithError(http.StatusNotFound, ErrNotImplemented)
}

func main() {
	r := gin.Default()

	// Routeur
	r.GET("/news", GetAllNews)
	r.GET("/news/:id", GetNews)
	r.PUT("/news", CreateNews)
	r.POST("/news/:id", UpdateNews)
	r.DELETE("/news/:id", DeleteNews)

	// Lancement
	r.Run()
}
