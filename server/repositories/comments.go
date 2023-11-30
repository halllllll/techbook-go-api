package repositories

import (
	"database/sql"

	"github.com/halllllll/techbook-go-api/server/models"
)

func InsertComent(db *sql.DB, comment models.Comment) (models.Comment, error) {

	var newComment models.Comment
	const sqlStr = `
		INSERT INTO comments (article_id, message, created_at) VALUES
		(?, ?, now());
	`

	newComment.ArticleID, newComment.Message = comment.CommentID, comment.Message

	result, err := db.Exec(sqlStr, newComment.ArticleID, newComment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	// インクリメントされたIDを受け取って更新して返却
	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)

	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		SELECT * FROM comments
		WHERE article_id = ?;
	`

	comments := make([]models.Comment, 0)

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		// エラーが起きても無視？
		// if err := rows.Scan(&comment); err != nil {
		// 	return []models.Comment{}, err
		// }
		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
