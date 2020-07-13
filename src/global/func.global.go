package global

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

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

func GetServicePublic() PublicService {
	return Services.Public
}

func GetServiceUser() UserService {
	return Services.User
}

func GetServiceOrder() OrderService {
	return Services.Order
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
			Domain:  GetDNSNameForCookie(),
		}
	} else {
		newCookie = http.Cookie{
			Name:    UserCookie[GetEnv()],
			Value:   cookie,
			Expires: time.Now().Add(EXPIRE_COOKIE),
			Domain:  GetDNSNameForCookie(),
		}
	}

	http.SetCookie(c.Writer, &newCookie)

	mapper["UserDetail"] = user
}

func GetDNSNameForCookie() string {
	return DomainCookie[GetEnv()]
}
