package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/5112100070/Trek/src/app"
	"github.com/5112100070/Trek/src/app/page/dashboard"
	"github.com/5112100070/Trek/src/conf"
	"github.com/5112100070/Trek/src/global"
	"github.com/5112100070/publib/constants/env"
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
		network = "development"
	}

	fileLocation := fmt.Sprintf("/etc/trek/%s.ini", cfgenv)
	log.Println(fmt.Sprintf("Using configuration : %s", fileLocation))
	log.Println(fmt.Sprintf("Running in network : %s", network))

	// init regex
	errInitRegexImage := global.InitRegexCompileImageName()
	if errInitRegexImage != nil {
		log.Fatalln("[FATAL] Failed when try to compile regex: ", errInitRegexImage)
	}

	var ok bool
	// using
	conf.GConfig, ok = conf.ReadConfig(fileLocation)
	if !ok && cfgenv == env.Development {
		fileLocation = fmt.Sprintf("./config/%s.ini", cfgenv)
		conf.GConfig, ok = conf.ReadConfig(fileLocation)
		if !ok {
			log.Fatal("Could not find configuration file")
		}
	} else if !ok && cfgenv == env.Production {
		log.Fatal("Could not find configuration file")
	}
	conf.GConfig.Env = cfgenv

	redis := conf.InitRedis(conf.GConfig)
	db := conf.InitDatabase(conf.GConfig.Database)

	repoBundle := global.DBBundle{
		RedisSession: redis,
		DB:           db,
	}

	global.InitRepoBundle(repoBundle)
}

func main() {
	r := initEngine(conf.GConfig.Env)

	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))

	r.GET("/", app.IndexPageHandler)
	r.GET("/home", app.IndexPageHandler)
	r.GET("/index", app.IndexPageHandler)
	r.GET("/about-us", app.AboutUsPageHandler)
	r.GET("/faq", app.FaqPageHandler)
	r.GET("/sewa", app.RentPageHandler)
	r.GET("/penyedia", app.ProviderPageHandler)
	r.GET("/syarat", app.RequirementPageHandler)
	r.GET("/kebijakan", app.PolicyPageHandler)
	r.GET("/permintaan", app.RequestPageHandler)
	r.GET("/thank-you", app.ThankYouPageHandler)
	r.GET("/not-found", app.NotFoundPageHandler)

	r.GET("/alat", app.MarketPlacePageHandler)
	r.GET("/alat/:product", app.DetailProductHandler)

	r.POST("/contact/save", app.SaveContact)
	r.POST("/send-request-item", app.SendRequestItem)

	r.GET("/login", app.LoginPageHandler)
	r.GET("/daftar", app.RegisterPageHandler)
	r.GET("/reset-password", app.ResetPasswordPageHandler)
	r.GET("/register-confirmation/:register_id", app.RegisterConfirmationPage)

	r.GET("/dashboard", dashboard.MainPageHandler)
	r.GET("/dashboard/config", dashboard.ConfigPageHandler)
	r.GET("/dashboard/users", dashboard.UserListPageHandler)
	r.GET("/dashboard/user", dashboard.UserDetailPageHandler)
	r.GET("/dashboard/create-user", dashboard.UserCreatePagehandler)
	r.GET("/dashboard/update-user", dashboard.UserUpdatePagehandler)
	r.GET("/dashboard/create-company", dashboard.CompanyCreatePagehandler)
	r.GET("/dashboard/update-company", dashboard.CompanyUpdatePagehandler)
	r.GET("/dashboard/companies", dashboard.CompaniesListPageHandler)
	r.GET("/dashboard/company", dashboard.CompanyDetailPageHandler)
	r.GET("/dashboard/order", dashboard.OrdersDetailPageHandler)
	r.GET("/dashboard/orders", dashboard.OrdersListPageHandler)
	r.GET("/dashboard/order/create-order", dashboard.CreateOrderPageHandler)
	r.GET("/dashboard/modules", dashboard.ModulesPageHandler)
	r.GET("/dashboard/features", dashboard.FeaturePageHandler)
	r.GET("/dashboard/feature-member", dashboard.MemberFeaturePageHandler)

	// need to be removed
	r.GET("/dashboard/tambah-anggota", app.AddMemberPageHandler)
	r.GET("/dashboard/daftar-perusahaan", app.RegisterCompanyPageHandler)
	r.GET("/dashboard/company-profile", app.CompanyProfilePageHandler)
	r.GET("/dashboard/ganti-password", app.ChangePasswordPageHandler)
	// border need to be removed

	r.POST("/login", app.ProcessMakeLogin)
	r.POST("/logout", app.ProcessMakeLogout)
	r.POST("/change-password", dashboard.ChangePassword)

	r.POST("/dashboard/user/register", dashboard.CreateNewAccount)
	r.POST("/dashboard/user/update", dashboard.UpdateAccount)
	r.POST("/dashboard/company/register", dashboard.CreateNewCompany)
	r.POST("/dashboard/company/update", dashboard.UpdateCompany)
	r.POST("/admin/change-password", dashboard.AdminChangePassword)
	r.POST("/admin/change-account-activation", dashboard.AdminChangeActivation)
	r.POST("/admin/create-order", dashboard.CreateOrderForAdmin)

	r.GET("/admin", app.AdminDashboardPage)
	r.GET("/admin/product", app.AdminProductList)
	r.GET("/admin/product/new", app.AdminProductNew)
	r.GET("/admin/product/edit", app.AdminProductEditPage)
	r.GET("/admin/user", app.AdminUserList)
	r.GET("/admin/user/new", app.AdminUserNew)
	r.GET("/admin/user/edit", app.AdminUserEditPage)
	r.GET("/admin/ganti-password", app.ChangePasswordPageHandler)

	r.GET("/admin/index", app.AdminDashboardPage)

	r.Run(":4001")
}

func initEngine(cfenv string) *gin.Engine {
	if cfenv == env.Development {
		return initEngineDevelopment()
	} else {
		return initEngineProd()
	}
}

func initEngineDevelopment() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("files/WEB-INF/pages/**/*")

	r.Static("/css", "files/WEB-INF/attr/css")
	r.Static("/scss", "files/WEB-INF/attr/scss")
	r.Static("/vendor", "files/WEB-INF/attr/vendor")
	r.Static("/img", "files/WEB-INF/attr/img")
	r.Static("/file", "files/WEB-INF/attr/files")
	r.Static("/etc", "files/WEB-INF/attr/etc")
	r.Static("/js", "files/WEB-INF/attr/js")

	r.Static("/dashboard/css", "files/WEB-INF/attr-dashboard/css")
	r.Static("/dashboard/js", "files/WEB-INF/attr-dashboard/js")
	r.Static("/dashboard/scss", "/var/www/trek/dashboard/scss")
	r.Static("/dashboard/vendor", "/var/www/trek/dashboard/vendor")
	r.Static("/dashboard/assets", "files/WEB-INF/attr-dashboard/assets")

	return r
}

func initEngineProd() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("/var/www/trek/pages/**/*")

	r.Static("/css", "/var/www/trek/css")
	r.Static("/scss", "/var/www/trek/scss")
	r.Static("/vendor", "/var/www/trek/vendor")
	r.Static("/img", "/var/www/trek/img")
	r.Static("/file", "/var/www/trek/files")
	r.Static("/etc", "/var/www/trek/etc")
	r.Static("/js", "/var/www/trek/js")

	r.Static("/dashboard/css", "/var/www/trek/dashboard/css")
	r.Static("/dashboard/js", "/var/www/trek/dashboard/js")
	r.Static("/dashboard/scss", "/var/www/trek/dashboard/scss")
	r.Static("/dashboard/vendor", "/var/www/trek/dashboard/vendor")
	r.Static("/dashboard/assets", "files/WEB-INF/attr-dashboard/assets")

	return r
}
