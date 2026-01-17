package profiler

import (
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/marcobouwmeester/proxytoolkit/internal/adapters"
)

type timeLoggerProps struct {
	start  time.Time
	logger *log.Logger
}

func (t *timeLoggerProps) onRequest(req *http.Request) {
	t.start = time.Now()
}

func (t *timeLoggerProps) onResponse(req *http.Request, res *http.Response) {
	duration := time.Since(t.start)

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
