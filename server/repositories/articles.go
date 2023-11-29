package repositories

import (
	"database/sql"

	"github.com/halllllll/techbook-go-api/server/models"
)

const (
	articleNumPerPage = 5
)

// 実行できるかのテストはあとの章でやっている
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	// sql文の中でniceとcreated_atの値を決めている
	const sqlStr = `
		INESRT INTO articles (title, contents, username, nice, created_at) 
		VALUES(?, ?, ?, 0, now());
	`
	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}
	// id取得
	id, _ := result.LastInsertId()
	newArticle.ID = int(id)
	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		SELECT article_id, title, contents, username, nice
		FROM articles
		LIMIT ? OFFSET ?;
	`

	articleArray := make([]models.Article, 0)

	rows, err := db.Query(sqlStr, articleNumPerPage, (page-1)*articleNumPerPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	var article models.Article
	var createdTime sql.NullTime
	const sqlStr = `
		SELECT * FROM articles WHERE article_id = ?;
	`
	row := db.QueryRow(sqlStr, articleID)
	// 確認必要
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	if err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime); err != nil {
		return models.Article{}, err
	}

	// sql.NullTime チェック（正常であればScanでロードしたcreatedTimeの値を使う）
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	const sqlGetNice = `
		SELECT nice
		FROM articles
		WHERE article_id = ?;
	`
	const sqlUpdateNice = `
		UPDATE articles SET nice = ? WHERE article_id = ?;
	`

	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var niceNum int
	if err := row.Scan(&niceNum); err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sqlUpdateNice, niceNum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		// tx.Rollback() ここでは必要ないらしい？
		return err
	}

	return nil
}
