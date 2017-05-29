package store

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yageek/glmf-go-api/models"
)

type MariaDBStore struct {
	db *sql.DB
}

func NewMariaDBStore(address string) (*MariaDBStore, error) {

	if db, err := sql.Open("mysql", address); err != nil {
		return nil, err
	} else {
		return &MariaDBStore{db}, nil
	}
}

func (s *MariaDBStore) CreateNews(title, content string) (models.News, error) {
	now := time.Now()
	result, err := s.db.Exec("INSERT INTO news (title, content, created_at, modified_at) VALUES (?, ?, ?, ?)", title, content, now, now)
	if err != nil {
		return models.News{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.News{}, err
	}
	return models.News{int(id), title, content, now, now}, nil
}

func (s *MariaDBStore) UpdateNews(ID int, news models.News) (models.News, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return models.News{}, err
	}
	_, err = s.db.Exec("UPDATE news SET title = ?, content = ?, modified_at = ? WHERE id=?", news.Title, news.Content, time.Now(), ID)
	if err != nil {
		return models.News{}, err
	}
	updatedNews, err := s.GetNews(ID)
	if err != nil {
		return models.News{}, err
	}
	return updatedNews, tx.Commit()
}

func (s *MariaDBStore) DeleteNews(ID int) error {
	_, err := s.db.Exec("DELETE FROM news WHERE id=?", ID)
	return err
}
func (s *MariaDBStore) GetNews(ID int) (models.News, error) {
	type row struct {
		ID       int64
		News     string
		Content  string
		Title    string
		Created  time.Time
		Modified time.Time
	}
	var r = row{}
	err := s.db.QueryRow("SELECT id,title,content,created_at,modified_at FROM news WHERE id=?", ID).Scan(&r.ID, &r.Title, &r.Content, &r.Created, &r.Modified)
	if err != nil {
		return models.News{}, err
	}
	return models.News{int(r.ID), r.Title, r.Content, r.Created, r.Modified}, nil
}

func (s *MariaDBStore) GetAllNews() ([]models.News, error) {
	rows, err := s.db.Query("SELECT id,title,content,created_at,modified_at FROM news")
	if err != nil {
		return []models.News{}, err
	}
	defer rows.Close()

	type row struct {
		ID       int64
		News     string
		Content  string
		Title    string
		Created  time.Time
		Modified time.Time
	}
	got := []models.News{}

	for rows.Next() {
		var r = row{}
		err := rows.Scan(&r.ID, &r.Title, &r.Content, &r.Created, &r.Modified)
		if err != nil {
			return []models.News{}, err
		}
		news := models.News{int(r.ID), r.Title, r.Content, r.Created, r.Modified}
		if err != nil {
			return []models.News{}, err
		}
		got = append(got, news)
	}
	return got, nil
}
