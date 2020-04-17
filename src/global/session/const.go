package session

import "time"

const (
	EXPIRE_COOKIE        = time.Duration((2 * 24 * time.Hour))
	EXPIRE_COOKIE_IN_SEC = 2 * 24 * 60 * 60
)

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
