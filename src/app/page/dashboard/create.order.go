package dashboard

import (
	"net/http"

	"github.com/5112100070/Trek/src/app/order"
	"github.com/5112100070/Trek/src/conf"
	"github.com/5112100070/Trek/src/global"
	"github.com/gin-gonic/gin"
)

func CreateOrderPageHandler(c *gin.Context) {
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

	renderData := gin.H{
		"UserDetail": accountResp.Data,
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
