package conf

import (
	"log"

	"github.com/5112100070/Trek/src/global"
	redigo "github.com/5112100070/Trek/src/global/redis"
	"github.com/ruizu/gcfg"
)

func ReadConfig(filePath string) (Config, bool) {
	var c Config
	if err := gcfg.ReadFileInto(&c, filePath); err != nil {
		log.Printf("%s\n", err)
		return Config{}, false
	}
	return c, true
}

func InitRedis(c Config) global.DBBundle {
	redisSession := redigo.New(c.Session.Redis, &c.RedigoDefault)

	DB := global.DBBundle{
		RedisSession: redisSession,
	}

	_, err := redisSession.PING()
	if err != nil {
		log.Fatal("Redis Session data not accessible, please check config")
	}

	return DB
}
