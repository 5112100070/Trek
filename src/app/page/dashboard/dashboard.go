package dashboard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/5112100070/Trek/src/app/order"
	"github.com/5112100070/Trek/src/app/session"
	"github.com/5112100070/Trek/src/conf"
	constErr "github.com/5112100070/Trek/src/constants/error"
	errorConst "github.com/5112100070/Trek/src/constants/error"
	statusConst "github.com/5112100070/Trek/src/constants/status"
	constUrl "github.com/5112100070/Trek/src/constants/url"
	"github.com/5112100070/Trek/src/global"
	gSession "github.com/5112100070/Trek/src/global/session"
	"github.com/gin-gonic/gin"
)

// MainPageHandler Dashboard page main handler
func MainPageHandler(c *gin.Context) {
	// Check user session
	accountResp, _, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func MainPageHandler error get user profile: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	if accountResp.Data.Company.ImageLogo == "" {
		accountResp.Data.Company.ImageLogo = "dashboard/assets/img/drawkit/color/drawkit-content-man-alt.svg"
	}

	renderData := gin.H{
		"UserDetail": accountResp.Data,
		"config":     conf.GConfig,
	}
	c.HTML(http.StatusOK, "main-dashboard.tmpl", renderData)
}

// ConfigPageHandler is handler for all configuration on account
func ConfigPageHandler(c *gin.Context) {
	// Check user session
	accountResp, _, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func ConfigPageHandler error get user profile: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	if accountResp.Data == nil {
		global.Error.Println("func ConfigPageHandler not have result from cgx. but not have error")
		global.RenderInternalServerErrorPage(c)
		return
	}

	accountRespJSON, _ := json.Marshal(accountResp.Data)

	renderData := gin.H{
		"UserDetail":     accountResp.Data,
		"UserDetailJSON": string(accountRespJSON),
		"config":         conf.GConfig,
	}
	c.HTML(http.StatusOK, "config.tmpl", renderData)
}

// OrdersDetailPageHandler is handler for show order detail for admin and company user
func OrdersDetailPageHandler(c *gin.Context) {
	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func OrdersDetailPageHandler error get detail account: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	// validate access to this feature
	featureCheckResp, err := global.GetServiceSession().CheckFeature(token, constUrl.URL_ADMIN_GET_ORDER_DETAIL, http.MethodGet)
	if err != nil {
		global.Error.Println("func OrdersDetailPageHandler error when check feature: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if featureCheckResp.Error != nil {
		handleErrorCheckFeature(c, featureCheckResp)
		return
	}

	// get list param
	orderID, _ := strconv.ParseInt(c.DefaultQuery("id", "1"), 10, 64)

	orderDetail, errCGX, err := global.GetServiceOrder().GetOrderDetailForAdmin(token, orderID)
	if err != nil || errCGX != nil {
		global.Error.Println("func OrdersDetailPageHandler error get order detail for admin: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	renderData := gin.H{
		"UserDetail":  accountResp.Data,
		"order":       orderDetail,
		"BadgeOrder":  statusConst.MAP_BADGE_BY_STATUS_ORDER,
		"BadgePickup": statusConst.MAP_BADGE_BY_STATUS_PICKUP,
		"config":      conf.GConfig,
	}
	c.HTML(http.StatusOK, "detail-order.tmpl", renderData)
}

// OrdersListPageHandler is handler for show list order for admin or company user
func OrdersListPageHandler(c *gin.Context) {
	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func OrdersListPageHandler error get detail account: ", errGetResponse)
		global.RenderInternalServerErrorPage(c)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		global.Error.Println("func OrdersListPageHandler error session: ", accountResp.Error.Detail)
		handleSessionErrorPage(c, *accountResp.Error, true)
		return
	}

	// validate access to this feature
	featureCheckResp, err := global.GetServiceSession().CheckFeature(token, constUrl.URL_ADMIN_GET_ORDER, http.MethodGet)
	if err != nil {
		global.Error.Println("func OrdersListPageHandler error when check feature: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	if featureCheckResp.Error != nil {
		handleErrorCheckFeature(c, featureCheckResp)
		return
	}

	// get list param
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if page <= 0 {
		page = 1
	}

	rows, _ := strconv.ParseInt(c.DefaultQuery("rows", "8"), 10, 64)
	orderType := c.DefaultQuery("order_type", "desc")
	companyID, _ := strconv.ParseInt(c.DefaultQuery("company_id", "0"), 10, 64)
	status, _ := strconv.Atoi(c.DefaultQuery("status", "-1"))
	createTimeFrom := c.DefaultQuery("create_time_from", "")
	createTimeTo := c.DefaultQuery("create_time_to", "")

	// check is valid parameter date
	if createTimeFrom != "" {
		_, errParse := time.Parse("2006-01-02", createTimeFrom)
		if errParse != nil {
			global.RenderNotFoundPage(c)
			return
		}
		createTimeFrom = createTimeFrom + " 00:00:00"
	}

	// check is valid parameter date
	if createTimeTo != "" {
		_, errParse := time.Parse("2006-01-02", createTimeTo)
		if errParse != nil {
			global.RenderNotFoundPage(c)
			return
		}
		createTimeTo = createTimeTo + " 23:59:59"
	}

	listOrderResp, err := global.GetServiceOrder().GetListOrders(token, order.ListOrderParam{
		Page:           int(page),
		Rows:           int(rows),
		OrderType:      orderType,
		CompanyID:      companyID,
		Status:         status,
		CreateTimeFrom: createTimeFrom,
		CreateTimeTo:   createTimeTo,
	})
	if err != nil {
		global.Error.Println("func OrdersListPageHandler error get list order: ", err)
		global.RenderInternalServerErrorPage(c)
		return
	}

	// fetching item data
	mapItems := make(map[int64][]order.ItemResponse)
	for _, orderDetail := range listOrderResp.Data.Orders {
		var items []order.ItemResponse
		for _, pickupDetail := range orderDetail.Pickups {
			for _, item := range pickupDetail.Items {
				items = append(items, item)
			}
		}
		mapItems[orderDetail.ID] = items
	}

	var filterBy string
	if companyID > 0 {
		filterBy = fmt.Sprintf("&company_id=%d", companyID)
	}

	accountRespJSON, _ := json.Marshal(accountResp.Data)

	renderData := gin.H{
		"UserDetail":     accountResp.Data,
		"UserDetailJSON": string(accountRespJSON),
		"orders":         listOrderResp.Data.Orders,
		"next":           listOrderResp.Data.Next,
		"mapItems":       mapItems,
		"config":         conf.GConfig,
		"filterBy":       filterBy,
	}
	c.HTML(http.StatusOK, "list-orders.tmpl", renderData)
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
		// redirect to user not register on any company
		global.RenderUnAuthorizePage(c)
		return
	} else {
		// Error Internal server - redirect to problem page
		global.RenderInternalServerErrorPage(c)
		return
	}
}

func handleErrorCheckFeature(c *gin.Context, featureCheckResp session.FeatureCheckResponse) {
	if featureCheckResp.Error.Code == constErr.ERROR_CODE_NOT_HAVE_PERMISSION_ON_FEATURE {
		global.RenderUnAuthorizePage(c)
		return
	} else if featureCheckResp.Error.Code == constErr.ERROR_CODE_SESSION_EXPIRE {
		handleSessionErrorPage(c, *featureCheckResp.Error, true)
		return
	} else {
		global.RenderInternalServerErrorPage(c)
		return
	}
}
