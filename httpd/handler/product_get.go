package handler

import (
	// "product/httpd/methods/"

	"product/httpd/methods/"

	"github.com/gin-gonic/gin"
)

func ProdutsGetHandler(product *methods.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := product.GetAll()
		c.JSON(200, results)
	}
}
