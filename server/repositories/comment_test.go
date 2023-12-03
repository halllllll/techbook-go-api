package repositories_test

import (
	"testing"

	"gihtub.com/halllllll/techbook-go-api/server/models"
	"gihtub.com/halllllll/techbook-go-api/server/repositories"
)

func TestSelectCommentList(t *testing.T) {
	articleID := 1
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Errorf("want comment of articleID %d but got %d\n", articleID, comment.ArticleID)
		}
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 2,
		Message:   "test comment",
	}
	expectedCommentID := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Fatal(err)
	}

	if newComment.CommentID != expectedCommentID {
		t.Errorf("newa comment article id is %d but expected %d\n", newComment.CommentID, expectedCommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			DELETE FROM comments
			WHERE comment_id = ? AND message = ?;
		`

		testDB.Exec(sqlStr, expectedCommentID, newComment.Message)
	})
}
