package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"gihtub.com/halllllll/techbook-go-api/server/controllers"
	"gihtub.com/halllllll/techbook-go-api/server/routers"
	"gihtub.com/halllllll/techbook-go-api/server/services"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 本書で書かれていた「実行時に環境変数をオプションで渡す」が動かなかったのでgodotenvで読み込むことにした
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var (
		DB_USER     = os.Getenv("DB_USER")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_NAME     = os.Getenv("DB_NAME")
	)

	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}
	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppController(ser)
	r := routers.NewRouter(con)

	log.Println("server start at prot 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
