package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"gihtub.com/halllllll/techbook-go-api/server/common"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	// 独自エラーに変換して扱う（独自エラーでない場合は作成する）
	var appErr *MyAppError
	if !errors.As(err, &appErr) {
		appErr = &MyAppError{
			ErrCode: Unkwon,
			Message: "Internal process faield",
			Err:     err,
		}
	}

	traceID := common.GetTraceID(req.Context())
	log.Printf("[%d] error: %s\n", traceID, appErr)

	var statusCode int
	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
