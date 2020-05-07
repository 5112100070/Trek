package global

import (
	"github.com/5112100070/Trek/src/app/session"
	"github.com/5112100070/Trek/src/app/user"
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
	User    UserService
}

type SessionService interface {
	GetUser(string) (session.AccountResponse, error)
	RequestLogin(string, string) (session.LoginResponse, error)
	RequestLogout(string) (*session.Error, error)
	SetUser(string, session.Account) error
	DelUser(string) error
}

type PublicService interface {
	SaveSubscriber(user entity.UserSubscriber) error
}

type UserService interface {
	GetListUsers(sessionID string, param user.ListUserParam) (user.MainListAccountResponse, error)
	GetListCompany(sessionID string, param user.ListCompanyParam) (user.MainListCompanyResponse, error)
}
