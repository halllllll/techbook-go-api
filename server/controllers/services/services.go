package services

import "gihtub.com/halllllll/techbook-go-api/server/models"

type MyAppServicer interface {
	PostArticleService(models.Article) (models.Article, error)
	GetArticleService(int) (models.Article, error)
	GetArticleListService(int) ([]models.Article, error)
	PostNiceService(models.Article) (models.Article, error)

	PostCommentService(models.Comment) (models.Comment, error)
}
