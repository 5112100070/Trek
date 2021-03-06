package global

import (
	"github.com/5112100070/Trek/src/app/module"
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
	Module  ModuleService

	// map order status
	// get from response CGX server.
	// must initiable. if fails must return FATAL
	MapOrderStatus map[int]string
	// map pickup status
	// get from response CGX server.
	// must initiable. if fails must return FATAL
	MapPickupStatus map[int]string
}

type SessionService interface {
	CheckFeature(sessionID, pathURL, method string) (session.FeatureCheckResponse, error)
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
	GetDetailCompany(sessionID string, param user.DetailCompanyParam) (user.MainDetailCompanyResponse, error)
	GetListUsers(sessionID string, param user.ListUserParam) (user.MainListAccountResponse, error)
	GetListCompany(sessionID string, param user.ListCompanyParam) (user.MainListCompanyResponse, error)
	CreateUser(sessionID string, param user.CreateAccountParam) (*user.Error, error)
	UpdateUser(sessionID string, param user.UpdateAccountParam) (*user.Error, error)
	CreateCompany(sessionID string, param user.CreateCompanyParam) (*user.Error, error)
	UpdateCompany(sessionID string, param user.UpdateCompanyParam) (*user.Error, error)
	ChangePassword(sessionID string, newPassword string) (*user.Error, error)
	ChangePasswordAdmin(sessionID string, param user.ChangePasswordParam) (*user.Error, error)
	ChangeStatusAccount(sessionID string, param user.ChangeStatusAccParam) (*user.Error, error)
}

type OrderService interface {
	GetListOrderStatusCGX() (map[string]string, error)
	GetListPickupStatusCGX() (map[string]string, error)
	CreateOrderForAdmin(sessionID string, payload order.CreateOrderParam) (*order.CreateOrderForAdminResponse, error)
	GetOrderDetailForAdmin(sessionID string, orderID int64) (order.OrderReponse, *order.ErrorOrder, error)
	GetListOrders(sessionID string, param order.ListOrderParam) (order.MainListOrderResponse, error)
	GetListUnitInOrder(sessionID string) (order.MainListUnitResponse, error)
}

type ModuleService interface {
	GetListModule(sessionID string, param module.ListModuleParam) (module.ModuleResponse, error)
	GetListFeature(sessionID string, param module.ListFeatureParam) (module.FeatureResponse, error)
}
