package main

import (
	"product/httpd/handler"
	"product/httpd/methods"

	"github.com/gin-gonic/gin"
)

func main() {
	_product := methods.New()

	r := gin.Default()
	// r.GET("/ping", handler.Pingget)
	r.GET("/produts", handler.ProdutsGetHandler(_product))
	r.Post("produts", handler.ProductPostHandler(_product))

	r.Run()
}
