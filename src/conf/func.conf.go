package conf

import (
	"log"

	redigo "github.com/5112100070/Trek/src/global/redis"
	"github.com/5112100070/publib/storage/database"
	"github.com/5112100070/publib/storage/database/sql_wrapper"
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

func InitRedis(c Config) redigo.Redis {
	redisSession := redigo.New(c.Session.Redis, &c.RedigoDefault)

	_, err := redisSession.PING()
	if err != nil {
		log.Println("Redis Session data not accessible, please check config")
	}

	return redisSession
}

func InitDatabase(cfg map[string]*DatabaseConfig) map[string]database.Database {
	res := map[string]database.Database{}

	for k, v := range cfg {
		dbModule := sql_wrapper.New(sql_wrapper.Config{
			Driver: v.Driver,
			Master: v.Master,
			Slave:  v.Slave,
		})

		res[k] = dbModule
	}

	return res
}
