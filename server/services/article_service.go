package services

import (
	"database/sql"
	"errors"
	"sync"

	"gihtub.com/halllllll/techbook-go-api/server/apperrors"
	"gihtub.com/halllllll/techbook-go-api/server/models"
	"gihtub.com/halllllll/techbook-go-api/server/repositories"
)

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
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
	// goroutineで分けてそれぞれ並行処理をした結果の戻り値として
	var article models.Article
	var commentList []models.Comment
	var articleGetError, commentGetError error

	// race condition対策　ロック・アンロック
	var aMux sync.Mutex
	var cMux sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	// (go文は戻り値がある関数には使えないので無名即時関数)
	go func(db *sql.DB, articleID int) {
		defer wg.Done()
		aMux.Lock()
		article, articleGetError = repositories.SelectArticleDetail(db, articleID)
		aMux.Unlock()
	}(s.db, articleID)

	go func(db *sql.DB, articleID int) {
		defer wg.Done()
		cMux.Lock()
		commentList, commentGetError = repositories.SelectCommentList(db, articleID)
		cMux.Unlock()
	}(s.db, articleID)

	wg.Wait()

	if articleGetError != nil {
		if errors.Is(articleGetError, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetError, "no data")
			return models.Article{}, err
		}
		err := apperrors.GetDataFailed.Wrap(articleGetError, "fail to get data")
		return models.Article{}, err
	}

	if commentGetError != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetError, "fail to get data")
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
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

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {

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
