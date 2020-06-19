package global

import (
	"github.com/5112100070/Trek/src/app/order"
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
	Order   OrderService
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
	GetDetailAccount(sessionID string, accountID int64) (user.MainDetailAccountResponse, error)
	GetDetailCompany(sessionID string, companyID int64) (user.MainDetailCompanyResponse, error)
	GetListUsers(sessionID string, param user.ListUserParam) (user.MainListAccountResponse, error)
	GetListCompany(sessionID string, param user.ListCompanyParam) (user.MainListCompanyResponse, error)
	CreateUser(sessionID string, param user.CreateAccountParam) (*user.Error, error)
	UpdateUser(sessionID string, param user.UpdateAccountParam) (*user.Error, error)
	CreateCompany(sessionID string, param user.CreateCompanyParam) (*user.Error, error)
	UpdateCompany(sessionID string, param user.UpdateCompanyParam) (*user.Error, error)
	ChangePassword(sessionID string, param user.ChangePasswordParam) (*user.Error, error)
	ChangeStatusAccount(sessionID string, param user.ChangeStatusAccParam) (*user.Error, error)
}

type OrderService interface {
	CreateOrderForAdmin(sessionID string, payload order.CreateOrderParam) (*order.CreateOrderForAdminResponse, error)
	GetListOrders(sessionID string, param order.ListOrderParam) (order.MainListOrderResponse, error)
}
