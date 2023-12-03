package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"gihtub.com/halllllll/techbook-go-api/server/apperrors"
	"gihtub.com/halllllll/techbook-go-api/server/controllers/services"
	"gihtub.com/halllllll/techbook-go-api/server/models"
	"github.com/gorilla/mux"
)

type ArticleController struct {
	service services.ArticleServicer
}

func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

func (c *ArticleController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, revision!!!\n")
}

func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		// どこで使う？ <- このあとすぐ回収される
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		// http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		apperrors.ErrorHandler(w, req, err)
		return
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		// http.Error(w, "fail to exec on PostArtice\n", http.StatusInternalServerError)
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			// どこで使う？
			err = apperrors.BadParam.Wrap(err, "queryparam must be number")
			// http.Error(w, "invalid query parameter", http.StatusBadRequest)
			// return
			apperrors.ErrorHandler(w, req, err)
		}
	} else {
		page = 1
	}

	artilceList, err := c.service.GetArticleListService(page)
	if err != nil {
		// http.Error(w, "fail to exec on GetArticleList", http.StatusInternalServerError)
		// return
		apperrors.ErrorHandler(w, req, err)
	}
	json.NewEncoder(w).Encode(artilceList)
}
func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "queryparam must include id")
		// http.Error(w, "invalid query parameter", http.StatusBadRequest)
		apperrors.ErrorHandler(w, req, err)
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		// http.Error(w, "fail internal exec on GetArticle\n", http.StatusInternalServerError)
		// return
		apperrors.ErrorHandler(w, req, err)
	}

	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		// どこで使う?
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		// http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		apperrors.ErrorHandler(w, req, err)
	}

	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		// http.Error(w, "fail to exec on PostNice", http.StatusInternalServerError)

		apperrors.ErrorHandler(w, req, err)

	}

	json.NewEncoder(w).Encode(article)

}
