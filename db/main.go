package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var (
		DB_USER     = os.Getenv("USERNAME")
		DB_PASSWORD = os.Getenv("USERPASS")
		DB_NAME     = os.Getenv("DATABASE")
	)

	dbUser := DB_USER
	dbPassword := DB_PASSWORD
	dbDatabase := DB_NAME
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
	result, err := tx.Exec(sqlUpdateNice, nicenum+1, article_id)
	if err != nil {
		// ロールバックを忘れずに
		fmt.Println(err)
		tx.Rollback()
		return
	}

	lid, _ := result.LastInsertId()
	fmt.Printf("last insertid: %d\n", lid)

	// コミットを忘れずに
	tx.Commit()

}
