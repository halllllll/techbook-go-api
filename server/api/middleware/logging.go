package middleware

import (
	"context"
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
		// 直下じゃなくてここでtraceidを宣言する
		traceID := newTraceID()
		log.Printf("[%d]%s %s\n", traceID, r.RequestURI, r.Method)
		// ctx := SetTraceID(r.Context(), traceID)
		ctx := r.Context()
		ctx = context.WithValue(ctx, traceIDKey{}, traceID)
		req := r.WithContext(ctx)

		// middlewareの後処理にはhttp.ResponseWriter型の構造体を自作し、その中でresを使う
		rlw := NewResLoggingWriter(w)

		next.ServeHTTP(rlw, req)

		// 後処理
		log.Printf("[%d] res: %d\n", traceID, rlw.code)
	})
}
