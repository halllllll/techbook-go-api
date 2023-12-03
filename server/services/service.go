package services

import (
	"database/sql"

	"gihtub.com/halllllll/techbook-go-api/server/models"
)

type MyAppService struct {
	db *sql.DB
}

// GetArticleListService implements services.MyAppServicer.
func (*MyAppService) GetArticleListService(int) ([]models.Article, error) {
	panic("unimplemented")
}

// PostArticleService implements services.MyAppServicer.
func (*MyAppService) PostArticleService(models.Article) (models.Article, error) {
	panic("unimplemented")
}

// PostNiceService implements services.MyAppServicer.
func (*MyAppService) PostNiceService(models.Article) (models.Article, error) {
	panic("unimplemented")
}

// コンストラクタ関数
func NewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}
