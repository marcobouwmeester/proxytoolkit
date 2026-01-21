package bruno

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/marcobouwmeester/proxytoolkit/internal/adapters"
	"github.com/marcobouwmeester/proxytoolkit/internal/config"
)

type brunoProps struct {
	Cfg config.Config
}

func (b brunoProps) OnRequest(req *http.Request) {
	if err := HandleRequest(req, b.Cfg); err != nil {
		log.Error("error handling request")
	}
}

/**
 * Should be implemented
 */
func New(cfg config.Config) adapters.InterceptionAdapter {
	bruno := &brunoProps{
		Cfg: cfg,
	}

	err := CreateBrunoConfigIfNotExists(cfg)
	if err != nil {
		log.Fatal("Error creating bruno config", err)
	}

	if err := CreateCollectionBruIfNotExists(cfg); err != nil {
		log.Fatal("Error creating collection.bru", err)
	}

	return adapters.InterceptionAdapter{
		OnRequest: bruno.OnRequest,
	}
}
