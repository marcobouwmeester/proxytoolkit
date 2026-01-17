package adapters

import "net/http"

type (
	RequestAdapter      func(req *http.Request)
	ResponseAdapter     func(req *http.Request, res *http.Response)
	InterceptionAdapter struct {
		OnRequest  RequestAdapter
		OnResponse ResponseAdapter
	}
)
