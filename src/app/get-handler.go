package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/5112100070/Trek/src/app/session"
	"github.com/5112100070/Trek/src/conf"
	errorConst "github.com/5112100070/Trek/src/constants/error"
	"github.com/5112100070/Trek/src/global"
	gSession "github.com/5112100070/Trek/src/global/session"
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

func FaqPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	global.GetDefaultUserAttribute(c, renderData)

	c.HTML(http.StatusOK, "faq.tmpl", renderData)
}

func NotFoundPageHandler(c *gin.Context) {
	renderData := gin.H{
		"config": conf.GConfig,
	}
	global.GetDefaultUserAttribute(c, renderData)

	c.HTML(http.StatusOK, "not-found.tmpl", renderData)
}

func LoginPageHandler(c *gin.Context) {
	// all login user cannot see login page and redirected to main page
	// all not login user can see login page

	authToken, errGetCookie := c.Cookie(global.UserCookie[global.GetEnv()])
	if errGetCookie == nil {
		// all login user cannot see login page

		accountResp, errGetResponse := global.GetServiceSession().GetUser(authToken)
		if errGetResponse != nil {
			global.Error.Println(errGetResponse)

			// TODO: redirect to problem page
			// ====================== REMOVE / NEED CHANGES ======================
			renderData := gin.H{
				"config": conf.GConfig,
			}
			c.HTML(http.StatusOK, "login.tmpl", renderData)
			// ====================================================================
			return
		}

		// expire - we remove cookie and redirect it
		if accountResp.Error != nil {
			handleSessionErrorPage(c, *accountResp.Error, false)
			return
		}

		baseUrl := conf.GConfig.BaseUrlConfig.BaseDNS
		http.Redirect(c.Writer, c.Request, baseUrl, http.StatusSeeOther)
		return
	}

	// all login user can see login page
	renderData := gin.H{
		"config": conf.GConfig,
	}

	c.HTML(http.StatusOK, "login.tmpl", renderData)
}

func RegisterPageHandler(c *gin.Context) {
	_, errGetCookie := c.Cookie(global.UserCookie[global.GetEnv()])
	if errGetCookie == nil {
		loginUrl := conf.GConfig.BaseUrlConfig.BaseDNS
		http.Redirect(c.Writer, c.Request, loginUrl, http.StatusSeeOther)
		return
	}

	renderData := gin.H{
		"config": conf.GConfig,
	}
	c.HTML(http.StatusOK, "register.tmpl", renderData)
}

func ResetPasswordPageHandler(c *gin.Context) {
	_, errGetCookie := c.Cookie(global.UserCookie[global.GetEnv()])
	if errGetCookie == nil {
		loginUrl := conf.GConfig.BaseUrlConfig.BaseDNS
		http.Redirect(c.Writer, c.Request, loginUrl, http.StatusSeeOther)
		return
	}

	renderData := gin.H{
		"config": conf.GConfig,
	}
	c.HTML(http.StatusOK, "reset-password.tmpl", renderData)
}

func RegisterConfirmationPage(c *gin.Context) {
	registerID := c.Param("register_id")

	baseUrl := conf.GConfig.BaseUrlConfig.ProductDNS
	path := "/confirmation"
	data := url.Values{}
	data.Set("register_id", registerID)

	u, _ := url.ParseRequestURI(baseUrl)
	u.Path = path
	urlStr := u.String()

	client := &http.Client{}
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		global.Error.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, errGetResp := client.Do(req)
	if err != nil {
		global.Error.Println(errGetResp)
		global.InternalServerErrorResponse(c, errorConst.WORDING_ERROR_INTERNAL_SERVER)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.Error.Println(err)
		global.InternalServerErrorResponse(c, errorConst.WORDING_ERROR_INTERNAL_SERVER)
		return
	}

	var resultResp struct {
		ServerMessage string                 `json:"server_message"`
		Data          map[string]interface{} `json:"data,omitempty"`
	}

	var message string
	errUnMarshal := json.Unmarshal(body, &resultResp)
	if errUnMarshal != nil {
		global.Error.Println(errUnMarshal)
		http.Redirect(c.Writer, c.Request, conf.GConfig.BaseUrlConfig.BaseDNS+"/not-found", http.StatusSeeOther)
		return
	} else if !resultResp.Data["is_success"].(bool) {
		global.Error.Println(resultResp.Data["error_message"])
		if resultResp.Data["message"].(string) != "" {
			message = resultResp.Data["message"].(string)
		} else {
			message = "Terjadi kesalahan pada proses pendaftaran"
		}
	} else {
		message = "Terima kasih karena telah melakukan pendaftaran. Silahkan menuju Login untuk memulai Session"
	}

	renderData := gin.H{
		"message": message,
		"config":  conf.GConfig,
	}

	c.HTML(http.StatusOK, "info.tmpl", renderData)
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

func AddMemberPageHandler(c *gin.Context) {
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

	if !(user.Data.Role == session.USER_TYPE_ADMIN || user.Data.Role == session.USER_TYPE_COMMON) {
		notFoundUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/not-found"
		http.Redirect(c.Writer, c.Request, notFoundUrl, http.StatusSeeOther)
		return
	}

	renderData := gin.H{
		"UserDetail": user,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "dashboard-add-member.tmpl", renderData)
}

func RegisterCompanyPageHandler(c *gin.Context) {
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

	// if global.IsUserCanAccess(user) {
	// 	notFoundUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/not-found"
	// 	http.Redirect(c.Writer, c.Request, notFoundUrl, http.StatusSeeOther)
	// 	return
	// }

	renderData := gin.H{
		"UserDetail": user,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "dashboard-register-company.tmpl", renderData)
}

func CompanyProfilePageHandler(c *gin.Context) {
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

	// if global.IsUserCanAccess(user) {
	// 	notFoundUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/not-found"
	// 	http.Redirect(c.Writer, c.Request, notFoundUrl, http.StatusSeeOther)
	// 	return
	// }

	renderData := gin.H{
		"UserDetail": user,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "dashboard-company-profile.tmpl", renderData)
}

func ChangePasswordPageHandler(c *gin.Context) {
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

	// if global.IsUserCanAccess(user) {
	// 	notFoundUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/not-found"
	// 	http.Redirect(c.Writer, c.Request, notFoundUrl, http.StatusSeeOther)
	// 	return
	// }

	renderData := gin.H{
		"UserDetail": user,
		"config":     conf.GConfig,
	}

	if user.Data.Role == session.USER_TYPE_ADMIN_TREK {
		c.HTML(http.StatusOK, "admin-change-password.tmpl", renderData)
	} else {
		c.HTML(http.StatusOK, "dashboard-change-password.tmpl", renderData)
	}
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

	if user.Data.Role != session.USER_TYPE_ADMIN_TREK {
		notFoundUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/not-found"
		http.Redirect(c.Writer, c.Request, notFoundUrl, http.StatusSeeOther)
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

	if user.Data.Role != session.USER_TYPE_ADMIN_TREK {
		notFoundUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/not-found"
		http.Redirect(c.Writer, c.Request, notFoundUrl, http.StatusSeeOther)
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

	if user.Data.Role != session.USER_TYPE_ADMIN_TREK {
		notFoundUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/not-found"
		http.Redirect(c.Writer, c.Request, notFoundUrl, http.StatusSeeOther)
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

	if user.Data.Role != session.USER_TYPE_ADMIN_TREK {
		notFoundUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/not-found"
		http.Redirect(c.Writer, c.Request, notFoundUrl, http.StatusSeeOther)
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

	if user.Data.Role != session.USER_TYPE_ADMIN_TREK {
		notFoundUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/not-found"
		http.Redirect(c.Writer, c.Request, notFoundUrl, http.StatusSeeOther)
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

	if user.Data.Role != session.USER_TYPE_ADMIN_TREK {
		notFoundUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/not-found"
		http.Redirect(c.Writer, c.Request, notFoundUrl, http.StatusSeeOther)
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

	if user.Data.Role != session.USER_TYPE_ADMIN_TREK {
		notFoundUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/not-found"
		http.Redirect(c.Writer, c.Request, notFoundUrl, http.StatusSeeOther)
		return
	}

	renderData := gin.H{
		"UserDetail": user,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "admin-user-edit.tmpl", renderData)
}

func handleSessionErrorPage(c *gin.Context, sessionErr session.Error, needRedirect bool) {
	// invalid or expire session will force to delete cookie
	if sessionErr.Code == errorConst.ERROR_CODE_SESSION_EXPIRE || sessionErr.Code == errorConst.ERROR_CODE_TOKEN_SESSION {
		// set cookie to expire
		cookie := gSession.SetExpireSessionCookie()
		http.SetCookie(c.Writer, &cookie)

		if !needRedirect {
			renderData := gin.H{
				"config": conf.GConfig,
			}
			c.HTML(http.StatusOK, "login.tmpl", renderData)
		} else {
			loginURL := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
			http.Redirect(c.Writer, c.Request, loginURL, http.StatusSeeOther)
		}

		return
	} else if sessionErr.Code == errorConst.ERROR_CODE_USER_NOT_REGISTER_ON_COMPANY {
		// TODO: redirect to user not register on any company
		if !needRedirect {
			// ====================== REMOVE / NEED CHANGES ======================
			renderData := gin.H{
				"config": conf.GConfig,
			}
			c.HTML(http.StatusOK, "login.tmpl", renderData)
			// ====================================================================
		} else {
			loginURL := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
			http.Redirect(c.Writer, c.Request, loginURL, http.StatusSeeOther)
		}

		return
	} else {
		// TODO: Error Internal server - redirect to problem page
		if !needRedirect {
			// ====================== REMOVE / NEED CHANGES ======================
			renderData := gin.H{
				"config": conf.GConfig,
			}
			c.HTML(http.StatusOK, "login.tmpl", renderData)
			// ====================================================================
		} else {
			loginURL := conf.GConfig.BaseUrlConfig.BaseDNS + "/login"
			http.Redirect(c.Writer, c.Request, loginURL, http.StatusSeeOther)
		}

		return
	}
}
