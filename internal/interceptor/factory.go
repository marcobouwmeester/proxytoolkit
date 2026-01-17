package interceptor

import (
	"github.com/marcobouwmeester/proxytoolkit/internal/adapters"
	"github.com/marcobouwmeester/proxytoolkit/internal/config"
	"github.com/marcobouwmeester/proxytoolkit/internal/modules/bruno"
	"github.com/marcobouwmeester/proxytoolkit/internal/modules/profiler"
)

type InterceptorFactoryProps struct {
	Cfg *config.Config
}

func (p InterceptorFactoryProps) New() *[]adapters.InterceptionAdapter {
	adapters := []adapters.InterceptionAdapter{}

	if p.Cfg.LogApiTiming {
		adapters = append(adapters, profiler.New())
	}
	if p.Cfg.GenerateBrunoConfig {
		adapters = append(adapters, bruno.New())
	}
	return &adapters
}
