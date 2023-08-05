package middleware

import (
	"fmt"
	"log"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("[method=%s|path=%s|remote_addr=%s]", r.Method, r.URL.Path, r.RemoteAddr)
		log.Default().Println(msg)
		next.ServeHTTP(w, r)
		log.Default().Println(msg, w.Header())
	})
}
