package global

import (
	"io"
	"log"

	"github.com/5112100070/Trek/src/app/order"
	"github.com/5112100070/Trek/src/app/public"
	"github.com/5112100070/Trek/src/app/session"
	"github.com/5112100070/Trek/src/app/user"
)

func InitLogError(errorHandle io.Writer) {
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func InitRepoBundle(dbBundle DBBundle) {
	Services = RepoBundle{
		Session: session.InitSessionRepo(dbBundle.RedisSession),
		Public:  public.InitPublicRepo(dbBundle.DB),
		User:    user.InitUserRepo(),
		Order:   order.InitOrderRepo(),
	}
}
