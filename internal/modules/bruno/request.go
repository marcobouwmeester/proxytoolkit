package bruno

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/marcobouwmeester/proxytoolkit/internal/config"
	"github.com/marcobouwmeester/proxytoolkit/internal/utils"
)

func mapToBruno(r *http.Request) (*BrunoRequest, error) {
	// --- Read body ---
	var rawBody string
	if r.Body != nil {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		rawBody = string(data)

		// restore body if needed later
		r.Body = io.NopCloser(bytes.NewBuffer(data))
	}

	// --- Query Params ---
	queryParams := map[string]string{}
	for k, v := range r.URL.Query() {
		if len(v) > 0 {
			queryParams[k] = v[0]
		}
	}

	// --- Headers ---
	headerVals := map[string]string{}
	for k, v := range r.Header {
		if len(v) > 0 {
			headerVals[k] = v[0]
		}
	}

	br := &BrunoRequest{
		Meta: MetaBlock{
			Name: utils.Slugify(r.URL.RequestURI()),
			Type: "http",
			Seq:  1,
		},
		Request: RequestBlock{
			Method: strings.ToLower(r.Method),
			URL:    r.URL.RequestURI(),
			Auth:   nil, // not available
		},
		Params: &ParamsBlock{
			Query: queryParams,
			Path:  map[string]string{}, // user call-in if needed
		},
		Headers: &HeadersBlock{
			Values: headerVals,
		},
		Body: &BodyBlock{
			Type: detectBodyType(r.Header, rawBody),
			Raw:  rawBody,
		},
		Settings: &SettingsBlock{
			EncodeURL: false,
			Timeout:   0,
		},
	}

	return br, nil
}

// Optional: detect JSON/form/text by headers
func detectBodyType(h http.Header, raw string) string {
	ct := h.Get("Content-Type")
	switch {
	case ct == "application/json":
		return "json"
	case ct == "application/x-www-form-urlencoded":
		return "form"
	case ct == "text/plain":
		return "text"
	case raw == "":
		return "none"
	default:
		return "raw"
	}
}

func HandleRequest(req *http.Request, cfg config.Config) error {
	data, err := mapToBruno(req)
	if err != nil {
		log.Error("Error mapping to bruno")
		return nil
	}

	fileName := fmt.Sprintf("%s.bru", data.Meta.Name)
	if err := CreateFileFromTemplate(
		cfg.ForwardAddr,
		"request.bru",
		data,
		&fileName,
	); err != nil {
		log.Error("Error creating collection.bru file")
	}
	return nil
}
