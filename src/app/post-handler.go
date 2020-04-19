package app

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/5112100070/Trek/src/global"
	gSession "github.com/5112100070/Trek/src/global/session"
	"github.com/5112100070/Trek/src/utils"
)

func SendRequestItem(c *gin.Context) {
	productName := c.PostForm("product_name")
	typeDuration := c.PostForm("type_duration")
	startDate := c.PostForm("start_date")
	email := c.PostForm("email")
	projectAddress := c.PostForm("project_address")

	if strings.TrimRight(productName, "\n") == "" ||
		strings.TrimRight(typeDuration, "\n") == "" ||
		strings.TrimRight(startDate, "\n") == "" ||
		strings.TrimRight(email, "\n") == "" ||
		strings.TrimRight(projectAddress, "\n") == "" {
		global.BadRequestResponse(c, nil)
		return
	}

	productID, errParse := strconv.ParseInt(c.PostForm("product_id"), 10, 64)
	if errParse != nil || productID <= 0 {
		global.Error.Println(errParse)
		global.BadRequestResponse(c, nil)
		return
	}

	duration, errParse := strconv.ParseInt(c.PostForm("duration"), 10, 64)
	if errParse != nil || duration <= 0 {
		global.Error.Println(errParse)
		global.BadRequestResponse(c, nil)
		return
	}

	total, errParse := strconv.ParseInt(c.PostForm("total"), 10, 64)
	if errParse != nil || total <= 0 {
		global.Error.Println(errParse)
		global.BadRequestResponse(c, nil)
		return
	}

	mail := utils.NewMail()
	mail.BuildReciever("g@trek.id", "wahyu.kukuh.herlambang@gmail.com")
	mail.BuildMessage(utils.BuildMessageForRequest(productID, productName, typeDuration, duration, total, startDate, email, projectAddress))
	mail.SendMail()

	global.OKResponse(c, nil)
}

func ProcessMakeLogin(c *gin.Context) {
	// make sure not too fast when give response
	defer func() {
		time.Sleep(4 * time.Second)
	}()

	email := c.PostForm("email")
	password := c.PostForm("password")

	service := global.GetServiceSession()
	loginResp, errResp := service.RequestLogin(email, password)
	if errResp != nil {
		log.Println(errResp)
		global.InternalServerErrorResponse(c, nil)
		return
	}

	//  return error cause from service server
	if loginResp.Error != nil {
		global.ErrorResponse(c, loginResp.Error)
		return
	}

	// return response from service server and keep to cookie
	if loginResp.Data != nil {
		cookie := gSession.SetSessionCookie(loginResp.Data.Token)

		http.SetCookie(c.Writer, &cookie)
		return
	}

	// return internal server error if not caught response from service server
	global.InternalServerErrorResponse(c, nil)
}

func ProcessMakeLogout(c *gin.Context) {
	token, errGetCookie := c.Cookie(global.UserCookie[global.GetEnv()])
	if errGetCookie != nil {
		global.Error.Println(errGetCookie)
		global.OKResponse(c, nil)
		return
	}

	service := global.GetServiceSession()
	errResp, err := service.RequestLogout(token)
	if err != nil {
		// internal server error from go.cgx.co.id
		global.Error.Println(err)
		global.InternalServerErrorResponse(c, nil)
	}

	if errResp != nil {
		// error response from go.cgx.co.id
		global.ErrorResponse(c, errResp)
		return
	}

	newCookie := gSession.SetExpireSessionCookie()

	http.SetCookie(c.Writer, &newCookie)
}
