package global

import (
	"log"
	"time"
)

const (
	expire_cookie = time.Duration((6 * time.Hour))
)

//Error Logger
var Error *log.Logger

//All Repository Access
var Services RepoBundle

var dns string

var UserCookie = map[string]string{
	"production":  "_TREK_",
	"staging":     "_TREK_TEST_",
	"development": "_TREK_DEV_",
}
