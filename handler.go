package handler

import (
	_ "github.com/eucalytus/embed-swagger-ui/statik"
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
	"net/http"
	"path"
	"strings"
)

const INDEX = "index.html"

var StaticFs, _ = fs.New()

func exists(fs http.FileSystem, prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		name := path.Join(prefix, p)
		f, err := fs.Open(name)
		if err != nil {
			return false
		}
		stats, err := f.Stat()
		if err != nil {
			return false
		}
		if stats.IsDir() {
			index := path.Join(name, INDEX)
			f, err := fs.Open(index)
			if err != nil {
				return false
			}
			stats, err := f.Stat()
			if err != nil {
				return false
			}
			stats.Name()
		}
		return true
	}
	return false
}

// Static returns a middleware handler that serves static files in the given directory.
func Serve(urlPrefix string) gin.HandlerFunc {
	fileserver := http.FileServer(StaticFs)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin.Context) {
		if exists(StaticFs, urlPrefix, c.Request.URL.Path) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}

var SwaggerUIHandler = http.FileServer(StaticFs)
