package main

import (
	"product/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handler.Pingget)
	r.Run()
}
