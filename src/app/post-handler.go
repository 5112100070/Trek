package app

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

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
