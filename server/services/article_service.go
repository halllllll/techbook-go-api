package services

import (
	"fmt"

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

func PostAritcleService(article models.Article) (models.Article, error) {
	// sql.DBを手に入れて、それ経由でrepositoryを操作
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	// 呼び出し側でclose
	defer db.Close()

	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

func GetAritcleListService(page int) ([]models.Article, error) {
	// sql.DBを手に入れて、それ経由でrepositoryを操作
	db, err := connectDB()
	if err != nil {
		return []models.Article{}, err
	}
	// 呼び出し側でclose
	defer db.Close()

	fmt.Println("よいしょ〜")
	articleList, err := repositories.SelectArticleList(db, page)

	if err != nil {
		return []models.Article{}, err
	}

	return articleList, nil
}

// 戻り値はarticle
func PostNiceService(article models.Article) (models.Article, error) {
	// sql.DBを手に入れて、それ経由でrepositoryを操作
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	// 呼び出し側でclose
	defer db.Close()

	err = repositories.UpdateNiceNum(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	// 戻り値はarticle
	// でもDBから取得するんじゃなくてここでプロパティ操作してる...
	// commentListはなぜか指定なし
	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
