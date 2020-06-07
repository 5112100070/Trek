package dashboard

import (
	"log"
	"strconv"

	"github.com/5112100070/Trek/src/app/user"
	"github.com/5112100070/Trek/src/global"
	"github.com/gin-gonic/gin"
)

// UpdateAccount is endpoint hadler to update common user data
func UpdateAccount(c *gin.Context) {
	id, errparse := strconv.ParseInt(c.PostForm("id"), 10, 64)
	if errparse != nil {
		global.BadRequestResponse(c, nil)
		return
	}

	fullname := c.PostForm("fullname")
	phone := c.PostForm("phone")
	role, _ := strconv.Atoi(c.PostForm("role"))

	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println(errGetResponse)
		global.InternalServerErrorResponse(c, nil)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		global.InternalServerErrorResponse(c, nil)
		return
	}

	// define service user
	userService := global.GetServiceUser()
	param := user.UpdateAccountParam{
		ID:       id,
		Fullname: fullname,
		Phone:    phone,
		Role:     role,
	}
	resp, err := userService.UpdateUser(token, param)
	if err != nil {
		log.Println("cannot update user. Err", err)
		global.InternalServerErrorResponse(c, err.Error())
		return
	}

	response := map[string]interface{}{
		"error": resp,
	}

	global.OKResponse(c, response)
}

// UpdateCompany is endpoint hadler to update common user data
func UpdateCompany(c *gin.Context) {
	id, errparse := strconv.ParseInt(c.PostForm("id"), 10, 64)
	if errparse != nil {
		global.BadRequestResponse(c, nil)
		return
	}

	address := c.PostForm("company_address")
	companyName := c.PostForm("company_name")
	phone := c.PostForm("phone")
	role, _ := strconv.Atoi(c.PostForm("role"))

	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println(errGetResponse)
		global.InternalServerErrorResponse(c, nil)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		global.InternalServerErrorResponse(c, nil)
		return
	}

	// define service user
	userService := global.GetServiceUser()
	param := user.UpdateCompanyParam{
		ID:      id,
		Name:    companyName,
		Address: address,
		Phone:   phone,
		Role:    role,
	}

	resp, err := userService.UpdateCompany(token, param)
	if err != nil {
		log.Println("cannot update company. Err", err)
		global.InternalServerErrorResponse(c, err.Error())
		return
	}

	response := map[string]interface{}{
		"error": resp,
	}

	global.OKResponse(c, response)
}
