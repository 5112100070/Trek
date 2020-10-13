package dashboard

import (
	"log"
	"net/http"

	"github.com/5112100070/Trek/src/app/user"

	"github.com/5112100070/Trek/src/app/order"
	"github.com/5112100070/Trek/src/conf"
	constErr "github.com/5112100070/Trek/src/constants/error"
	constRole "github.com/5112100070/Trek/src/constants/role"
	constUrl "github.com/5112100070/Trek/src/constants/url"
	"github.com/5112100070/Trek/src/global"
	"github.com/gin-gonic/gin"
)

func CreateOrderPageHandler(c *gin.Context) {
	// Check user session
	accountResp, sessionID, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.RenderInternalServerErrorPage(c)
		global.Error.Println("func CreateOrderPageHandler error when get profile: ", errGetResponse)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	// validate access to this feature
	featureCheckResp, err := global.GetServiceSession().CheckFeature(sessionID, constUrl.URL_ADMIN_CREATE_ORDER_FOR_ADMIN, http.MethodPost)
	if err != nil {
		global.Error.Println("func CompaniesListPageHandler error when check feature: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if featureCheckResp.Error != nil {
		handleErrorCheckFeature(c, featureCheckResp)
		return
	}

	unitsResp, errGetUnitData := global.GetServiceOrder().GetListUnitInOrder(sessionID)
	if errGetUnitData != nil {
		// add internal server error page response
		log.Println("func CreateOrderPageHandler error when call get list unit from CGX: ", errGetUnitData)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if unitsResp.Error != nil {
		log.Println("func CreateOrderPageHandler error expected when call get list unit from CGX: ", unitsResp.Error.Detail)
		global.RenderInternalServerErrorPage(c)
	}

	var listCompany []user.CompanyProfile
	if accountResp.Data.Company.Role != constRole.ROLE_COMPANY_CLIENT {
		companiesResp, errGetCompany := global.GetServiceUser().GetListCompany(sessionID, user.ListCompanyParam{
			Page:             1,
			Rows:             50,
			OrderType:        "desc",
			FilterByIsEnable: "1",
		})
		if errGetCompany != nil {
			// add internal server error page response
			log.Println("func CreateOrderPageHandler error when call get list unit to CGX: ", errGetCompany)
			global.RenderInternalServerErrorPage(c)
			return
		}

		if companiesResp.Error != nil {
			// possibility error
			if companiesResp.Error.Code == constErr.ERROR_CODE_SESSION_EXPIRE {
				// ERROR_CODE_SESSION_EXPIRE
				handleSessionErrorPage(c, *accountResp.Error, true)
			} else if companiesResp.Error.Code == constErr.ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS || companiesResp.Error.Code == constErr.ERROR_CODE_NOT_HAVE_RESULT {
				// ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS
				global.RenderUnAuthorizePage(c)
			} else if companiesResp.Error.Code == constErr.ERROR_CODE_INVALID_PARAMETER {
				// ERROR_CODE_INVALID_PARAMETER
				global.RenderNotFoundPage(c)
			} else {
				// ERROR_CODE_INTERNAL_SERVER
				global.RenderInternalServerErrorPage(c)
			}

			return
		}

		listCompany = companiesResp.Data.Companies
	} else {
		detailResp, err := global.GetServiceUser().GetDetailCompany(sessionID, accountResp.Data.Company.ID)
		if err != nil {
			global.Error.Println("func UserCreatePagehandler error get detail company: ", err)
			global.RenderInternalServerErrorPage(c)
			return
		}

		if detailResp.Error != nil {
			// possibility error
			if detailResp.Error.Code == constErr.ERROR_CODE_SESSION_EXPIRE {
				// ERROR_CODE_SESSION_EXPIRE
				handleSessionErrorPage(c, *accountResp.Error, true)
			} else if detailResp.Error.Code == constErr.ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS {
				// ERROR_CODE_ACCOUNT_NOT_HAVE_ACCESS
				global.RenderUnAuthorizePage(c)
			} else {
				// ERROR_CODE_INTERNAL_SERVER
				global.RenderInternalServerErrorPage(c)
			}
			return
		}

		listCompany = append(listCompany, detailResp.Data)
	}

	renderData := gin.H{
		"UserDetail": accountResp.Data,
		"Units":      unitsResp.Data,
		"Companies":  listCompany,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "create-order.tmpl", renderData)
}

// CreateOrderForAdmin method to create order only for admin
func CreateOrderForAdmin(c *gin.Context) {
	var body order.CreateOrderParam

	// Check user session
	_, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func CreateOrderForAdmin error get user profile: ", errGetResponse)
		global.ForbiddenResponse(c, constErr.WORDING_ERROR_FORBIDDEN)
		return
	}

	errBindJSON := c.BindJSON(&body)
	if errBindJSON != nil {
		global.Error.Printf("func CreateOrderForAdmin error when marshal json: %v\n", errBindJSON)
		global.BadRequestResponse(c, "invalid request")
		return
	}

	resp, err := global.GetServiceOrder().CreateOrderForAdmin(token, body)
	if err != nil {
		global.Error.Println("func CreateOrderForAdmin error when create order for admin: ", err)
		global.InternalServerErrorResponse(c, constErr.WORDING_ERROR_INTERNAL_SERVER)
		return
	}

	response := map[string]interface{}{
		"error": resp.Error,
		"data":  resp.Data,
	}

	global.OKResponse(c, response)

	return
}
