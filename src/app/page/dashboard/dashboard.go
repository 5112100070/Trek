package dashboard

import (
	"net/http"
	"strconv"

	"github.com/5112100070/Trek/src/entity"

	"github.com/5112100070/Trek/src/app/session"
	"github.com/5112100070/Trek/src/app/user"
	"github.com/5112100070/Trek/src/conf"
	errorConst "github.com/5112100070/Trek/src/constants/error"
	roleConst "github.com/5112100070/Trek/src/constants/role"
	statusConst "github.com/5112100070/Trek/src/constants/status"
	"github.com/5112100070/Trek/src/global"
	gSession "github.com/5112100070/Trek/src/global/session"
	"github.com/gin-gonic/gin"
)

// MainPageHandler Dashboard page main handler
func MainPageHandler(c *gin.Context) {
	// Check user session
	accountResp, _, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println(errGetResponse)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	// TODO:
	// ONLY ROLE CGX ADMIN CAN USE THIS PAGE
	// if !(user.Data.Role == session.USER_TYPE_ADMIN || user.Data.Role == session.USER_TYPE_COMMON) {
	// 	notFoundUrl := conf.GConfig.BaseUrlConfig.BaseDNS + "/not-found"
	// 	http.Redirect(c.Writer, c.Request, notFoundUrl, http.StatusSeeOther)
	// 	return
	// }

	renderData := gin.H{
		"UserDetail": accountResp.Data,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "main-dashboard.tmpl", renderData)
}

// UserListPageHandler is handler for show list user for admin or company user
func UserListPageHandler(c *gin.Context) {
	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println(errGetResponse)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	// get list param
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	rows, _ := strconv.ParseInt(c.DefaultQuery("rows", "10"), 10, 64)
	orderType := c.DefaultQuery("order_type", "desc")

	listUserResp, err := global.GetServiceUser().GetListUsers(token, user.ListUserParam{
		Page:      int(page),
		Rows:      int(rows),
		OrderType: orderType,
	})
	if err != nil {
		// need internal page error handler
	}

	templatePage := conf.GConfig.BaseUrlConfig.BaseDNS + "/dashboard/users"
	totalPage := listUserResp.Data.Total / int(rows)

	// get additional page
	// total data 22 row 10
	// must result 3 pages
	if listUserResp.Data.Total%int(rows) > 0 {
		totalPage++
	}

	// wording for type user
	for i, u := range listUserResp.Data.Accounts {
		listUserResp.Data.Accounts[i].RoleWording = roleConst.ROLE_ACCOUNT_WORDING[u.Role]
	}

	// handle pagination
	pagination := entity.Pagination{
		Template:  templatePage,
		Page:      int(page),
		NextPage:  int(page) + 1,
		PrevPage:  int(page) - 1,
		Rows:      int(rows),
		TotalPage: totalPage,
		ListPage:  global.GenerateListPage(totalPage),
	}

	renderData := gin.H{
		"UserDetail": accountResp.Data,
		"accounts":   listUserResp.Data.Accounts,
		"pagination": pagination,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "list-user.tmpl", renderData)
}

// CompaniesListPageHandler is handler for show list company for admin or company user
func CompaniesListPageHandler(c *gin.Context) {
	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println(errGetResponse)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	// get list param
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	rows, _ := strconv.ParseInt(c.DefaultQuery("rows", "10"), 10, 64)
	orderType := c.DefaultQuery("order_type", "desc")

	listUserResp, err := global.GetServiceUser().GetListCompany(token, user.ListCompanyParam{
		Page:      int(page),
		Rows:      int(rows),
		OrderType: orderType,
	})
	if err != nil {
		// need internal page error handler
	}

	// define status activation company
	for i, c := range listUserResp.Data.Companies {
		listUserResp.Data.Companies[i].StatusActivation = statusConst.COMPANY_IS_ENABLED_WORDING[c.IsEnabled]
	}

	templatePage := conf.GConfig.BaseUrlConfig.BaseDNS + "/dashboard/companies"
	totalPage := listUserResp.Data.Total / int(rows)

	// get additional page
	// total data 22 row 10
	// must result 3 pages
	if listUserResp.Data.Total%int(rows) > 0 {
		totalPage++
	}

	// handle pagination
	pagination := entity.Pagination{
		Template:  templatePage,
		Page:      int(page),
		NextPage:  int(page) + 1,
		PrevPage:  int(page) - 1,
		Rows:      int(rows),
		TotalPage: totalPage,
		ListPage:  global.GenerateListPage(totalPage),
	}

	renderData := gin.H{
		"UserDetail": accountResp.Data,
		"companies":  listUserResp.Data.Companies,
		"pagination": pagination,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "list-companies.tmpl", renderData)
}

// getUserProfile -- get detail user based on active cookie
func getUserProfile(c *gin.Context) (session.AccountResponse, string, error) {
	var result session.AccountResponse
	var loginURL = conf.GConfig.BaseUrlConfig.BaseDNS + "/login"

	authToken, errGetCookie := c.Cookie(global.UserCookie[global.GetEnv()])
	if errGetCookie != nil {
		global.Error.Println(errGetCookie)
		http.Redirect(c.Writer, c.Request, loginURL, http.StatusSeeOther)
		return result, authToken, errGetCookie
	}

	result, errGetResponse := global.GetServiceSession().GetUser(authToken)
	if errGetResponse != nil {
		global.Error.Println(errGetResponse)
		// TODO: redirect to problem page
		// ====================== REMOVE / NEED CHANGES ======================
		global.Error.Println(errGetResponse)
		http.Redirect(c.Writer, c.Request, loginURL, http.StatusSeeOther)
		// ====================================================================
		return result, authToken, errGetResponse
	}

	return result, authToken, nil
}

// handleSessionErrorPage is mitigation error because expire or invalid session (not error because internal error)
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