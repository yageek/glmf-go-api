package store

import (
	"sync"
	"time"

	"github.com/yageek/glmf-go-api/models"
)

type MemoryStore struct {
	Contents map[int]models.News
	mu       sync.Mutex
	counter  int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{Contents: map[int]models.News{}}
}
func (m *MemoryStore) GetAllNews() ([]models.News, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	resultSet := make([]models.News, len(m.Contents))
	index := 0
	for _, value := range m.Contents {
		resultSet[index] = value
		index++
	}
	return resultSet, nil
}

func (m *MemoryStore) CreateNews(title, content string) (models.News, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	t := time.Now()
	news := models.News{
		ID:           m.counter,
		Title:        title,
		Content:      content,
		CreatedAt:    t,
		LastModified: t,
	}
	m.Contents[news.ID] = news
	m.counter++
	return news, nil
}
func (m *MemoryStore) UpdateNews(ID int, news models.News) (models.News, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	storeNews, exist := m.Contents[ID]
	if !exist {
		return models.News{}, ErrNewsNotFound
	}
	storeNews.Content = news.Content
	storeNews.Title = news.Title
	storeNews.LastModified = time.Now()
	m.Contents[ID] = storeNews
	return storeNews, nil
}
func (m *MemoryStore) DeleteNews(ID int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exist := m.Contents[ID]; !exist {
		return ErrNewsNotFound
	}
	delete(m.Contents, ID)
	return nil
}

func (m *MemoryStore) GetNews(ID int) (models.News, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	storeNews, exist := m.Contents[ID]
	if !exist {
		return models.News{}, ErrNewsNotFound
	}
	return storeNews, nil
}
