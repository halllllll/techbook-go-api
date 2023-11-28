package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/halllllll/techbook-go-api/db/models"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	// articleID := 1

	// const sqlStr = `
	// 	SELECT * FROM articles WHERE article_id = ?;
	// `
	// row := db.QueryRow(sqlStr, articleID)
	// if err := row.Err(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// var article models.Article
	// var createdTime sql.NullTime

	// err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if createdTime.Valid {
	// 	article.CreatedAt = createdTime.Time
	// }

	article := models.Article{
		Title:    "insert test ONE",
		Contents: "Can I Insert Data Correctly?",
		UserName: "saki-ikas",
	}
	const sqlStr = `
		INSERT INTO articles (
			title, contents, username, nice, created_at
		) VALUES (
			?, ?, ?, 0, now()
		);
	`

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}