package dashboard

import (
	"fmt"
	"log"
	"strconv"

	"github.com/5112100070/Trek/src/conf"

	"github.com/gin-gonic/gin"

	"github.com/5112100070/Trek/src/app/user"
	"github.com/5112100070/Trek/src/global"
)

// CreateNewAccount is endpoint hadler to create new user account
func CreateNewAccount(c *gin.Context) {
	fullname := c.PostForm("fullname")
	if fullname == "" {
		global.BadRequestResponse(c, "Masukkan nama lengkap")
		return
	}

	email := c.PostForm("email")
	if email == "" {
		global.BadRequestResponse(c, "Masukkan alamat email")
		return
	}

	phone := c.PostForm("phone")
	if phone == "" {
		global.BadRequestResponse(c, "Masukkan nomor telepon yang dapat dihubungi")
		return
	}

	role, _ := strconv.Atoi(c.PostForm("role"))
	if role == 0 {
		global.BadRequestResponse(c, "Invalid role")
		return
	}

	companyID, _ := strconv.ParseInt(c.PostForm("company_id"), 10, 64)
	if companyID == 0 {
		global.BadRequestResponse(c, "Perusahaan tidak valid")
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		global.Error.Println("func CreateNewAccount error when get multipart form: ", err)
		global.BadRequestResponse(c, "Invalid request payload")
		return
	}

	// Check user session
	accountResp, token, errGetResponse := getUserProfile(c)
	if errGetResponse != nil {
		global.Error.Println("func CreateNewAccount error when getUserProfile", errGetResponse)
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
			global.Error.Println("func CreateNewAccount error when create profile image ", errOpenFile)
			global.InternalServerErrorResponse(c, errOpenFile)
			return
		}
		defer profileImage.Close()

		var errUpload error
		filename, errUpload = global.UploadProfileImage(profileImage, email)
		if errUpload != nil {
			global.Error.Println("func CreateNewAccount error when upload profile image ", errUpload)
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
	param := user.CreateAccountParam{
		Fullname:     fullname,
		Phone:        phone,
		Email:        email,
		CompanyID:    companyID,
		Role:         role,
		ProfileImage: filename,
	}

	resp, err := userService.CreateUser(token, param)
	if err != nil {
		log.Println("func CreateNewAccount cannot save new account. Err", err)
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
	if companyName == "" {
		global.BadRequestResponse(c, "Masukkan nama perusahaan")
		return
	}

	address := c.PostForm("address")
	if address == "" {
		global.BadRequestResponse(c, "Masukkan alamat perusahaan")
		return
	}

	phone := c.PostForm("phone")
	if phone == "" {
		global.BadRequestResponse(c, "Masukkan nomor telepon perusahaan")
		return
	}

	role, _ := strconv.Atoi(c.PostForm("role"))
	if role == 0 {
		global.BadRequestResponse(c, "Invalid role")
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
		global.InternalServerErrorResponse(c, errGetResponse)
		return
	}

	// expire - we remove cookie and redirect it
	if accountResp.Error != nil {
		global.BadRequestResponse(c, "Session telah habis. silahkan reload halaman ini")
		return
	}

	// create/save image on location
	files := form.File["company_image"]

	var filename string
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
	} else {
		// if no files detected. return error
		global.BadRequestResponse(c, "Please upload the company image")
		return
	}

	// define service user
	userService := global.GetServiceUser()
	param := user.CreateCompanyParam{
		Name:      companyName,
		Phone:     phone,
		Address:   address,
		Role:      role,
		ImageLogo: filename,
	}

	// send request to cgx to create company data
	resp, err := userService.CreateCompany(token, param)
	if err != nil {
		global.Error.Println("func CreateNewCompany error when request create company ", err)
		global.InternalServerErrorResponse(c, err.Error())
		return
	}

	response := map[string]interface{}{
		"error": resp,
	}

	global.OKResponse(c, response)
}
