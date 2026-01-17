package profiler

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/marcobouwmeester/proxytoolkit/internal/adapters"
)

type ctxKey string

const REQUEST_START_KEY ctxKey = "request-start-time"

type timeLoggerProps struct {
	logger *log.Logger
}

func (t *timeLoggerProps) onRequest(req *http.Request) {
	ctx := context.WithValue(
		req.Context(),
		REQUEST_START_KEY,
		time.Now(),
	)

	*req = *req.WithContext(ctx)
}

func (t *timeLoggerProps) onResponse(req *http.Request, res *http.Response) {
	startTime := req.Context().Value(REQUEST_START_KEY)
	if startTime == nil {
		t.logger.Warn(
			"missing start time",
			"path", req.RequestURI,
			"method", req.Method,
		)
		return
	}
	duration := time.Since(startTime.(time.Time))

	t.logger.Info(
		req.RequestURI,
		"duration", duration.String(),
		"method", req.Method,
		"status", res.StatusCode,
	)
}

func New() adapters.InterceptionAdapter {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "⏰",
	})

	timeLogger := &timeLoggerProps{
		logger: logger,
	}

	return adapters.InterceptionAdapter{
		OnRequest:  timeLogger.onRequest,
		OnResponse: timeLogger.onResponse,
	}
}
