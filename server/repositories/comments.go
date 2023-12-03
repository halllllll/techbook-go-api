package repositories

import (
	"database/sql"

	"gihtub.com/halllllll/techbook-go-api/server/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
	INSERT INTO comments (article_id, message, created_at)
	VALUES(?, ?, now());
	`
	var newComment models.Comment
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, nil
	}
	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)

	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {

	const sqlStr = `
	SELECT * FROM comments WHERE article_id = ?;
	`
	var commentArray []models.Comment

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return []models.Comment{}, err
	}

	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		// エラーは無視
		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)
		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}
		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
