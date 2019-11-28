package main

import (
	handler "github.com/eucalytus/embed-swagger-ui"
	_ "github.com/eucalytus/embed-swagger-ui/statik"
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
	"log"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	engine := gin.Default()
	engine.Use(handler.Serve("/", statikFS))

	log.Fatal("gin run failed", engine.Run(":80"))
}
