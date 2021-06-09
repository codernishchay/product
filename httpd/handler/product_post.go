package handler

import (
	"product/httpd/methods/"

	"github.com/gin-gonic/gin"
)

type productPost struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Stock       string `json:"stock"`
}

func ProductPostHandler(products *methods.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := productPost{}
		c.Bind(&requestBody)
		_products := methods.Product{
			Title:       requestBody.Title,
			Description: requestBody.Description,
			Stock:       requestBody.Stock,
		}
		products.Add(_products)
	}

}
