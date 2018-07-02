package app

import (
	"net/http"
	"strconv"

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
	providerStr := "provider"
	productStr := "product"

	providerID := c.Param(providerStr)
	productID, _ := strconv.ParseInt(c.Param(productStr), 10, 64)

	renderData := gin.H{
		"title":      "Detail Product",
		"providerID": providerID,
		"productID":  productID,
	}

	c.HTML(http.StatusOK, "detail-product.tmpl", renderData)
}
