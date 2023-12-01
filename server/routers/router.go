package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/halllllll/techbook-go-api/server/controllers"
)

func NewRouter(con *controllers.MyAppController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/hello", con.HelloHandler)
	r.HandleFunc("/article", con.PostAritcleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	return r
}
