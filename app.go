package main

import (
	"github.com/gin-gonic/gin"

	"github.com/5112100070/Trek/src/app"
)

func main() {
	r := initEngine()

	r.GET("/", app.IndexPageHandler)
	r.GET("/index", app.IndexPageHandler)
	r.GET("/sewa", app.RentPageHandler)
	r.GET("/penyedia", app.ProviderPageHandler)
	r.GET("/syarat", app.RequirementPageHandler)
	r.GET("/kebijakan", app.PolicyPageHandler)
	r.GET("/permintaan", app.RequestPageHandler)

	r.Run(":4000") // listen and serve on 0.0.0.0:8080
}

func initEngine() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("/var/www/trek/pages/*")
	r.Static("/css", "/var/www/trek/css")
	r.Static("/scss", "/var/www/trek/scss")
	r.Static("/vendor", "/var/www/trek/vendor")
	r.Static("/img", "/var/www/trek/img")
	r.Static("/file", "/var/www/trek/files")
	r.Static("/etc", "/var/www/trek/etc")

	return r
}
