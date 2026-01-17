package main

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/marcobouwmeester/proxytoolkit/internal/config"
	"github.com/marcobouwmeester/proxytoolkit/internal/interceptor"
	"github.com/marcobouwmeester/proxytoolkit/internal/proxy"
)

func main() {
	cfg := config.Load()

	interceptorFactory := interceptor.InterceptorFactoryProps{Cfg: cfg}
	adapters := interceptorFactory.New()

	p, err := proxy.New(cfg.ForwardAddr, adapters)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Listening on :%s → %s", cfg.Port, cfg.ForwardAddr)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, p))
}
