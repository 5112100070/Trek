package session

import (
	"time"

	redigo "github.com/5112100070/Trek/src/global/redis"
)

const (
	redis_key_cookie = "cookie:c_"
	redis_timeout    = time.Duration((6 * time.Hour))

	USER_TYPE_ADMIN_TREK = 0
	USER_TYPE_ADMIN      = 1
	USER_TYPE_COMMON     = 2
)

type UserSession struct {
	ID       int64  `json:"user_id"`
	FullName string `json:"fullname"`
	Type     int    `json:"type"`
}

func (us UserSession) GetFullname() string {
	return us.FullName
}

type sessionRepo struct {
	redis redigo.Redis
}
