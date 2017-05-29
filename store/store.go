package store

import (
	"errors"

	"github.com/yageek/glmf-go-api/models"
)

var (
	ErrNewsNotFound     = errors.New("The news was not found.")
	ErrCanNotDeleteNews = errors.New("Can not delete news")
)

type Store interface {
	GetAllNews() ([]models.News, error)
	CreateNews(title, content string) (models.News, error)
	UpdateNews(ID int, news models.News) (models.News, error)
	DeleteNews(ID int) error
	GetNews(ID int) (models.News, error)
}
