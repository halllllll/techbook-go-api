package controllers

import (
	"encoding/json"
	"net/http"

	"gihtub.com/halllllll/techbook-go-api/server/apperrors"
	"gihtub.com/halllllll/techbook-go-api/server/controllers/services"
	"gihtub.com/halllllll/techbook-go-api/server/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		// どこで使う？
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail to exec on PostComment", http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
