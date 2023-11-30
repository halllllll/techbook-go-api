package services

import (
	"github.com/halllllll/techbook-go-api/server/models"
	"github.com/halllllll/techbook-go-api/server/repositories"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {
	// sql.DBを手に入れて、それ経由でrepositoryを操作
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	// 呼び出し側でclose
	defer db.Close()

	newComment, err := repositories.InsertComent(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
