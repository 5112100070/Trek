package dashboard

import (
	"log"
	"net/http"
	"strconv"

	"github.com/5112100070/Trek/src/app/user"

	"github.com/5112100070/Trek/src/app/order"
	"github.com/5112100070/Trek/src/conf"
	constErr "github.com/5112100070/Trek/src/constants/error"
	constRole "github.com/5112100070/Trek/src/constants/role"
	constStatus "github.com/5112100070/Trek/src/constants/status"
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
		"response": resp,
	}

	global.OKResponse(c, response)

	return
}

// ApproveOrderForAdmin method to create order only for admin
func ApproveOrderForAdmin(c *gin.Context) {
	// Check user session
	_, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func ApproveOrderForAdmin error get user profile: ", errGetResponse)
		global.ForbiddenResponse(c, constErr.WORDING_ERROR_FORBIDDEN)
		return
	}

	orderID, errParse := strconv.ParseInt(c.PostForm("order_id"), 10, 64)
	if errParse != nil {
		global.Error.Printf("func ApproveOrderForAdmin error when parsing : %v\n", errParse)
		global.BadRequestResponse(c, "invalid request")
		return
	}

	awb := c.PostForm("awb")

	resp, err := global.GetServiceOrder().ApproveOrderForAdmin(token, orderID, awb)
	if err != nil {
		global.Error.Println("func ApproveOrderForAdmin error when create order for admin: ", err)
		global.InternalServerErrorResponse(c, constErr.WORDING_ERROR_INTERNAL_SERVER)
		return
	}

	response := map[string]interface{}{
		"response": resp,
	}

	global.OKResponse(c, response)

	return
}

// RejectOrderForAdmin method to create order only for admin
func RejectOrderForAdmin(c *gin.Context) {
	// Check user session
	_, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func RejectOrderForAdmin error get user profile: ", errGetResponse)
		global.ForbiddenResponse(c, constErr.WORDING_ERROR_FORBIDDEN)
		return
	}

	orderID, errParse := strconv.ParseInt(c.PostForm("order_id"), 10, 64)
	if errParse != nil {
		global.Error.Printf("func RejectOrderForAdmin error when parsing : %v\n", errParse)
		global.BadRequestResponse(c, "invalid request")
		return
	}

	resp, err := global.GetServiceOrder().RejectOrderForAdmin(token, orderID)
	if err != nil {
		global.Error.Println("func RejectOrderForAdmin error when create order for admin: ", err)
		global.InternalServerErrorResponse(c, constErr.WORDING_ERROR_INTERNAL_SERVER)
		return
	}

	response := map[string]interface{}{
		"response": resp,
	}

	global.OKResponse(c, response)

	return
}

// DispatchOrder method to dispatch order
func DispatchOrder(c *gin.Context) {
	// Check user session
	_, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func DispatchOrderForAdmin error get user profile: ", errGetResponse)
		global.ForbiddenResponse(c, constErr.WORDING_ERROR_FORBIDDEN)
		return
	}

	orderID, errParse := strconv.ParseInt(c.PostForm("order_id"), 10, 64)
	if errParse != nil {
		global.Error.Printf("func DispatchOrderForAdmin error when parsing : %v\n", errParse)
		global.BadRequestResponse(c, "invalid request")
		return
	}

	// Check is action is allowed dispatch status
	action := c.PostForm("action")
	if !constStatus.ALLOWED_STATUS_DISPATCH[action] {
		global.Error.Printf("func DispatchOrderForAdmin error when parsing : %v\n", errParse)
		global.BadRequestResponse(c, "invalid request")
		return
	}

	var resp *order.SuccessCRUDResponse
	var err error
	if action == constStatus.STATUS_ORDER_DISPATCH_TO_FULFILMENT_CENTER {
		resp, err = global.GetServiceOrder().DispatchOrderToFulfilmentCenter(token, orderID)
		if err != nil {
			global.Error.Println("func DispatchOrderForAdmin error when dispatch order to fulfilment center: ", err)
			global.InternalServerErrorResponse(c, constErr.WORDING_ERROR_INTERNAL_SERVER)
			return
		}
	} else if action == constStatus.STATUS_ORDER_DISPATCH_TO_DRIVER {
		resp, err = global.GetServiceOrder().DispatchOrderToDriver(token, orderID)
		if err != nil {
			global.Error.Println("func DispatchOrderForAdmin error when dispatch order to fulfilment center: ", err)
			global.InternalServerErrorResponse(c, constErr.WORDING_ERROR_INTERNAL_SERVER)
			return
		}
	} else {
		global.Error.Println("func DispatchOrderForAdmin undefined action")
		global.BadRequestResponse(c, nil)
		return
	}

	// prevent any panic on frontend process rendering
	if resp == nil {
		global.Error.Println("func DispatchOrderForAdmin error not have response from cgx: ", err)
		global.InternalServerErrorResponse(c, constErr.WORDING_ERROR_INTERNAL_SERVER)
		return
	}

	response := map[string]interface{}{
		"response": resp,
	}

	global.OKResponse(c, response)

	return
}

// PickUpItem method to pick up item in order. specific for 1 order 1 pickup id
func PickUpItem(c *gin.Context) {
	// Check user session
	_, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func PickUpItem error get user profile: ", errGetResponse)
		global.ForbiddenResponse(c, constErr.WORDING_ERROR_FORBIDDEN)
		return
	}

	orderID, errParse := strconv.ParseInt(c.PostForm("order_id"), 10, 64)
	if errParse != nil {
		global.Error.Printf("func PickUpItem error when parsing : %v\n", errParse)
		global.BadRequestResponse(c, "invalid order id")
		return
	}

	pickupID, errParse := strconv.ParseInt(c.PostForm("pickup_id"), 10, 64)
	if errParse != nil {
		global.Error.Printf("func PickUpItem error when parsing : %v\n", errParse)
		global.BadRequestResponse(c, "invalid pickup ID")
		return
	}

	param := order.PickUpParam{
		PickUpIDs: []int64{
			pickupID,
		},
		DriverName:  c.PostForm("driver_name"),
		DriverPhone: c.PostForm("driver_phone"),
	}

	resp, err := global.GetServiceOrder().PickUpOrderToDriver(token, orderID, param)
	if err != nil {
		global.Error.Println("func PickUpItem error when dispatch order to fulfilment center: ", err)
		global.InternalServerErrorResponse(c, constErr.WORDING_ERROR_INTERNAL_SERVER)
		return
	}

	// prevent any panic on frontend process rendering
	if resp == nil {
		global.Error.Println("func PickUpItem error not have response from cgx: ", err)
		global.InternalServerErrorResponse(c, constErr.WORDING_ERROR_INTERNAL_SERVER)
		return
	}

	response := map[string]interface{}{
		"response": resp,
	}

	global.OKResponse(c, response)

	return
}

// RejectPickUpItem method to reject pick up based on pick up id
func RejectPickUpItem(c *gin.Context) {
	// Check user session
	_, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func RejectPickUpItem error get user profile: ", errGetResponse)
		global.ForbiddenResponse(c, constErr.WORDING_ERROR_FORBIDDEN)
		return
	}

	orderID, errParse := strconv.ParseInt(c.PostForm("order_id"), 10, 64)
	if errParse != nil {
		global.Error.Printf("func RejectPickUpItem error when parsing : %v\n", errParse)
		global.BadRequestResponse(c, "invalid order id")
		return
	}

	pickupID, errParse := strconv.ParseInt(c.PostForm("pickup_id"), 10, 64)
	if errParse != nil {
		global.Error.Printf("func RejectPickUpItem error when parsing : %v\n", errParse)
		global.BadRequestResponse(c, "invalid pickup ID")
		return
	}

	resp, err := global.GetServiceOrder().RejectPickUpOrder(token, orderID, pickupID)
	if err != nil {
		global.Error.Println("func RejectPickUpItem error when dispatch order to fulfilment center: ", err)
		global.InternalServerErrorResponse(c, constErr.WORDING_ERROR_INTERNAL_SERVER)
		return
	}

	// prevent any panic on frontend process rendering
	if resp == nil {
		global.Error.Println("func RejectPickUpItem error not have response from cgx: ", err)
		global.InternalServerErrorResponse(c, constErr.WORDING_ERROR_INTERNAL_SERVER)
		return
	}

	response := map[string]interface{}{
		"response": resp,
	}

	global.OKResponse(c, response)

	return
}

// FinishPickUpItem method to finish pick up
func FinishPickUpItem(c *gin.Context) {
	// Check user session
	_, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func FinishPickUpItem error get user profile: ", errGetResponse)
		global.ForbiddenResponse(c, constErr.WORDING_ERROR_FORBIDDEN)
		return
	}

	var requestParam struct {
		OrderID   int64                   `json:"order_id"`
		PickUpIDs []int64                 `json:"pickup_ids"`
		Items     []order.ItemPickUpParam `json:"items"`
	}

	errBindRequest := c.BindJSON(&requestParam)
	if errBindRequest != nil {
		global.Error.Printf("func FinishPickUpItem error when parsing : %v\n", errBindRequest)
		global.BadRequestResponse(c, "invalid parameter request")
		return
	}

	var param order.FinishPickupParam
	param.PickUpIDs = requestParam.PickUpIDs
	param.Items = requestParam.Items

	resp, err := global.GetServiceOrder().FinishPickUpOrder(token, requestParam.OrderID, param)
	if err != nil {
		global.Error.Println("func FinishPickUpItem error when dispatch order to fulfilment center: ", err)
		global.InternalServerErrorResponse(c, constErr.WORDING_ERROR_INTERNAL_SERVER)
		return
	}

	// prevent any panic on frontend process rendering
	if resp == nil {
		global.Error.Println("func FinishPickUpItem error not have response from cgx: ", err)
		global.InternalServerErrorResponse(c, constErr.WORDING_ERROR_INTERNAL_SERVER)
		return
	}

	response := map[string]interface{}{
		"response": resp,
	}

	global.OKResponse(c, response)

	return
}

// DeliveryOrder method to handle delivery process order
func DeliveryOrder(c *gin.Context) {
	// Check user session
	_, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func DeliveryOrder error get user profile: ", errGetResponse)
		global.ForbiddenResponse(c, constErr.WORDING_ERROR_FORBIDDEN)
		return
	}

	var requestParam struct {
		OrderID     int64                   `json:"order_id"`
		DriverName  string                  `json:"driver_name"`
		DriverPhone string                  `json:"driver_phone"`
		Items       []order.ItemPickUpParam `json:"items"`
	}

	errBindRequest := c.BindJSON(&requestParam)
	if errBindRequest != nil {
		global.Error.Printf("func DeliveryOrder error when parsing : %v\n", errBindRequest)
		global.BadRequestResponse(c, "invalid parameter request")
		return
	}

	var param order.DeliveryParam
	param.DriverName = requestParam.DriverName
	param.DriverPhone = requestParam.DriverPhone
	param.Items = requestParam.Items

	resp, err := global.GetServiceOrder().DeliveryOrder(token, requestParam.OrderID, param)
	if err != nil {
		global.Error.Println("func DeliveryOrder error when set process order to deliver step: ", err)
		global.InternalServerErrorResponse(c, constErr.WORDING_ERROR_INTERNAL_SERVER)
		return
	}

	// prevent any panic on frontend process rendering
	if resp == nil {
		global.Error.Println("func DeliveryOrder error not have response from cgx: ", err)
		global.InternalServerErrorResponse(c, constErr.WORDING_ERROR_INTERNAL_SERVER)
		return
	}

	response := map[string]interface{}{
		"response": resp,
	}

	global.OKResponse(c, response)

	return
}
