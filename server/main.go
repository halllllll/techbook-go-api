package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/halllllll/techbook-go-api/server/handlers"
	"github.com/joho/godotenv"
)

func init() {
	// 本書で書かれていた「実行時に環境変数をオプションで渡す」が動かなかったのでgodotenvで読み込むことにした
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
