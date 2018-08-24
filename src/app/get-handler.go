package app

import (
	"fmt"
	"net/http"

	"github.com/5112100070/Trek/src/conf"
	"github.com/5112100070/Trek/src/global"
	"github.com/gin-gonic/gin"
)

func IndexPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	global.GetDefaultUserAttribute(c, renderData)

	c.HTML(http.StatusOK, "index.tmpl", renderData)
}

func RentPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	global.GetDefaultUserAttribute(c, renderData)

	c.HTML(http.StatusOK, "sewa.tmpl", renderData)
}

func ProviderPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	global.GetDefaultUserAttribute(c, renderData)

	c.HTML(http.StatusOK, "penyedia.tmpl", renderData)
}

func RequirementPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	global.GetDefaultUserAttribute(c, renderData)

	c.HTML(http.StatusOK, "syarat.tmpl", renderData)
}

func PolicyPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	global.GetDefaultUserAttribute(c, renderData)

	c.HTML(http.StatusOK, "kebijakan.tmpl", renderData)
}

func RequestPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	global.GetDefaultUserAttribute(c, renderData)

	c.HTML(http.StatusOK, "permintaan.tmpl", renderData)
}

func ThankYouPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	global.GetDefaultUserAttribute(c, renderData)

	c.HTML(http.StatusOK, "thank-you.tmpl", renderData)
}

func AboutUsPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	global.GetDefaultUserAttribute(c, renderData)

	c.HTML(http.StatusOK, "about-us.tmpl", renderData)
}

func LoginPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	global.GetDefaultUserAttribute(c, renderData)

	c.HTML(http.StatusOK, "login.tmpl", renderData)
}

func MarketPlacePageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	global.GetDefaultUserAttribute(c, renderData)

	c.HTML(http.StatusOK, "market-place.tmpl", renderData)
}

func DetailProductHandler(c *gin.Context) {
	productStr := "product"

	productName := c.Param(productStr)

	productPath := fmt.Sprintf("/alat/%v", productName)

	renderData := gin.H{
		"title":       "Detail Product",
		"productPath": productPath,
		"config":      conf.GConfig,
	}
	global.GetDefaultUserAttribute(c, renderData)

	c.HTML(http.StatusOK, "detail-product.tmpl", renderData)
}

func AdminDashboardPage(c *gin.Context) {
	cookie, errGetCookie := c.Cookie(global.UserCookie[global.GetEnv()])
	if errGetCookie != nil {
		global.Error.Println(errGetCookie)
		loginUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
		http.Redirect(c.Writer, c.Request, loginUrl, http.StatusSeeOther)
		return
	}

	service := global.GetServiceSession()

	user, errLogin := service.GetUser(cookie)
	if errLogin != nil {
		global.Error.Println(errLogin)
		loginUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
		http.Redirect(c.Writer, c.Request, loginUrl, http.StatusSeeOther)
		return
	}

	renderData := gin.H{
		"UserDetail": user,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "admin-index.tmpl", renderData)
}

func AdminProductList(c *gin.Context) {
	cookie, errGetCookie := c.Cookie(global.UserCookie[global.GetEnv()])
	if errGetCookie != nil {
		global.Error.Println(errGetCookie)
		loginUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
		http.Redirect(c.Writer, c.Request, loginUrl, http.StatusSeeOther)
		return
	}

	service := global.GetServiceSession()

	user, errLogin := service.GetUser(cookie)
	if errLogin != nil {
		global.Error.Println(errLogin)
		loginUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
		http.Redirect(c.Writer, c.Request, loginUrl, http.StatusSeeOther)
		return
	}

	renderData := gin.H{
		"UserDetail": user,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "admin-product-list.tmpl", renderData)
}

func AdminProductNew(c *gin.Context) {
	cookie, errGetCookie := c.Cookie(global.UserCookie[global.GetEnv()])
	if errGetCookie != nil {
		global.Error.Println(errGetCookie)
		loginUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
		http.Redirect(c.Writer, c.Request, loginUrl, http.StatusSeeOther)
		return
	}

	service := global.GetServiceSession()

	user, errLogin := service.GetUser(cookie)
	if errLogin != nil {
		global.Error.Println(errLogin)
		loginUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
		http.Redirect(c.Writer, c.Request, loginUrl, http.StatusSeeOther)
		return
	}

	renderData := gin.H{
		"UserDetail": user,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "admin-product-new.tmpl", renderData)
}

func AdminProductEditPage(c *gin.Context) {
	cookie, errGetCookie := c.Cookie(global.UserCookie[global.GetEnv()])
	if errGetCookie != nil {
		global.Error.Println(errGetCookie)
		loginUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
		http.Redirect(c.Writer, c.Request, loginUrl, http.StatusSeeOther)
		return
	}

	service := global.GetServiceSession()

	user, errLogin := service.GetUser(cookie)
	if errLogin != nil {
		global.Error.Println(errLogin)
		loginUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
		http.Redirect(c.Writer, c.Request, loginUrl, http.StatusSeeOther)
		return
	}

	renderData := gin.H{
		"UserDetail": user,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "admin-product-edit.tmpl", renderData)
}

func AdminUserList(c *gin.Context) {
	cookie, errGetCookie := c.Cookie(global.UserCookie[global.GetEnv()])
	if errGetCookie != nil {
		global.Error.Println(errGetCookie)
		loginUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
		http.Redirect(c.Writer, c.Request, loginUrl, http.StatusSeeOther)
		return
	}

	service := global.GetServiceSession()

	user, errLogin := service.GetUser(cookie)
	if errLogin != nil {
		global.Error.Println(errLogin)
		loginUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
		http.Redirect(c.Writer, c.Request, loginUrl, http.StatusSeeOther)
		return
	}

	renderData := gin.H{
		"UserDetail": user,
		"config":     conf.GConfig,
	}

	c.HTML(http.StatusOK, "admin-user-list.tmpl", renderData)
}

func AdminUserNew(c *gin.Context) {
	cookie, errGetCookie := c.Cookie(global.UserCookie[global.GetEnv()])
	if errGetCookie != nil {
		global.Error.Println(errGetCookie)
		loginUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
		http.Redirect(c.Writer, c.Request, loginUrl, http.StatusSeeOther)
		return
	}

	service := global.GetServiceSession()

	user, errLogin := service.GetUser(cookie)
	if errLogin != nil {
		global.Error.Println(errLogin)
		loginUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
		http.Redirect(c.Writer, c.Request, loginUrl, http.StatusSeeOther)
		return
	}

	renderData := gin.H{
		"UserDetail": user,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "admin-user-new.tmpl", renderData)
}

func AdminUserEditPage(c *gin.Context) {
	cookie, errGetCookie := c.Cookie(global.UserCookie[global.GetEnv()])
	if errGetCookie != nil {
		global.Error.Println(errGetCookie)
		loginUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
		http.Redirect(c.Writer, c.Request, loginUrl, http.StatusSeeOther)
		return
	}

	service := global.GetServiceSession()

	user, errLogin := service.GetUser(cookie)
	if errLogin != nil {
		global.Error.Println(errLogin)
		loginUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
		http.Redirect(c.Writer, c.Request, loginUrl, http.StatusSeeOther)
		return
	}

	renderData := gin.H{
		"UserDetail": user,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "admin-user-edit.tmpl", renderData)
}
