package main

import (
	"fmt"
	"log"
	"net/http"

	"tuanm.dev/golb/config"
	m "tuanm.dev/golb/middleware"
)

func main() {
	conf, err := config.Load(".")
	if err == nil {
		log.Default().Printf("config loaded: %s\n", conf.String())
		var h http.Handler
		if len(conf.Servers) > 0 {
			h = m.RoundRobin(conf.Servers)
		}
		if conf.RateLimit > 0 {
			h = m.RateLimit(h, conf)
		}
		if conf.Logging {
			h = m.Logging(h)
		}
		err = http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), h)
	}
	log.Fatalln(err.Error())
}
