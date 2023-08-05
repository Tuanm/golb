package middleware

import (
	"net/http"
	"net/http/httputil"

	"tuanm.dev/golb/config"
)

func RoundRobin(servers []string) http.Handler {
	rps := make([]*httputil.ReverseProxy, len(servers))
	for i, server := range servers {
		rps[i], _ = config.GetReverseProxy(server)
	}
	var count int

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rp := rps[count%len(rps)]
		count++
		rp.ServeHTTP(w, r)
	})
}
