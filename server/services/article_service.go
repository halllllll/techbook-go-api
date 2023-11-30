package services

import (
	"github.com/halllllll/techbook-go-api/server/models"
	"github.com/halllllll/techbook-go-api/server/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {
	// 記事と、その記事へのコメントが必要
	// sql.DBを手に入れて、それ経由でrepositoryを操作
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	// 呼び出し側でclose
	defer db.Close()

	// 1. 記事を取得
	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// 2. コメント一覧を取得
	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// 3. 1に2を連結
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil

}
