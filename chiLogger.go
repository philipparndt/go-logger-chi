package loggerchi

import (
    "fmt"
    "net/http"
    "time"

    "github.com/philipparndt/go-logger"
)

type LogEntry interface {
    Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{})
    Panic(v interface{}, stack []byte)
}

type LogFormatter interface {
    NewLogEntry(r *http.Request) LogEntry
}

type CustomLogFormatter struct {
}

var okLevel string = "debug"

func SetOkLevel(level string) {
    okLevel = level
}

func (f *CustomLogFormatter) NewLogEntry(r *http.Request) LogEntry {
    return &CustomLogEntry{request: r}
}

type CustomLogEntry struct {
    request *http.Request
}

func (l *CustomLogEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
    scheme := "http"
    if l.request.TLS != nil {
        scheme = "https"
    }
    message := fmt.Sprintf("\"%s %s://%s%s %s\" %d %dB in %s from %s",
        l.request.Method,
        scheme,
        l.request.Host,
        l.request.RequestURI,
        l.request.Proto,
        status,
        bytes,
        elapsed,
        l.request.RemoteAddr,
    )
    switch {
    case status >= 500:
        logger.Error(message)
    case status >= 400:
        logger.Warn(message)
    default:
        logger.Log(okLevel, message)
    }
}

func (l *CustomLogEntry) Panic(v interface{}, stack []byte) {
    logger.Panic("Request panic", v, string(stack))
}
