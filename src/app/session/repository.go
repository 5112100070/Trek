package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/5112100070/Trek/src/conf"
	headerConst "github.com/5112100070/Trek/src/constants/header"
	urlConst "github.com/5112100070/Trek/src/constants/url"
	redigo "github.com/5112100070/Trek/src/global/redis"
)

func InitSessionRepo(sessionRedis redigo.Redis) *sessionRepo {
	return &sessionRepo{
		redis: sessionRedis,
	}
}

func (repo sessionRepo) GetUser(sessionID string) (AccountResponse, error) {
	var result AccountResponse

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_GET_USER_PROFILE
	urlStr := u.String()

	client := &http.Client{}
	req, err := http.NewRequest("GET", urlStr, strings.NewReader(""))
	if err != nil {
		log.Println(err)
		return result, err
	}

	req.Header.Add(headerConst.AUTHORIZATION, sessionID)

	resp, errGetResp := client.Do(req)
	if err != nil {
		log.Println(errGetResp)
		return result, errGetResp
	}

	if resp.Body == nil {
		log.Println("no response from cgx service")
		return result, errors.New("no response from cgx service")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer resp.Body.Close()

	errUnMarshal := json.Unmarshal(body, &result)
	if errUnMarshal != nil {
		log.Println(errUnMarshal)
		return result, errUnMarshal
	}

	return result, nil
}

func (repo sessionRepo) SetUser(cookie string, user Account) error {
	key := fmt.Sprintf("%v%v", redis_key_cookie, cookie)

	m, _ := json.Marshal(user)

	err := repo.redis.SETEX(key, redis_timeout, string(m))
	if err != nil {
		return err
	}

	return nil
}

func (repo sessionRepo) DelUser(cookie string) error {
	key := fmt.Sprintf("%v%v", redis_key_cookie, cookie)

	err := repo.redis.DEL(key)
	if err != nil {
		return err
	}
	return nil
}
