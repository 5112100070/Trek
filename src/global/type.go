package global

import (
	"github.com/5112100070/Trek/src/app/session"
	"github.com/5112100070/Trek/src/entity"
	redigo "github.com/5112100070/Trek/src/global/redis"
	"github.com/5112100070/publib/storage/database"
)

// DBBundle : Data type of DB
type DBBundle struct {
	RedisSession redigo.Redis
	DB           map[string]database.Database
}

type RepoBundle struct {
	Session SessionService
	Public  PublicService
}

type SessionService interface {
	GetUser(string) (session.AccountResponse, error)
	SetUser(string, session.Account) error
	DelUser(string) error
}

type PublicService interface {
	SaveSubscriber(user entity.UserSubscriber) error
}
