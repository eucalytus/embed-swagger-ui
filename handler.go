package handler

import (
	_ "github.com/eucalytus/embed-swagger-ui/statik"
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
	"net/http"
	"net/url"
	"path"
	"strconv"
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

func ServeWithCustomIndexHtml(prefix string, customIndexHtml string) http.Handler {
	needHandleIndexHtml := strings.TrimSpace(customIndexHtml) != ""
	buf := []byte(customIndexHtml)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if p := strings.TrimPrefix(r.URL.Path, prefix); len(p) < len(r.URL.Path) {
			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = p
			if needHandleIndexHtml && (p == "index.html" || p == "/") {
				w.WriteHeader(200)
				w.Header().Set("Content-Type", "text/html")
				w.Header().Set("Content-Length", strconv.FormatInt(int64(len(buf)), 10))
				w.Write([]byte(customIndexHtml))
			} else {
				SwaggerUIHandler.ServeHTTP(w, r2)
			}
		} else {
			http.NotFound(w, r)
		}
	})
}
