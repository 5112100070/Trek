package global

import (
	"os"

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
