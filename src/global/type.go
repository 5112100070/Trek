package global

import (
	"github.com/5112100070/Trek/src/app/session"
	redigo "github.com/5112100070/Trek/src/global/redis"
)

// DBBundle : Data type of DB
type DBBundle struct {
	RedisSession redigo.Redis
}

type RepoBundle struct {
	Session SessionService
}

type SessionService interface {
	GetUser(string) (session.UserSession, error)
	SetUser(string, session.UserSession) error
}
