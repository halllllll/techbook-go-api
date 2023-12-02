package services

import (
	"fmt"

	"github.com/halllllll/techbook-go-api/server/models"
	"github.com/halllllll/techbook-go-api/server/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {

	// 1. 記事を取得
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		fmt.Println(err)
		return models.Article{}, err
	}

	// 2. コメント一覧を取得
	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// 3. 1に2を連結
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil

}

// func PostAritcleService(article models.Article) (models.Article, error) {
func (s *MyAppService) PostAritcleService(article models.Article) (models.Article, error) {
	// // sql.DBを手に入れて、それ経由でrepositoryを操作
	// db, err := connectDB()
	// if err != nil {
	// 	return models.Article{}, err
	// }
	// // 呼び出し側でclose
	// defer db.Close()

	// newArticle, err := repositories.InsertArticle(db, article)
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

func (s *MyAppService) GetAritcleListService(page int) ([]models.Article, error) {

	articleList, err := repositories.SelectArticleList(s.db, page)

	if err != nil {
		return []models.Article{}, err
	}

	return articleList, nil
}

// 戻り値はarticle
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {

	err := repositories.UpdateNiceNum(s.db, article.ID)
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
