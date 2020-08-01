package dashboard

import (
	"fmt"
	"log"
	"strconv"

	"github.com/5112100070/Trek/src/app/user"
	"github.com/5112100070/Trek/src/conf"
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
	if fullname == "" {
		global.BadRequestResponse(c, "Tidak boleh menghapus nama profile")
		return
	}

	phone := c.PostForm("phone")
	if phone == "" {
		global.BadRequestResponse(c, "Tidak boleh menghapus nomor telepon pengguna")
		return
	}

	role, _ := strconv.Atoi(c.PostForm("role"))
	if role == 0 {
		global.BadRequestResponse(c, "invalid account role")
		return
	}

	email := c.PostForm("email")
	if fullname == "" {
		global.BadRequestResponse(c, "Invalid request payload")
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		global.Error.Println("func UpdateAccount error when get multipart form: ", err)
		global.BadRequestResponse(c, "Invalid request payload")
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

	// create/save image on location
	files := form.File["profile_image"]

	var filename string
	if len(files) > 0 {
		profileImage, errOpenFile := files[0].Open()
		if errOpenFile != nil {
			global.Error.Println("func UpdateAccount error when create profile image ", errOpenFile)
			global.InternalServerErrorResponse(c, errOpenFile)
			return
		}
		defer profileImage.Close()

		var errUpload error
		filename, errUpload = global.UploadProfileImage(profileImage, email)
		if errUpload != nil {
			global.Error.Println("func UpdateAccount error when upload profile image ", errUpload)
			global.InternalServerErrorResponse(c, errUpload)
			return
		}

		filename = fmt.Sprintf("%s/img/user/%s", conf.GConfig.BaseUrlConfig.BaseDNS, filename)
	} else {
		// if no files detected. using default image
		filename = fmt.Sprintf("%s/img/user/default-account.png", conf.GConfig.BaseUrlConfig.BaseDNS)
	}

	// define service user
	userService := global.GetServiceUser()
	param := user.UpdateAccountParam{
		ID:           id,
		Fullname:     fullname,
		Phone:        phone,
		Role:         role,
		ProfileImage: filename,
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

	companyName := c.PostForm("company_name")
	if companyName == "" {
		global.BadRequestResponse(c, "Tidak boleh menghapus nama perusahaan")
		return
	}

	address := c.PostForm("company_address")
	if address == "" {
		global.BadRequestResponse(c, "Tidak boleh menghapus alamat perusahaan")
		return
	}

	phone := c.PostForm("phone")
	if phone == "" {
		global.BadRequestResponse(c, "Tidak boleh menghapus no telepon perusahaan")
		return
	}

	role, _ := strconv.Atoi(c.PostForm("role"))
	if role == 0 {
		global.BadRequestResponse(c, "Tidak boleh menghapus role")
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		global.Error.Println("func CreateNewCompany error when get multipart form: ", err)
		global.BadRequestResponse(c, "Silahkan Upload logo perusahaan")
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

	// create/save image on location
	files := form.File["company_image"]

	var filename string
	// if files detected. filename will filled
	// if no files detected. no need to update image
	if len(files) > 0 {
		companyImage, errOpenFile := files[0].Open()
		if errOpenFile != nil {
			global.Error.Println("func CreateNewCompany error when create company image ", errOpenFile)
			global.InternalServerErrorResponse(c, errOpenFile)
			return
		}
		defer companyImage.Close()

		var errUpload error
		filename, errUpload = global.UploadCompanyImage(companyImage, companyName)
		if errUpload != nil {
			global.Error.Println("func CreateNewCompany error when upload company image ", errUpload)
			global.InternalServerErrorResponse(c, errUpload)
			return
		}

		filename = fmt.Sprintf("%s/img/partner/%s", conf.GConfig.BaseUrlConfig.BaseDNS, filename)
	}

	// define service user
	userService := global.GetServiceUser()
	param := user.UpdateCompanyParam{
		ID:        id,
		Name:      companyName,
		Address:   address,
		Phone:     phone,
		Role:      role,
		ImageLogo: filename,
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
