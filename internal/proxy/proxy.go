package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/marcobouwmeester/proxytoolkit/internal/adapters"
)

type interceptingTransport struct {
	base     http.RoundTripper
	adapters *[]adapters.InterceptionAdapter
}

func (t *interceptingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for _, ad := range *t.adapters {
		if ad.OnRequest != nil {
			(ad.OnRequest)(req)
		}
	}

	resp, err := t.base.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	for _, ad := range *t.adapters {
		if ad.OnResponse != nil {
			(ad.OnResponse)(req, resp)
		}
	}

	return resp, nil
}

// New returns a reverse proxy handler with interception
func New(target string, adapters *[]adapters.InterceptionAdapter) (http.Handler, error) {
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	rp := httputil.NewSingleHostReverseProxy(targetURL)

	// install our intercepting transport
	rp.Transport = &interceptingTransport{
		base:     http.DefaultTransport,
		adapters: adapters,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Host = targetURL.Hostname()
		rp.ServeHTTP(w, r)
	}), nil
}
