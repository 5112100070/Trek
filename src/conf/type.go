package conf

import (
	redigo "github.com/5112100070/Trek/src/global/redis"
)

type Config struct {
	BaseUrlConfig BaseUrl
	RedigoDefault redigo.RedisConfig
	Session       SessionConfig
}

type BaseUrl struct {
	BaseDNS    string
	ProductDNS string
}

type SessionConfig struct {
	Redis string
}
