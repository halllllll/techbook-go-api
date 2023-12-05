package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		log.Println(req.RequestURI, req.Method)
		next.ServeHTTP(res, req)
	})
}
