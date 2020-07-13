package global

import (
	"log"
	"time"
)

const (
	EXPIRE_COOKIE = time.Duration((2 * 24 * time.Hour))
)

//Error Logger
var Error *log.Logger

//All Repository Access
var Services RepoBundle

var UserCookie = map[string]string{
	"production":  "_CGX_",
	"staging":     "_CGX_TEST_",
	"development": "_CGX_DEV_",
}

var DomainCookie = map[string]string{
	"production":  ".cgx.co.id",
	"staging":     ".",
	"development": ".trek.ndvl",
}
