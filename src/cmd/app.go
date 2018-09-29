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

	db := conf.InitRedis(conf.GConfig)
	global.InitRepoBundle(db)
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
	r.GET("/thank-you", app.ThankYouPageHandler)
	r.GET("/tentang-kami", app.AboutUsPageHandler)
	r.GET("/not-found", app.NotFoundPageHandler)

	r.GET("/alat", app.MarketPlacePageHandler)
	r.GET("/alat/:product", app.DetailProductHandler)

	r.POST("/send-request-item", app.SendRequestItem)

	r.GET("/login", app.LoginPageHandler)
	r.GET("/daftar", app.RegisterPageHandler)
	r.GET("/reset-password", app.ResetPasswordPageHandler)
	r.GET("/register-confirmation/:register_id", app.RegisterConfirmationPage)
	r.POST("/login", app.ProcessMakeLogin)
	r.POST("/logout", app.ProcessMakeLogout)

	r.GET("/dashboard", app.DashboardPageHandler)
	r.GET("/dashboard/tambah-anggota", app.AddMemberPageHandler)
	r.GET("/dashboard/daftar-perusahaan", app.RegisterCompanyPageHandler)

	r.GET("/admin", app.AdminDashboardPage)
	r.GET("/admin/product", app.AdminProductList)
	r.GET("/admin/product/new", app.AdminProductNew)
	r.GET("/admin/product/edit", app.AdminProductEditPage)
	r.GET("/admin/user", app.AdminUserList)
	r.GET("/admin/user/new", app.AdminUserNew)
	r.GET("/admin/user/edit", app.AdminUserEditPage)

	r.GET("/admin/index", app.AdminDashboardPage)

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
