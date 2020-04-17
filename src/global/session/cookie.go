package session

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func SetSessionCookie(token string) http.Cookie {
	location, _ := time.LoadLocation("Asia/Jakarta")
	wib := time.Now().In(location)

	return http.Cookie{
		Name:    UserCookie[getEnv()],
		Value:   token,
		Domain:  DomainCookie[getEnv()],
		Expires: wib.Add(EXPIRE_COOKIE),
	}
}

func SetExpireSessionCookieGin(c *gin.Context) {
	// c.SetCookie(UserCookie[getEnv()], "", -1, "/", DomainCookie[getEnv()], http.SameSiteDefaultMode, true, true)
}

func SetExpireSessionCookie() http.Cookie {
	return http.Cookie{
		Name:   UserCookie[getEnv()],
		Domain: DomainCookie[getEnv()],
		MaxAge: -1,
	}
}

func getEnv() string {
	cfgenv := os.Getenv("CGXENV")
	if cfgenv == "" {
		cfgenv = "development"
	}
	return cfgenv
}
