package loggerchi

import (
    "github.com/philipparndt/go-logger"
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
    // do nothing
}

func TestRequestLoggerMiddleware(t *testing.T) {
    defer logger.LogTo(nil)

    SetOkLevel("info")

    var buf strings.Builder
    logger.LogTo(&buf)

    loggerWrapper := Middleware()

    w := httptest.NewRecorder()
    r, _ := http.NewRequest("GET", "/", nil)

    loggerWrapper(http.HandlerFunc(defaultHandler)).ServeHTTP(w, r)

    assert.Contains(t, buf.String(), "GET http:// HTTP/1.1")
}
