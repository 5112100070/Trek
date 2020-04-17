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

type AccountResponse struct {
	Data              *Account `json:"data", omitempty`
	Error             *Error   `json:"error", omitempty`
	ServerProcessTime string   `json:"server_process_time"`
}

type Account struct {
	ID         int64          `json:"user_id"`
	Fullname   string         `json:"fullname"`
	Email      string         `json:"email"`
	CreateTime string         `json:"create_time"`
	Role       int            `json:"role"`
	Company    CompanyProfile `json:"company"`
}

type CompanyProfile struct {
	ID          int64  `json:"company_id"`
	CompanyName string `json:"company_name"`
	SecretKey   string `json:"secret_key"`
	IsEnabled   bool   `json:"is_enabled"`
	Role        int    `json:"role"`
}

type Error struct {
	Code    int    `json:"code"`
	Massage string `json:"massage"`
}

type sessionRepo struct {
	redis redigo.Redis
}
