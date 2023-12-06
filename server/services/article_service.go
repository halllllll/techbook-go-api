package services

import (
	"database/sql"
	"errors"

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
	var article models.Article
	var commentList []models.Comment
	var articleGetError, commentGetError error

	// var aMux sync.Mutex
	// var cMux sync.Mutex

	// var wg sync.WaitGroup
	// wg.Add(2)

	// WaitGroupを使った並列処理をチャネルを使ったものに置き換える
	type articleResult struct {
		article models.Article
		err     error
	}
	articleChan := make(chan articleResult)
	defer close(articleChan)

	// go func(db *sql.DB, articleID int) {
	// 	// defer wg.Done()
	// 	// aMux.Lock()
	// 	article, articleGetError = repositories.SelectArticleDetail(db, articleID)
	// 	// aMux.Unlock()
	// }(s.db, articleID)

	go func(ch chan<- articleResult, db *sql.DB, articleID int) {
		article, err := repositories.SelectArticleDetail(db, articleID)
		ch <- articleResult{article: article, err: err}
	}(articleChan, s.db, articleID)

	// WaitGroupを使った並列処理をチャネルを使ったものに置き換える
	type commentResult struct {
		commentList *[]models.Comment
		err         error
	}
	commentChan := make(chan commentResult)
	defer close(commentChan)

	// go func(db *sql.DB, articleID int) {
	// 	// defer wg.Done()
	// 	// cMux.Lock()
	// 	commentList, commentGetError = repositories.SelectCommentList(db, articleID)
	// 	// cMux.Unlock()
	// }(s.db, articleID)

	go func(ch chan<- commentResult, db *sql.DB, articleID int) {
		commentList, err := repositories.SelectCommentList(db, articleID)
		ch <- commentResult{
			commentList: &commentList, err: err,
		}
	}(commentChan, s.db, articleID)

	// wg.Wait()
	// WaitGroupではなくチャネルを使った並列処理に変える
	// select文は上から順に評価されるわけではなく、待ち受けているcaseに値が入ったら実行される
	// forで2回回しているので、チャネルからの受信を計2回受け付けている
	for i := 0; i < 2; i++ {
		select {
		case ar := <-articleChan:
			article, articleGetError = ar.article, ar.err
		case cr := <-commentChan:
			commentList, commentGetError = *cr.commentList, cr.err
		}
	}

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
