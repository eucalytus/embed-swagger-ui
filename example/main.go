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
