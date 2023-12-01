package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/halllllll/techbook-go-api/server/controllers"
	"github.com/halllllll/techbook-go-api/server/routers"
	"github.com/halllllll/techbook-go-api/server/services"
	"github.com/joho/godotenv"
)

func init() {
	// 本書で書かれていた「実行時に環境変数をオプションで渡す」が動かなかったのでgodotenvで読み込むことにした
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

// var (
// 	dbUser     = os.Getenv("DB_USER")
// 	dbPassword = os.Getenv("DB_PASSWORD")
// 	dbDatabase = os.Getenv("DB_NAME")
// 	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
// )

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_NAME")
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}
	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppController(ser)

	r := routers.NewRouter(con)
	// r := mux.NewRouter()

	// r.HandleFunc("/article", con.PostAritcleHandler).Methods(http.MethodPost)
	// r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	// r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	// r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)
	// r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
