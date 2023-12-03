package repositories

import (
	"database/sql"

	"gihtub.com/halllllll/techbook-go-api/server/models"
)

var (
	articleNumPerPage = 5
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	INSERT INTO articles (title, contents, username, nice, created_at)
	VALUES (?, ?, ?, 0, now());
	`
	// わざわざ用意しているがそんなもんなのかもしれない
	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	tx, err := db.Begin()
	if err != nil {
		return models.Article{}, err
	}

	result, err := tx.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		tx.Rollback()
		return models.Article{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return models.Article{}, err
	}

	if err := tx.Commit(); err != nil {
		return models.Article{}, err
	}
	newArticle.ID = int(id)
	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
	SELECT article_id, title, contents, username, nice
	FROM articles
	LIMIT ? OFFSET ?;
	`

	tx, err := db.Begin()
	if err != nil {
		return []models.Article{}, err
	}

	rows, err := tx.Query(sqlStr, articleNumPerPage, ((page - 1) * articleNumPerPage))
	if err != nil {
		tx.Rollback()
		return []models.Article{}, err
	}
	defer rows.Close()

	articleList := make([]models.Article, 0)

	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum); err != nil {
			continue
		}

		articleList = append(articleList, article)
	}

	if err := tx.Commit(); err != nil {
		return []models.Article{}, err
	}
	return articleList, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {

	const sqlStr = `
	SELECT * FROM articles WHERE article_id = ?;
	`
	var article models.Article
	var createdTime sql.NullTime
	tx, err := db.Begin()

	if err != nil {
		tx.Rollback()
		return models.Article{}, err
	}

	row := tx.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return models.Article{}, err
	}

	if err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime); err != nil {
		tx.Rollback()
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	if err := tx.Commit(); err != nil {
		return models.Article{}, err
	}
	return article, nil

}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
		SELECT nice FROM articles WHERE article_id = ?;
	`

	var niceNum int
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}
	if err := row.Scan(&niceNum); err != nil {
		tx.Rollback()
		return err
	}

	const sqlUpdateNice = `
		UPDATE articles SET nice = ? WHERE article_id = ?;
	`

	if _, err := tx.Exec(sqlUpdateNice, niceNum+1, articleID); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
