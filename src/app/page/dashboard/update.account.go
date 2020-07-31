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
		global.BadRequestResponse(c, "invalid account id")
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
		global.BadRequestResponse(c, "Invalid Company ID")
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

// AdminChangePassword is endpoint to changes all user password
func AdminChangePassword(c *gin.Context) {
	userID, errparse := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
	if errparse != nil {
		global.BadRequestResponse(c, "invalid account")
		return
	}

	newPassword := c.PostForm("new_password")

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
	param := user.ChangePasswordParam{
		UserID:      userID,
		NewPassword: newPassword,
	}

	resp, err := userService.ChangePasswordAdmin(token, param)
	if err != nil {
		log.Println("cannot change password account. Err", err)
		global.InternalServerErrorResponse(c, err.Error())
		return
	}

	response := map[string]interface{}{
		"error": resp,
	}

	global.OKResponse(c, response)
}

func ChangePassword(c *gin.Context) {
	newPassword := c.PostForm("new_password")

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
	resp, err := userService.ChangePassword(token, newPassword)
	if err != nil {
		log.Println("cannot change password account. Err", err)
		global.InternalServerErrorResponse(c, err.Error())
		return
	}

	response := map[string]interface{}{
		"error": resp,
	}

	global.OKResponse(c, response)
}

// AdminChangeActivation is endpoint to change status activation account
func AdminChangeActivation(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
	companyID, _ := strconv.ParseInt(c.PostForm("company_id"), 10, 64)

	var isEnabled bool
	payloadStatus := c.PostForm("is_enabled")
	if payloadStatus == "true" {
		isEnabled = true
	} else if payloadStatus == "false" {
		isEnabled = false
	} else {
		log.Printf("func AdminChangeActivation error when parse status. payload: %v \n", payloadStatus)
		global.BadRequestResponse(c, "invalid is_enable value")
		return
	}

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
	param := user.ChangeStatusAccParam{
		UserID:    userID,
		CompanyID: companyID,
		IsEnabled: isEnabled,
	}

	resp, err := userService.ChangeStatusAccount(token, param)
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
