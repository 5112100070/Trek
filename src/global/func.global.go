package global

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func SetCookie(c *gin.Context, key string, value string, base_url string) {
	c.SetCookie(key, value, (5 * 3600), "/", base_url, true, true)
}

func GetEnv() string {
	cfgenv := os.Getenv("TREKENV")
	if cfgenv == "" {
		cfgenv = "development"
	}
	return cfgenv
}

func GetServiceSession() SessionService {
	return Services.Session
}

func GetDefaultUserAttribute(c *gin.Context, mapper map[string]interface{}) {
	service := GetServiceSession()
	cookie, errGetCookie := c.Cookie(UserCookie[GetEnv()])
	if errGetCookie != nil {
		Error.Println(errGetCookie)
		return
	}

	var newCookie http.Cookie
	user, errLogin := service.GetUser(cookie)
	if errLogin != nil {
		newCookie = http.Cookie{
			Name:    UserCookie[GetEnv()],
			Expires: time.Now().Add(time.Duration(0)),
		}
	} else {
		newCookie = http.Cookie{
			Name:    UserCookie[GetEnv()],
			Value:   cookie,
			Expires: time.Now().Add(expire_cookie),
		}
	}

	http.SetCookie(c.Writer, &newCookie)

	mapper["UserDetail"] = user
}
