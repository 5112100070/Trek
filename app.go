package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/5112100070/Trek/src/app"
	"github.com/5112100070/Trek/src/conf"
	"github.com/5112100070/Trek/src/global"
)

func init() {
	fmt.Println("Lihat Disini say")
	// init error logging
	global.InitLogError(os.Stderr)

	cfgenv := os.Getenv("TREKENV")
	network := os.Getenv("NETWORK")
	if cfgenv == "" {
		log.Println("[trek] No environment set. Using 'development'.")
		log.Println("[trek] Use 'export TKPENV=[development|alpha|staging|production]' to change.")
		cfgenv = "development"
	}

	fileLocation := fmt.Sprintf("/etc/trek/%s.ini", cfgenv)
	log.Println(fmt.Sprintf("Using configuration : %s", fileLocation))
	log.Println(fmt.Sprintf("Running in network : %s", network))

	var ok bool
	conf.GConfig, ok = conf.ReadConfig(fileLocation)
	if !ok {
		log.Fatal("Could not find configuration file")
	}
}

func main() {
	r := initEngine()

	r.GET("/", app.IndexPageHandler)
	r.GET("/index", app.IndexPageHandler)
	r.GET("/sewa", app.RentPageHandler)
	r.GET("/penyedia", app.ProviderPageHandler)
	r.GET("/syarat", app.RequirementPageHandler)
	r.GET("/kebijakan", app.PolicyPageHandler)
	r.GET("/permintaan", app.RequestPageHandler)

	r.GET("/provider/:provider/:product", app.DetailProductHandler)

	r.Run(":4000") // listen and serve on 0.0.0.0:8080
}

func initEngine() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("/var/www/trek/pages/**/*")

	r.Static("/css", "/var/www/trek/css")
	r.Static("/scss", "/var/www/trek/scss")
	r.Static("/vendor", "/var/www/trek/vendor")
	r.Static("/img", "/var/www/trek/img")
	r.Static("/file", "/var/www/trek/files")
	r.Static("/etc", "/var/www/trek/etc")
	r.Static("/js", "/var/www/trek/js")

	return r
}
