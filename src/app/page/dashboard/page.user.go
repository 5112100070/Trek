package dashboard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/5112100070/Trek/src/entity"

	"github.com/5112100070/Trek/src/app/user"
	"github.com/5112100070/Trek/src/conf"
	constErr "github.com/5112100070/Trek/src/constants/error"
	constRole "github.com/5112100070/Trek/src/constants/role"
	constUrl "github.com/5112100070/Trek/src/constants/url"
	"github.com/5112100070/Trek/src/global"
	"github.com/gin-gonic/gin"
)

// UserListPageHandler is handler for show list user for admin or company user
func UserListPageHandler(c *gin.Context) {
	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func UserListPageHandler error get user profile: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	// validate access to this feature
	featureCheckResp, err := global.GetServiceSession().CheckFeature(token, constUrl.URL_ADMIN_GET_LIST_USER, http.MethodGet)
	if err != nil {
		global.Error.Println("func UserListPageHandler error when check feature: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if featureCheckResp.Error != nil {
		handleErrorCheckFeature(c, featureCheckResp)
		return
	}

	// get list param
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	rows, _ := strconv.ParseInt(c.DefaultQuery("rows", "10"), 10, 64)
	companyID, _ := strconv.ParseInt(c.DefaultQuery("company_id", "0"), 10, 64)
	orderType := c.DefaultQuery("order_type", "desc")

	listUserResp, err := global.GetServiceUser().GetListUsers(token, user.ListUserParam{
		Page:      int(page),
		Rows:      int(rows),
		OrderType: orderType,
		CompanyID: companyID,
	})
	if err != nil {
		global.Error.Println("func UserListPageHandler error get list user: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if listUserResp.Error != nil {
		// possibility error
		if listUserResp.Error.Code == constErr.ERROR_CODE_SESSION_EXPIRE {
			// ERROR_CODE_SESSION_EXPIRE
			handleSessionErrorPage(c, *accountResp.Error, true)
		} else if listUserResp.Error.Code == constErr.ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS || listUserResp.Error.Code == constErr.ERROR_CODE_NOT_HAVE_RESULT {
			// ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS
			global.RenderUnAuthorizePage(c)
		} else if listUserResp.Error.Code == constErr.ERROR_CODE_INVALID_PARAMETER {
			// ERROR_CODE_INVALID_PARAMETER
			global.RenderNotFoundPage(c)
		} else {
			// ERROR_CODE_INTERNAL_SERVER
			global.RenderInternalServerErrorPage(c)
		}

		return
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
		listUserResp.Data.Accounts[i].RoleWording = constRole.ROLE_ACCOUNT_WORDING[u.Role]
		listUserResp.Data.Accounts[i].RoleColor = constRole.ROLE_ACCOUNT_COLOR[u.Role]
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

	accountRespJSON, _ := json.Marshal(accountResp.Data)

	var filterBy string
	if companyID > 0 {
		filterBy = fmt.Sprintf("&company_id=%d", companyID)
	}

	renderData := gin.H{
		"UserDetail":     accountResp.Data,
		"UserDetailJSON": string(accountRespJSON),
		"accounts":       listUserResp.Data.Accounts,
		"pagination":     pagination,
		"config":         conf.GConfig,
		"filterBy":       filterBy,
	}
	c.HTML(http.StatusOK, "list-user.tmpl", renderData)
}

// UserDetailPageHandler is handler for show detail company for admin
func UserDetailPageHandler(c *gin.Context) {
	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func UserDetailPageHandler error get user profile: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	// validate access to this feature
	featureCheckResp, err := global.GetServiceSession().CheckFeature(token, constUrl.URL_ADMIN_GET_DETAIL_ACCOUNT, http.MethodGet)
	if err != nil {
		global.Error.Println("func UserDetailPageHandler error when check feature: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if featureCheckResp.Error != nil {
		handleErrorCheckFeature(c, featureCheckResp)
		return
	}

	// get list param
	accountID, _ := strconv.ParseInt(c.DefaultQuery("id", "1"), 10, 64)

	accDetail, err := global.GetServiceUser().GetDetailAccount(token, accountID)
	if err != nil {
		global.Error.Println("func UserDetailPageHandler error get detail account: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if accDetail.Error != nil {
		// possibility error
		if accDetail.Error.Code == constErr.ERROR_CODE_SESSION_EXPIRE {
			// ERROR_CODE_SESSION_EXPIRE
			handleSessionErrorPage(c, *accountResp.Error, true)
		} else if accDetail.Error.Code == constErr.ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS || accDetail.Error.Code == constErr.ERROR_CODE_NOT_HAVE_RESULT {
			// ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS
			global.RenderUnAuthorizePage(c)
		} else if accDetail.Error.Code == constErr.ERROR_CODE_INVALID_PARAMETER {
			// ERROR_CODE_INVALID_PARAMETER
			global.RenderNotFoundPage(c)
		} else {
			// ERROR_CODE_INTERNAL_SERVER
			global.RenderInternalServerErrorPage(c)
		}

		return
	}

	accountRespJSON, _ := json.Marshal(accountResp.Data)

	renderData := gin.H{
		"UserDetail":     accountResp.Data,
		"UserDetailJSON": string(accountRespJSON),
		"account":        accDetail.Data,
		"config":         conf.GConfig,
	}
	c.HTML(http.StatusOK, "detail-user.tmpl", renderData)
}

// UserCreatePagehandler is handler for show form create user
func UserCreatePagehandler(c *gin.Context) {
	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func UserCreatePagehandler error get list user: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	// validate access to this feature
	featureCheckResp, err := global.GetServiceSession().CheckFeature(token, constUrl.URL_ADMIN_CREATE_ACCOUNT, http.MethodPost)
	if err != nil {
		global.Error.Println("func UserCreatePagehandler error when check feature: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if featureCheckResp.Error != nil {
		handleErrorCheckFeature(c, featureCheckResp)
		return
	}

	var listCompany []user.CompanyProfile
	if accountResp.Data.Company.Role != constRole.ROLE_COMPANY_CLIENT {
		listCompResp, err := global.GetServiceUser().GetListCompany(token, user.ListCompanyParam{
			Page:             1,
			Rows:             20,
			OrderType:        "desc",
			FilterByIsEnable: "1",
			IsInternal:       true,
		})
		if err != nil {
			global.Error.Println("func UserCreatePagehandler error get list company: ", err)
			global.RenderInternalServerErrorPage(c)
			return
		}

		if listCompResp.Error != nil {
			// possibility error
			if listCompResp.Error.Code == constErr.ERROR_CODE_SESSION_EXPIRE {
				// ERROR_CODE_SESSION_EXPIRE
				handleSessionErrorPage(c, *accountResp.Error, true)
			} else if listCompResp.Error.Code == constErr.ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS || listCompResp.Error.Code == constErr.ERROR_CODE_NOT_HAVE_PERMISSION_ON_FEATURE {
				// ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS
				global.RenderUnAuthorizePage(c)
			} else {
				// ERROR_CODE_INTERNAL_SERVER
				global.RenderInternalServerErrorPage(c)
			}

			return
		}

		listCompany = listCompResp.Data.Companies
	} else {
		// for client only can create account for theirselves
		listCompany = append(listCompany, user.CompanyProfile{
			ID:          accountResp.Data.Company.ID,
			CompanyName: accountResp.Data.Company.CompanyName,
		})
	}

	renderData := gin.H{
		"UserDetail": accountResp.Data,
		"companies":  listCompany,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "create-user.tmpl", renderData)
}

// UserUpdatePagehandler is handler for show update user form
func UserUpdatePagehandler(c *gin.Context) {
	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func UserUpdatePagehandler error get user profile: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	// validate access to this feature
	featureCheckResp, err := global.GetServiceSession().CheckFeature(token, constUrl.URL_ADMIN_UPDATE_ACCOUNT, http.MethodPost)
	if err != nil {
		global.Error.Println("func UserCreatePagehandler error when check feature: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if featureCheckResp.Error != nil {
		handleErrorCheckFeature(c, featureCheckResp)
		return
	}

	accountID, _ := strconv.ParseInt(c.DefaultQuery("id", "1"), 10, 64)
	accDetail, err := global.GetServiceUser().GetDetailAccount(token, accountID)
	if err != nil {
		global.Error.Println("func UserUpdatePagehandler error get detail account: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if accDetail.Error != nil {
		// possibility error
		if accDetail.Error.Code == constErr.ERROR_CODE_SESSION_EXPIRE {
			// ERROR_CODE_SESSION_EXPIRE
			handleSessionErrorPage(c, *accountResp.Error, true)
		} else if accDetail.Error.Code == constErr.ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS || accDetail.Error.Code == constErr.ERROR_CODE_NOT_HAVE_RESULT {
			// ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS
			global.RenderUnAuthorizePage(c)
		} else if accDetail.Error.Code == constErr.ERROR_CODE_INVALID_PARAMETER {
			// ERROR_CODE_INVALID_PARAMETER
			global.RenderNotFoundPage(c)
		} else {
			// ERROR_CODE_INTERNAL_SERVER
			global.RenderInternalServerErrorPage(c)
		}

		return
	}

	renderData := gin.H{
		"UserDetail": accountResp.Data,
		"account":    accDetail.Data,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "update-user.tmpl", renderData)
}
