package services

import (
	"gihtub.com/halllllll/techbook-go-api/server/models"
)

type ArticleServicer interface {
	PostArticleService(models.Article) (models.Article, error)
	GetArticleService(int) (models.Article, error)
	GetArticleListService(int) ([]models.Article, error)
	PostNiceService(models.Article) (models.Article, error)
}

type CommentServicer interface {
	PostCommentService(models.Comment) (models.Comment, error)
}
