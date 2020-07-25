package dashboard

import (
	"log"
	"net/http"

	"github.com/5112100070/Trek/src/app/user"

	"github.com/5112100070/Trek/src/app/order"
	"github.com/5112100070/Trek/src/conf"
	"github.com/5112100070/Trek/src/global"
	"github.com/gin-gonic/gin"
)

func CreateOrderPageHandler(c *gin.Context) {
	// Check user session
	accountResp, sessionID, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println(errGetResponse)
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
		return
	}

	if unitsResp.Error != nil {
		log.Println("func CreateOrderPageHandler error expected when call get list unit from CGX: ", unitsResp.Error.Detail)
	}

	companiesResp, errGetCompany := global.GetServiceUser().GetListCompany(sessionID, user.ListCompanyParam{
		Page:             1,
		Rows:             50,
		OrderType:        "desc",
		FilterByIsEnable: "1",
	})
	if errGetCompany != nil {
		// add internal server error page response
		log.Println("func CreateOrderPageHandler error when call get list unit to CGX: ", errGetCompany)
		return
	}

	if companiesResp.Error != nil {
		log.Println("func CreateOrderPageHandler error expected when call get list unit to CGX: ", companiesResp.Error.Detail)
	}

	renderData := gin.H{
		"UserDetail": accountResp.Data,
		"Units":      unitsResp.Data.Units,
		"Companies":  companiesResp.Data.Companies,
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
		global.ForbiddenResponse(c, nil)
		return
	}

	errBindJSON := c.BindJSON(&body)
	if errBindJSON != nil {
		global.Error.Printf("func CreateOrderForAdmin error when marshal json: %v\n", errBindJSON)
		global.BadRequestResponse(c, nil)
		return
	}

	resp, err := global.GetServiceOrder().CreateOrderForAdmin(token, body)
	if err != nil {
		global.Error.Println("func CreateOrderForAdmin error when create order for admin: ", err)
		global.InternalServerErrorResponse(c, nil)
		return
	}

	response := map[string]interface{}{
		"response": resp,
	}

	global.OKResponse(c, response)

	return
}
