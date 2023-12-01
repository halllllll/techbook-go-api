package services

import (
	"github.com/halllllll/techbook-go-api/server/models"
	"github.com/halllllll/techbook-go-api/server/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {

	newComment, err := repositories.InsertComent(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
