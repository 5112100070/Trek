package session

import (
	"encoding/json"
	"fmt"

	redigo "github.com/5112100070/Trek/src/global/redis"
)

func InitSessionRepo(sessionRedis redigo.Redis) *sessionRepo {
	return &sessionRepo{
		redis: sessionRedis,
	}
}

func (repo sessionRepo) GetUser(cookie string) (UserSession, error) {
	var user UserSession
	key := fmt.Sprintf("%v%v", redis_key_cookie, cookie)

	result, err := repo.redis.GET(key)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal([]byte(result), &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo sessionRepo) SetUser(cookie string, user UserSession) error {
	key := fmt.Sprintf("%v%v", redis_key_cookie, cookie)

	m, _ := json.Marshal(user)

	err := repo.redis.SETEX(key, redis_timeout, string(m))
	if err != nil {
		return err
	}

	return nil
}
