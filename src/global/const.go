package global

import (
	"log"
)

//Error Logger
var Error *log.Logger

var UserCookie = map[string]string{
	"production":  "_TREK_",
	"staging":     "_TREK_TEST_",
	"development": "_TREK_DEV_",
}
