package services

import (
	"gihtub.com/halllllll/techbook-go-api/server/apperrors"
	"gihtub.com/halllllll/techbook-go-api/server/models"
	"gihtub.com/halllllll/techbook-go-api/server/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {

	// これこのままreturnしても中身同じだと思うけどあとで本書でリファクタリングされる可能性があるからとりあえず愚直にそのまま写経
	// (見事独自エラーのところで回収)
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
