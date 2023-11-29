package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

	// トランザクション
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	article_id := 1
	const sqlGetNice = `
		SELECT nice
		FROM articles
		WHERE article_id = ?
	`

	row := tx.QueryRow(sqlGetNice, article_id)
	if err := row.Err(); err != nil {
		// ロールバックを忘れずに
		fmt.Println(err)
		tx.Rollback()
		return
	}

	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		// ロールバックを忘れずに
		fmt.Println(err)
		tx.Rollback()
		return
	}

	const sqlUpdateNice = `
		UPDATE articles
		SET nice = ?
		WHERE article_id = ?
	`
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, article_id)
	if err != nil {
		// ロールバックを忘れずに
		fmt.Println(err)
		tx.Rollback()
		return
	}

	// コミットを忘れずに
	tx.Commit()

}
