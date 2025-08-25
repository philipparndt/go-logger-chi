package chi

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

type chiLogFormatterAdapter struct {
	formatter LogFormatter
}

func (a *chiLogFormatterAdapter) NewLogEntry(r *http.Request) middleware.LogEntry {
	return &chiLogEntryAdapter{entry: a.formatter.NewLogEntry(r)}
}

type chiLogEntryAdapter struct {
	entry LogEntry
}

func (a *chiLogEntryAdapter) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	a.entry.Write(status, bytes, header, elapsed, extra)
}

func (a *chiLogEntryAdapter) Panic(v interface{}, stack []byte) {
	a.entry.Panic(v, stack)
}

func ChiLogger() func(next http.Handler) http.Handler {
	formatter := &CustomLogFormatter{}
	adapter := &chiLogFormatterAdapter{formatter: formatter}
	return middleware.RequestLogger(adapter)
}
