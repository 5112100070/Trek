package global

import (
	"io"
	"log"

	"github.com/5112100070/Trek/src/app/session"
)

func InitLogError(errorHandle io.Writer) {
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func InitRepoBundle(dbBundle DBBundle) {
	Services = RepoBundle{
		Session: session.InitSessionRepo(dbBundle.RedisSession),
	}
}

func InitDNSName(DnsName string) {
	dns = DnsName
}
