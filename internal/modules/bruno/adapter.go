package bruno

import (
	"net/http"

	"github.com/marcobouwmeester/proxytoolkit/internal/adapters"
)

type brunoProps struct{}

func (b brunoProps) OnRequest(req *http.Request) {
	// to be implemented
}

func (b brunoProps) OnResponse(req *http.Request, res *http.Response) {
	// to be implemented
}

/**
 * Should be implemented
 */
func New() adapters.InterceptionAdapter {
	bruno := &brunoProps{}

	return adapters.InterceptionAdapter{
		OnRequest:  bruno.OnRequest,
		OnResponse: bruno.OnResponse,
	}
}
