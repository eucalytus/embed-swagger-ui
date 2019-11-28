# embed-swagger-ui
embed swagger ui for golang

# Usage

## Gin 
```go
package main
import (
	handler "github.com/eucalytus/embed-swagger-ui"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	engine.Use(handler.Serve("/"))
	log.Fatal("gin run failed", engine.Run(":80"))
}
```

Golang Http

```go
package main

import (
	handler "github.com/eucalytus/embed-swagger-ui"
	"net/http"
)

func main() {
	http.Handle("/", handler.SwaggerUIHandler)
	http.ListenAndServe(":80", nil)
}
```