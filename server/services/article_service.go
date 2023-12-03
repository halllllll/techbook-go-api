package services

import (
	"database/sql"
	"errors"

	"gihtub.com/halllllll/techbook-go-api/server/apperrors"
	"gihtub.com/halllllll/techbook-go-api/server/models"
	"gihtub.com/halllllll/techbook-go-api/server/repositories"
)

func (s *MyAppService) PostAricleService(article models.Article) (models.Article, error) {
	// これそのままreturnしてもいいとは思うがあとで本書でリファクタリングするかもしれないのでとりあえず
	// (追記：独自エラーのところで見事回収)
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}

	return newArticle, nil
}

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return models.Article{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *MyAppService) GetAriticleListService(page int) ([]models.Article, error) {

	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return []models.Article{}, err
	}
	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return []models.Article{}, err
	}

	return articleList, nil
}

func (s *MyAppService) PostNiceSerivece(article models.Article) (models.Article, error) {

	// 奇妙奇天烈だが、あとで本書でリファクタリングされるかもしれない実装
	// (ここで+1をハードコーディングしている)
	if err := repositories.UpdateNiceNum(s.db, article.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to update nice count")
		return models.Article{}, err
	}

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
