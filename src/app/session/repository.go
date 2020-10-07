package session

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
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
	req, err := http.NewRequest(http.MethodGet, urlStr, strings.NewReader(""))
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

	if resp == nil || resp.Body == nil {
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

func (repo sessionRepo) RequestLogout(sessionID string) (*Error, error) {
	cgxResponse := struct {
		SuccessResp *struct {
			IsSuccess bool   `json:"is_success"`
			Message   string `json:"message"`
		} `json:"data"`
		ErrResp           *Error `json:"error"`
		ServerProcessTime string `json:"server_process_time"`
	}{}

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_REQUEST_LOGOUT
	urlStr := u.String()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(""))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Add(headerConst.AUTHORIZATION, sessionID)

	resp, errGetResp := client.Do(req)
	if err != nil {
		log.Println(errGetResp)
		return nil, errGetResp
	}

	if resp == nil || resp.Body == nil {
		log.Println("no response from cgx service")
		return nil, errors.New("no response from cgx service")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	errUnMarshal := json.Unmarshal(body, &cgxResponse)
	if errUnMarshal != nil {
		log.Println(errUnMarshal)
		return nil, errUnMarshal
	}

	return cgxResponse.ErrResp, nil
}

func (repo sessionRepo) RequestLogin(email, password string) (LoginResponse, error) {
	var result LoginResponse

	data := url.Values{}
	data.Set("email", email)
	data.Set("password", password)

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_REQUEST_LOGIN
	urlStr := u.String()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		log.Println(err)
		return result, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, errGetResp := client.Do(req)
	if err != nil {
		log.Println(errGetResp)
		return result, errGetResp
	}

	if resp == nil || resp.Body == nil {
		return result, errors.New("no response from cgx service")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return result, err
	}

	errUnMarshal := json.Unmarshal(body, &result)
	if errUnMarshal != nil {
		log.Println(errUnMarshal)
		return result, errUnMarshal
	}

	return result, nil
}

func (repo sessionRepo) CheckFeature(sessionID, pathURL, method string) (FeatureCheckResponse, error) {
	var result FeatureCheckResponse
	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_CHECK_FEATURE
	urlStr := u.String()

	bodyPayload := struct {
		PathURL string `json:"path_url"`
		Method  string `json:"method"`
	}{
		PathURL: pathURL,
		Method:  method,
	}

	payload, _ := json.Marshal(bodyPayload)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer(payload))
	if err != nil {
		log.Println("func CheckFeature error when create request: ", err)
		return result, err
	}

	// set header
	req.Header.Add(headerConst.AUTHORIZATION, sessionID)

	resp, errGetResp := client.Do(req)
	if err != nil {
		log.Println(errGetResp)
		return result, errGetResp
	}

	if resp == nil || resp.Body == nil {
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
		log.Println("func CheckFeature error when unMarshal: ", errUnMarshal)
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
