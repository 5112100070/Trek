package app

import (
	"net/http"
	"strconv"

	"github.com/5112100070/Trek/src/conf"
	"github.com/gin-gonic/gin"
)

func IndexPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	c.HTML(http.StatusOK, "index.tmpl", renderData)
}

func RentPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	c.HTML(http.StatusOK, "sewa.tmpl", renderData)
}

func ProviderPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	c.HTML(http.StatusOK, "penyedia.tmpl", renderData)
}

func RequirementPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	c.HTML(http.StatusOK, "syarat.tmpl", renderData)
}

func PolicyPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	c.HTML(http.StatusOK, "kebijakan.tmpl", renderData)
}

func RequestPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	c.HTML(http.StatusOK, "permintaan.tmpl", renderData)
}

func ThankYouPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	c.HTML(http.StatusOK, "thank-you.tmpl", renderData)
}

func AboutUsPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	c.HTML(http.StatusOK, "about-us.tmpl", renderData)
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
		"config":     conf.GConfig,
	}

	c.HTML(http.StatusOK, "detail-product.tmpl", renderData)
}
