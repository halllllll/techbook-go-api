package services

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// ここでもDBリポジトリと同じようにsql.DB定義
// GetEnvによる読み込み(go run main.go のオプションで流し込む予定)
var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

// closeは呼び出し側でやる
func connectDB() (*sql.DB, error) {
	fmt.Printf("dbconn: %s\n", dbConn)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
