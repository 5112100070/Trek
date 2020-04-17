package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/5112100070/Trek/src/conf"
	urlConst "github.com/5112100070/Trek/src/constants/url"
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

	data := url.Values{}
	data.Set("email", email)
	data.Set("password", password)

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_REQUEST_LOGIN
	urlStr := u.String()

	client := &http.Client{}
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		global.Error.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, errGetResp := client.Do(req)
	if err != nil {
		global.Error.Println(errGetResp)
		global.InternalServerErrorResponse(c, nil)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.Error.Println(err)
		global.InternalServerErrorResponse(c, nil)
		return
	}

	/*
		Example response from go.cgx.co.id
		success without error
				{
			    "data": {
			        "message": "success generate token",
			        "token": "eyJleHAiOjE1ODY5MDY4MTksImlhdCI6MTU4Njg4ODgxOSwia"
			    },
			    "server_process_time": "253.438488ms"
			}

		success with error
			{
			    "error": {
			        "code": 4000,
			        "massage": "invalid email or password"
			    },
			    "server_process_time": "9.102529ms"
			}
	*/

	var resultResp struct {
		ServerProcessTime string                 `json:"server_process_time"`
		Data              map[string]interface{} `json:"data,omitempty"`
		Error             map[string]interface{} `json:"error,omitempty"`
	}

	errUnMarshal := json.Unmarshal(body, &resultResp)
	if errUnMarshal != nil {
		global.Error.Println(errUnMarshal)
		global.InternalServerErrorResponse(c, nil)
		return
	}

	//  return error cause from service server
	if resultResp.Error != nil {
		global.ErrorResponse(c, resultResp.Error)
		return
	}

	// return response from service server and keep to cookie
	if resultResp.Data != nil {
		cookie := gSession.SetSessionCookie(resultResp.Data["token"].(string))

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
	err := service.DelUser(token)
	if err != nil {
		global.Error.Println(err)
	}

	newCookie := http.Cookie{
		Name:    global.UserCookie[global.GetEnv()],
		Value:   token,
		Expires: time.Now().Add(time.Duration(0)),
	}

	http.SetCookie(c.Writer, &newCookie)
}
