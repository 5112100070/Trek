package app

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/5112100070/Trek/src/global"
	"github.com/5112100070/Trek/src/utils"
)

func SendRequestItem(c *gin.Context) {
	productName := c.PostForm("product_name")
	typeDuration := c.PostForm("type_duration")
	startDate := c.PostForm("start_date")
	email := c.PostForm("email")

	productID, errParse := strconv.ParseInt(c.PostForm("product_id"), 10, 64)
	if errParse != nil {
		global.Error.Println(errParse)
		global.BadRequestResponse(c, nil)
		return
	}

	duration, errParse := strconv.ParseInt(c.PostForm("duration"), 10, 64)
	if errParse != nil {
		global.Error.Println(errParse)
		global.BadRequestResponse(c, nil)
		return
	}

	total, errParse := strconv.ParseInt(c.PostForm("total"), 10, 64)
	if errParse != nil {
		global.Error.Println(errParse)
		global.BadRequestResponse(c, nil)
		return
	}

	mail := utils.NewMail()
	mail.BuildReciever("ghazi.binarandi@gmail.com", "hanifnarendra@gmail.com", "wahyu.kukuh.herlambang@gmail.com")
	mail.BuildMessage(utils.BuildMessageForRequest(productID, productName, typeDuration, duration, total, startDate, email))
	mail.SendMail()

	global.OKResponse(c, nil)
}
