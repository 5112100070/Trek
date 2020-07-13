package dashboard

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/5112100070/Trek/src/app/user"
	"github.com/5112100070/Trek/src/global"
)

// CreateNewAccount is endpoint hadler to create new user account
func CreateNewAccount(c *gin.Context) {
	fullname := c.PostForm("fullname")
	email := c.PostForm("email")
	phone := c.PostForm("phone")

	role, _ := strconv.Atoi(c.PostForm("role"))
	companyID, _ := strconv.ParseInt(c.PostForm("company_id"), 10, 64)

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
	param := user.CreateAccountParam{
		Fullname:  fullname,
		Phone:     phone,
		Email:     email,
		CompanyID: companyID,
		Role:      role,
	}
	resp, err := userService.CreateUser(token, param)
	if err != nil {
		log.Println("cannot save subscriber. Err", err)
		global.InternalServerErrorResponse(c, err.Error())
		return
	}

	response := map[string]interface{}{
		"error": resp,
	}

	global.OKResponse(c, response)
}

// CreateNewCompany is endpoint handler to create new user account
func CreateNewCompany(c *gin.Context) {
	companyName := c.PostForm("name")
	address := c.PostForm("address")
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
	param := user.CreateCompanyParam{
		Name:    companyName,
		Phone:   phone,
		Address: address,
		Role:    role,
	}
	resp, err := userService.CreateCompany(token, param)
	if err != nil {
		log.Println("cannot save subscriber. Err", err)
		global.InternalServerErrorResponse(c, err.Error())
		return
	}

	response := map[string]interface{}{
		"error": resp,
	}

	global.OKResponse(c, response)
}
