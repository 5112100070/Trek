package dashboard

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/5112100070/Trek/src/entity"

	"github.com/5112100070/Trek/src/app/user"
	"github.com/5112100070/Trek/src/conf"
	constErr "github.com/5112100070/Trek/src/constants/error"
	constRole "github.com/5112100070/Trek/src/constants/role"
	statusConst "github.com/5112100070/Trek/src/constants/status"
	constUrl "github.com/5112100070/Trek/src/constants/url"
	"github.com/5112100070/Trek/src/global"
	"github.com/gin-gonic/gin"
)

// CompaniesListPageHandler is handler for show list company for admin or company user
func CompaniesListPageHandler(c *gin.Context) {
	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func CompaniesListPageHandler error get detail account: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	// validate access to this feature
	featureCheckResp, err := global.GetServiceSession().CheckFeature(token, constUrl.URL_ADMIN_GET_LIST_COMPANY, http.MethodGet)
	if err != nil {
		global.Error.Println("func CompaniesListPageHandler error when check feature: ", err)
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
	orderType := c.DefaultQuery("order_type", "desc")
	isEnabled := c.Query("is_enabled")

	listCompanyResp, err := global.GetServiceUser().GetListCompany(token, user.ListCompanyParam{
		Page:             int(page),
		Rows:             int(rows),
		OrderType:        orderType,
		FilterByIsEnable: isEnabled,
	})
	if err != nil {
		global.Error.Println("func CompaniesListPageHandler error get list company: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if listCompanyResp.Error != nil {
		// possibility error
		if listCompanyResp.Error.Code == constErr.ERROR_CODE_SESSION_EXPIRE {
			// ERROR_CODE_SESSION_EXPIRE
			handleSessionErrorPage(c, *accountResp.Error, true)
		} else if listCompanyResp.Error.Code == constErr.ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS || listCompanyResp.Error.Code == constErr.ERROR_CODE_NOT_HAVE_RESULT {
			// ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS
			global.RenderUnAuthorizePage(c)
		} else if listCompanyResp.Error.Code == constErr.ERROR_CODE_INVALID_PARAMETER {
			// ERROR_CODE_INVALID_PARAMETER
			global.RenderNotFoundPage(c)
		} else {
			// ERROR_CODE_INTERNAL_SERVER
			global.RenderInternalServerErrorPage(c)
		}

		return
	}

	// define status activation company
	for i, c := range listCompanyResp.Data.Companies {
		listCompanyResp.Data.Companies[i].StatusActivation = statusConst.COMPANY_IS_ENABLED_WORDING[c.IsEnabled]
	}

	templatePage := conf.GConfig.BaseUrlConfig.BaseDNS + "/dashboard/companies"
	totalPage := listCompanyResp.Data.Total / int(rows)

	// get additional page
	// total data 22 row 10
	// must result 3 pages
	if listCompanyResp.Data.Total%int(rows) > 0 {
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
		"companies":  listCompanyResp.Data.Companies,
		"pagination": pagination,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "list-companies.tmpl", renderData)
}

// CompanyDetailPageHandler is handler for show detail company for admin
func CompanyDetailPageHandler(c *gin.Context) {
	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func CompanyDetailPageHandler error get detail account: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	// validate access to this feature
	featureCheckResp, err := global.GetServiceSession().CheckFeature(token, constUrl.URL_ADMIN_GET_DETAIL_COMPANY, http.MethodGet)
	if err != nil {
		global.Error.Println("func CompanyDetailPageHandler error when check feature: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if featureCheckResp.Error != nil {
		handleErrorCheckFeature(c, featureCheckResp)
		return
	}

	// get list param
	companyID, _ := strconv.ParseInt(c.DefaultQuery("id", "1"), 10, 64)

	companyDetail, err := global.GetServiceUser().GetDetailCompany(token, user.DetailCompanyParam{
		CompanyID:  companyID,
		IsInternal: true,
	})
	if err != nil {
		global.Error.Println("func CompanyDetailPageHandler error get detail company: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if companyDetail.Error != nil {
		// possibility error
		// ERROR_CODE_INTERNAL_SERVER
		if companyDetail.Error.Code == constErr.ERROR_CODE_SESSION_EXPIRE {
			// ERROR_CODE_SESSION_EXPIRE
			handleSessionErrorPage(c, *accountResp.Error, true)
		} else if companyDetail.Error.Code == constErr.ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS || companyDetail.Error.Code == constErr.ERROR_CODE_NOT_HAVE_RESULT {
			// ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS
			global.RenderUnAuthorizePage(c)
		} else if companyDetail.Error.Code == constErr.ERROR_CODE_INVALID_PARAMETER && companyDetail.Error.Code == constErr.ERROR_CODE_NOT_HAVE_RESULT {
			// ERROR_CODE_INVALID_PARAMETER
			global.RenderNotFoundPage(c)
		} else {
			// ERROR_CODE_INTERNAL_SERVER
			global.RenderInternalServerErrorPage(c)
		}

		return
	}

	companyDetail.Data.StatusActivation = statusConst.COMPANY_IS_ENABLED_WORDING[companyDetail.Data.IsEnabled]
	companyDetail.Data.CompanyName = strings.ToUpper(companyDetail.Data.CompanyName)
	companyDetail.Data.StatusActivation = strings.ToUpper(companyDetail.Data.StatusActivation)

	accountRespJSON, _ := json.Marshal(accountResp.Data)

	renderData := gin.H{
		"UserDetail":     accountResp.Data,
		"UserDetailJSON": string(accountRespJSON),
		"company":        companyDetail.Data,
		"config":         conf.GConfig,
	}
	c.HTML(http.StatusOK, "detail-company.tmpl", renderData)
}

// func CompanyCreatePagehandler is handler for show form create company
func CompanyCreatePagehandler(c *gin.Context) {
	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func CompanyCreatePagehandler error get detail account: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	// validate access to this feature
	featureCheckResp, err := global.GetServiceSession().CheckFeature(token, constUrl.URL_ADMIN_CREATE_COMPANY, http.MethodPost)
	if err != nil {
		global.Error.Println("func CompanyCreatePagehandler error when check feature: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if featureCheckResp.Error != nil {
		handleErrorCheckFeature(c, featureCheckResp)
		return
	}

	// all client company cannot get this page
	if accountResp.Data.Company.Role == constRole.ROLE_COMPANY_CLIENT {
		global.RenderUnAuthorizePage(c)
		return
	}

	// only admin on admin company cannot get this page
	if accountResp.Data.Company.Role != constRole.ROLE_COMPANY_CLIENT {
		if accountResp.Data.Role != constRole.ROLE_ACCOUNT_ADMIN {
			global.RenderUnAuthorizePage(c)
			return
		}
	}

	accountRespJSON, _ := json.Marshal(accountResp.Data)

	renderData := gin.H{
		"UserDetail":     accountResp.Data,
		"UserDetailJSON": string(accountRespJSON),
		"config":         conf.GConfig,
		"IsGodUser":      accountResp.Data.Company.Role == constRole.ROLE_COMPANY_GOD,
	}
	c.HTML(http.StatusOK, "create-company.tmpl", renderData)
}

// CompanyUpdatePagehandler is handler for show companya update for admin
func CompanyUpdatePagehandler(c *gin.Context) {
	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func CompanyUpdatePagehandler error get detail account: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	// validate access to this feature
	featureCheckResp, err := global.GetServiceSession().CheckFeature(token, constUrl.URL_ADMIN_UPDATE_COMPANY, http.MethodPost)
	if err != nil {
		global.Error.Println("func CompanyUpdatePagehandler error when check feature: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if featureCheckResp.Error != nil {
		handleErrorCheckFeature(c, featureCheckResp)
		return
	}

	companyID, _ := strconv.ParseInt(c.DefaultQuery("id", "1"), 10, 64)
	accDetail, err := global.GetServiceUser().GetDetailCompany(token, user.DetailCompanyParam{
		CompanyID:  companyID,
		IsInternal: true,
	})
	if err != nil {
		global.Error.Println("func CompanyUpdatePagehandler error get detail company: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	renderData := gin.H{
		"UserDetail": accountResp.Data,
		"company":    accDetail.Data,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "update-company.tmpl", renderData)
}
