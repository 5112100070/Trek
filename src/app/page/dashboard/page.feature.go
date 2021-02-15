package dashboard

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/5112100070/Trek/src/entity"

	"github.com/5112100070/Trek/src/app/module"
	"github.com/5112100070/Trek/src/conf"
	constErr "github.com/5112100070/Trek/src/constants/error"
	"github.com/5112100070/Trek/src/global"
	"github.com/gin-gonic/gin"
)

// ModulesPageHandler is handler for show list module which we have
func ModulesPageHandler(c *gin.Context) {
	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func ModulesPageHandler error get user profile: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
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

	listModules, err := global.GetServiceModule().GetListModule(token, module.ListModuleParam{
		Page:      int(page),
		Rows:      int(rows),
		OrderType: orderType,
	})
	if err != nil {
		global.Error.Println("func ModulesPageHandler error get list user: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if listModules.Error != nil {
		// possibility error
		if listModules.Error.Code == constErr.ERROR_CODE_SESSION_EXPIRE {
			// ERROR_CODE_SESSION_EXPIRE
			handleSessionErrorPage(c, *accountResp.Error, true)
		} else if listModules.Error.Code == constErr.ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS || listModules.Error.Code == constErr.ERROR_CODE_NOT_HAVE_RESULT {
			// ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS
			global.RenderUnAuthorizePage(c)
		} else if listModules.Error.Code == constErr.ERROR_CODE_INVALID_PARAMETER {
			// ERROR_CODE_INVALID_PARAMETER
			global.RenderNotFoundPage(c)
		} else {
			// ERROR_CODE_INTERNAL_SERVER
			global.RenderInternalServerErrorPage(c)
		}

		return
	}

	templatePage := conf.GConfig.BaseUrlConfig.BaseDNS + "/dashboard/modules"
	totalPage := listModules.Data.Total / int(rows)

	// get additional page
	// total data 22 row 10
	// must result 3 pages
	if listModules.Data.Total%int(rows) > 0 {
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
		"modules":    listModules.Data.Modules,
		"pagination": pagination,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "list-module.tmpl", renderData)
}

// FeaturePageHandler is handler for show list feature which we have
func FeaturePageHandler(c *gin.Context) {
	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func FeaturePageHandler error get user profile: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
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
	moduleID, _ := strconv.ParseInt(c.DefaultQuery("module_id", "0"), 10, 64)

	listModules, err := global.GetServiceModule().GetListFeature(token, module.ListFeatureParam{
		Page:      int(page),
		Rows:      int(rows),
		ModuleID:  moduleID,
		OrderType: orderType,
	})
	if err != nil {
		global.Error.Println("func FeaturePageHandler error get list user: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if listModules.Error != nil {
		// possibility error
		if listModules.Error.Code == constErr.ERROR_CODE_SESSION_EXPIRE {
			// ERROR_CODE_SESSION_EXPIRE
			handleSessionErrorPage(c, *accountResp.Error, true)
		} else if listModules.Error.Code == constErr.ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS || listModules.Error.Code == constErr.ERROR_CODE_NOT_HAVE_RESULT {
			// ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS
			global.RenderUnAuthorizePage(c)
		} else if listModules.Error.Code == constErr.ERROR_CODE_INVALID_PARAMETER {
			// ERROR_CODE_INVALID_PARAMETER
			global.RenderNotFoundPage(c)
		} else {
			// ERROR_CODE_INTERNAL_SERVER
			global.RenderInternalServerErrorPage(c)
		}

		return
	}

	totalPage := listModules.Data.Total / int(rows)

	Url, _ := url.Parse("")
	Url.Path += "/dashboard/features"
	parameters := url.Values{}
	parameters.Add("rows", strconv.FormatInt(rows, 10))
	if moduleID > 0 {
		parameters.Add("module_id", strconv.FormatInt(moduleID, 10))
	}
	Url.RawQuery = parameters.Encode()

	// get additional page
	// total data 22 row 10
	// must result 3 pages
	if listModules.Data.Total%int(rows) > 0 {
		totalPage++
	}

	// handle pagination
	pagination := entity.Pagination{
		Template:  Url.String(),
		Page:      int(page),
		NextPage:  int(page) + 1,
		PrevPage:  int(page) - 1,
		TotalPage: totalPage,
		ListPage:  global.GenerateListPage(totalPage),
	}

	renderData := gin.H{
		"UserDetail": accountResp.Data,
		"features":   listModules.Data.Features,
		"pagination": pagination,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "list-feature.tmpl", renderData)
}

// MemberFeaturePageHandler is handler for show list feature which we have
func MemberFeaturePageHandler(c *gin.Context) {
	// Check user session
	accountResp, _, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func MemberFeaturePageHandler error get user profile: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	renderData := gin.H{
		"UserDetail": accountResp.Data,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "list-member-feature.tmpl", renderData)
}
