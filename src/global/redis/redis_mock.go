package redis

import (
	"errors"
	"time"
)

type RedisMock struct {
	// for unit testing. Default is false (return nil on error), if true will return error
	MakeErrorResponse bool
	Value             string
}

func (r *RedisMock) PING() (string, error) {
	if r.MakeErrorResponse {
		return "", errors.New("Unit Testing Expected Error")
	}
	return "Pong", nil
}

func (r *RedisMock) TTL(key string) (time.Duration, error) {
	if r.MakeErrorResponse {
		return time.Duration(0), errors.New("Unit Testing Expected Error")
	}
	return time.Duration(14400), nil
}

func (r *RedisMock) EXISTS(key string) (bool, error) {
	if r.MakeErrorResponse {
		return false, errors.New("Unit Testing Expected Error")
	}
	return true, nil
}

func (r *RedisMock) EXPIRE(key string, duration time.Duration) error {
	if r.MakeErrorResponse {
		return errors.New("Unit Testing Expected Error")
	}
	return nil
}

func (r *RedisMock) INCR(key string) (int64, error) {
	if r.MakeErrorResponse {
		return 0, errors.New("Unit Testing Expected Error")
	}
	return 0, nil
}

func (r *RedisMock) INCRBY(key string, val int) error {
	if r.MakeErrorResponse {
		return errors.New("Unit Testing Expected Error")
	}
	return nil
}

func (r *RedisMock) SET(key string, value interface{}) error {
	if r.MakeErrorResponse {
		return errors.New("Unit Testing Expected Error")
	}
	return nil
}

func (r *RedisMock) SETEX(key string, dur time.Duration, value interface{}) error {
	if r.MakeErrorResponse {
		return errors.New("Unit Testing Expected Error")
	}
	return nil
}

func (r *RedisMock) GET(key string) (string, error) {
	if r.MakeErrorResponse {
		return "", errors.New("Unit Testing Expected Error")
	}
	return "", nil
}

func (r *RedisMock) MGET(key ...string) ([]string, error) {
	if r.MakeErrorResponse {
		return nil, errors.New("Unit Testing Expected Error")
	}
	return nil, nil
}

func (r *RedisMock) DEL(key string) error {
	if r.MakeErrorResponse {
		return errors.New("Unit Testing Expected Error")
	}
	return nil
}

func (r *RedisMock) HSET(key string, fields string, value string) error {
	if r.MakeErrorResponse {
		return errors.New("Unit Testing Expected Error")
	}
	return nil
}

func (r *RedisMock) HSETEX(key string, duration time.Duration, field string, value string) error {
	if r.MakeErrorResponse {
		return errors.New("Unit Testing Expected Error")
	}
	return nil
}

func (r *RedisMock) HMSET(key string, data map[string]interface{}) error {
	if r.MakeErrorResponse {
		return errors.New("Unit Testing Expected Error")
	}
	return nil
}

func (r *RedisMock) HGET(key string, field string) (string, error) {
	if r.MakeErrorResponse {
		return "", errors.New("Unit Testing Expected Error")
	}
	return "Mock Reply", nil
}

func (r *RedisMock) HMGET(key string, field ...string) (map[string]string, error) {
	if r.MakeErrorResponse {
		return nil, errors.New("Unit Testing Expected Error")
	}
	return nil, nil
}

func (r *RedisMock) HGETALL(key string) (map[string]string, error) {
	if r.MakeErrorResponse {
		return nil, errors.New("Unit Testing Expected Error")
	}
	return nil, nil
}

func (r *RedisMock) HDEL(key string, fields ...string) error {
	if r.MakeErrorResponse {
		return errors.New("Unit Testing Expected Error")
	}
	return nil
}

func (r *RedisMock) HEXISTS(key string, field string) (bool, error) {
	if r.MakeErrorResponse {
		return false, errors.New("Unit Testing Expected Error")
	}
	return true, nil
}

func (r *RedisMock) ZADD(key string, fields ...Z) error {
	if r.MakeErrorResponse {
		return errors.New("Unit Testing Expected Error")
	}
	return nil
}

func (r *RedisMock) ZSCORE(key string, value interface{}) (int, error) {
	if r.MakeErrorResponse {
		return 0, errors.New("Unit Testing Expected Error")
	}
	return 0, nil
}

func (r *RedisMock) ZREM(key string, fields ...string) error {
	if r.MakeErrorResponse {
		return errors.New("Unit Testing Expected Error")
	}
	return nil
}

func (r *RedisMock) ZRANGE(key string, start int, end int) ([]string, error) {
	if r.MakeErrorResponse {
		return nil, errors.New("Unit Testing Expected Error")
	}
	return []string{
		"{\"inbox_id\":13062,\"talk_id\":5410,\"shop_id\":2182,\"product_id\":26388035,\"user_id\":34293,\"status\":1,\"message\":\"\\u003ca rel=\\\"nofollow\\\" target=\\\"_blank\\\" href=\\\"https://tkp.me/r?url=http://new.wh-wahyu.ndvl/roti/roti-apel-roti-apel/talk\\\"\\u003ehttp://new\\u0026bull;wh-wahyu\\u0026bull;ndvl/roti/roti-apel-roti-ap...\\u003c/a\\u003e\",\"total_comment\":0,\"create_time_db\":\"2018-03-27T17:50:47Z\",\"shop_admin\":null,\"fraud_status\":0}",
	}, nil
}

func (r *RedisMock) ZRANGEBYSCORE(key string, opt ZRangeByScore) ([]string, error) {
	if r.MakeErrorResponse {
		return nil, errors.New("Unit Testing Expected Error")
	}
	return []string{
		"{\"inbox_id\":13062,\"talk_id\":5410,\"shop_id\":2182,\"product_id\":26388035,\"user_id\":34293,\"status\":1,\"message\":\"\\u003ca rel=\\\"nofollow\\\" target=\\\"_blank\\\" href=\\\"https://tkp.me/r?url=http://new.wh-wahyu.ndvl/roti/roti-apel-roti-apel/talk\\\"\\u003ehttp://new\\u0026bull;wh-wahyu\\u0026bull;ndvl/roti/roti-apel-roti-ap...\\u003c/a\\u003e\",\"total_comment\":0,\"create_time_db\":\"2018-03-27T17:50:47Z\",\"shop_admin\":null,\"fraud_status\":0}",
	}, nil
}

func (r *RedisMock) HGETP(keys []string, field string) (map[string]interface{}, error) {
	if r.MakeErrorResponse {
		return nil, errors.New("Unit Testing Expected Error")
	}
	return nil, nil
}

func (r *RedisMock) HMGETP(key []string, field ...string) (map[string]interface{}, error) {
	if r.MakeErrorResponse {
		return nil, errors.New("Unit Testing Expected Error")
	}
	return nil, nil
}

func (r *RedisMock) LTRIM(key string, start interface{}, end interface{}) error {
	if r.MakeErrorResponse {
		return errors.New("Unit Testing Expected Error")
	}
	return nil
}

func (r *RedisMock) LPOP(key string) (string, error) {
	if r.MakeErrorResponse {
		return "", errors.New("Unit Testing Expected Error")
	}
	return r.Value, nil
}

func (r *RedisMock) LPUSH(key string, value interface{}) error {
	if r.MakeErrorResponse {
		return errors.New("Unit Testing Expected Error")
	}
	return nil
}

func (r *RedisMock) LLEN(key string) (int64, error) {
	if r.MakeErrorResponse {
		return 0, errors.New("Unit Testing Expected Error")
	}
	return 0, nil
}

func (r *RedisMock) LPUSHEX(key string, duration time.Duration, value interface{}) error {
	if r.MakeErrorResponse {
		return errors.New("Unit Testing Expected Error")
	}
	return nil
}

func (r *RedisMock) LREM(key string, count interface{}, value interface{}) (int64, error) {
	if r.MakeErrorResponse {
		return 0, errors.New("Unit Testing Expected Error")
	}
	return 0, nil
}

func (r *RedisMock) LRANGE(key string, start interface{}, stop interface{}) ([]string, error) {
	if r.MakeErrorResponse {
		return nil, errors.New("Unit Testing Expected Error")
	}
	return nil, nil
}

func (r *RedisMock) LRANGEP(keys []string, fields ...string) (map[string][]int, error) {
	if r.MakeErrorResponse {
		return nil, errors.New("Unit Testing Expected Error")
	}
	return nil, nil
}

func (r *RedisMock) SADD(key string, val ...interface{}) error {
	if r.MakeErrorResponse {
		return errors.New("Unit Testing Expected Error")
	}
	return nil
}

func (r *RedisMock) SREM(key string, val ...interface{}) error {
	if r.MakeErrorResponse {
		return errors.New("Unit Testing Expected Error")
	}
	return nil
}

func (r *RedisMock) SMEMBERS(key string) ([]string, error) {
	if r.MakeErrorResponse {
		return nil, errors.New("Unit Testing Expected Error")
	}
	return nil, nil
}

func (r *RedisMock) SCARD(key string) (int64, error) {
	if r.MakeErrorResponse {
		return 0, errors.New("Unit Testing Expected Error")
	}
	return 0, nil
}

func (r *RedisMock) SISMEMBERS(key string, field interface{}) (bool, error) {
	if r.MakeErrorResponse {
		return false, errors.New("Unit Testing Expected Error")
	}
	return true, nil
}

func (r *RedisMock) RPUSH(key string, value interface{}) (int64, error) {
	if r.MakeErrorResponse {
		return 0, errors.New("Unit Testing Expected Error")
	}
	return 0, nil
}
