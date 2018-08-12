package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/5112100070/Trek/src/conf"
	"github.com/5112100070/Trek/src/global"
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
	username := c.PostForm("username")
	password := c.PostForm("secret")

	baseUrl := conf.GConfig.BaseUrlConfig.ProductDNS
	path := "/make-login"
	data := url.Values{}
	data.Set("username", username)
	data.Set("secret", password)

	u, _ := url.ParseRequestURI(baseUrl)
	u.Path = path
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

	var resultResp struct {
		ServerMessage string                 `json:"server_message"`
		Data          map[string]interface{} `json:"data,omitempty"`
	}

	errUnMarshal := json.Unmarshal(body, &resultResp)
	if errUnMarshal != nil {
		global.Error.Println(errUnMarshal)
		global.ForbiddenResponse(c, nil)
		return
	} else if !resultResp.Data["is_success"].(bool) {
		global.ForbiddenResponse(c, nil)
		return
	} else {
		cookie := http.Cookie{
			Name:  global.UserCookie[global.GetEnv()],
			Value: resultResp.Data["nekot"].(string),
		}
		http.SetCookie(c.Writer, &cookie)

		http.Redirect(c.Writer, c.Request, conf.GConfig.BaseUrlConfig.BaseDNS, http.StatusSeeOther)
		return
	}
}
