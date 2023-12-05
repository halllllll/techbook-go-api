package middleware

import (
	"log"
	"net/http"
)

type resLoggingWriter struct {
	http.ResponseWriter // インターフェースの型をそのまま使うことで移譲できる（インターフェースで宣言されているメソッドを明示的に実装しなくてもそのままつかえる
	// -> つまりresLogginWriterはhttp.ResponseWriterインターフェースを満たす
	// ついでにレスポンスコード保存用のフィールドも追加
	code int
}

// 移譲しているのでメソッドを明示的に書かなくてもいいが、書くとオーバーライド的なことができる
func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code

	// 本来のwriterheaderの機能を呼び出す
	rsw.ResponseWriter.WriteHeader(code)
}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.RequestURI, r.Method)

		// middlewareの後処理にはhttp.ResponseWriter型の構造体を自作し、その中でresを使う
		rlw := NewResLoggingWriter(w)

		next.ServeHTTP(rlw, r)

		// 後処理
		log.Println("res: ", rlw.code)
	})
}
