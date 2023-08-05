package config

import (
	"net/http/httputil"
	"net/url"
)

func GetReverseProxy(addr string) (rp *httputil.ReverseProxy, err error) {
	u, err := url.Parse(addr)
	if err == nil {
		rp = httputil.NewSingleHostReverseProxy(u)
	}
	return
}
