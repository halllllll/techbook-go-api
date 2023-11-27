package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/halllllll/techbook-go-api/models"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")

}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, "cannot get content length\n", http.StatusBadRequest)
		return
	}

	reqBodybuffer := make([]byte, length)

	if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
		// もしも errors.Is 関数の結果が false だった場合エラーは io.EOF ではない 何か別の異常ということなので、その場合にはサーバー内で異常が起きたことを示す 500 番エラー を返すようにしています。
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	var reqArticle models.Article
	if err := json.Unmarshal(reqBodybuffer, &reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	// モックを返す
	article := reqArticle
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)

}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	// モックを返す
	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (page %d)\n", page)
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	w.Write(jsonData)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "invalid query parameter", http.StatusBadRequest)
		return
	}

	// モックを返す
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (articleID %d)\n", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// モックを返す
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "faile to encode nice json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// モックを返す
	comment := models.Comment2
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
