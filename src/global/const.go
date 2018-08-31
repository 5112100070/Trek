package global

import (
	"log"
	"time"
)

const (
	EXPIRE_COOKIE = time.Duration((6 * time.Hour))
)

//Error Logger
var Error *log.Logger

//All Repository Access
var Services RepoBundle

var UserCookie = map[string]string{
	"production":  "_TREK_",
	"staging":     "_TREK_TEST_",
	"development": "_TREK_DEV_",
}

var DomainCookie = map[string]string{
	"production":  ".trek.id",
	"staging":     ".",
	"development": ".trek.ndvl",
}
