package user

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/5112100070/Trek/src/conf"
	headerConst "github.com/5112100070/Trek/src/constants/header"
	urlConst "github.com/5112100070/Trek/src/constants/url"
)

// InitUserRepo - initialization for userRepo
func InitUserRepo() *userRepo {
	return &userRepo{}
}

func (repo userRepo) GetListUsers(sessionID string, param ListUserParam) (MainListAccountResponse, error) {
	var result MainListAccountResponse

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_ADMIN_GET_LIST_USER
	urlStr := u.String()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, urlStr, strings.NewReader(""))
	if err != nil {
		log.Println(err)
		return result, err
	}

	// set header
	req.Header.Add(headerConst.AUTHORIZATION, sessionID)
	// set query param
	q := req.URL.Query()
	q.Add("rows", strconv.FormatInt(int64(param.Rows), 10))
	q.Add("page", strconv.FormatInt(int64(param.Page), 10))
	q.Add("order_by", param.OrderType)

	req.URL.RawQuery = q.Encode()

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

func (repo userRepo) GetListCompany(sessionID string, param ListCompanyParam) (MainListCompanyResponse, error) {
	var result MainListCompanyResponse

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_ADMIN_GET_LIST_COMPANY
	urlStr := u.String()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, urlStr, strings.NewReader(""))
	if err != nil {
		log.Println(err)
		return result, err
	}

	// set header
	req.Header.Add(headerConst.AUTHORIZATION, sessionID)
	// set query param
	q := req.URL.Query()
	q.Add("rows", strconv.FormatInt(int64(param.Rows), 10))
	q.Add("page", strconv.FormatInt(int64(param.Page), 10))
	q.Add("order_by", param.OrderType)

	req.URL.RawQuery = q.Encode()

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
