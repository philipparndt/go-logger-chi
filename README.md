# go-logger-chi

Shared logger configuration for my go chi projects.


```bash
go get github.com/philipparndt/go-logger-chi
```

Usage:
```go
package yourpackage

import (
  "github.com/philipparndt/go-logger-chi"
  "github.com/go-chi/chi/v5"
)


func main() {
  router := chi.NewRouter()
  router.Use(loggerchi.Middleware())
}
```
