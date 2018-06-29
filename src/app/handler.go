package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func RentPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "sewa.tmpl", nil)
}

func ProviderPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "penyedia.tmpl", nil)
}

func RequirementPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "syarat.tmpl", nil)
}

func PolicyPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "kebijakan.tmpl", nil)
}

func RequestPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "permintaan.tmpl", nil)
}

func DetailProductHandler(c *gin.Context) {
	provider := c.Param("provider")
	product := c.Param("product")

	fmt.Println(provider)
	fmt.Println(product)

	c.HTML(http.StatusOK, "detail-product.tmpl", nil)
}
