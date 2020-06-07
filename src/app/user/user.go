package user

import (
	"bytes"
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

func (repo userRepo) GetDetailAccount(sessionID string, accountID int64) (MainDetailAccountResponse, error) {
	var result MainDetailAccountResponse

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_ADMIN_GET_DETAIL_ACCOUNT
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
	q.Add("id", strconv.FormatInt(accountID, 10))

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

func (repo userRepo) GetDetailCompany(sessionID string, companyID int64) (MainDetailCompanyResponse, error) {
	var result MainDetailCompanyResponse

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_ADMIN_GET_DETAIL_COMPANY
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
	q.Add("id", strconv.FormatInt(companyID, 10))

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
	q.Add("is_enabled", param.FilterByIsEnable)

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

func (repo userRepo) CreateUser(sessionID string, param CreateAccountParam) (*Error, error) {
	var result *Error

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_ADMIN_CREATE_ACCOUNT
	urlStr := u.String()

	bodyReq, _ := json.Marshal(param)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer(bodyReq))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// set header
	req.Header.Add(headerConst.AUTHORIZATION, sessionID)
	req.Header.Add(headerConst.CONTENT_TYPE, "application/json")

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

	var resultResp = struct {
		Error             *Error `json:"error", omitempty`
		ServerProcessTime string `json:"server_process_time"`
	}{}

	errUnMarshal := json.Unmarshal(body, &resultResp)
	if errUnMarshal != nil {
		log.Println(errUnMarshal)
		return nil, errUnMarshal
	}

	result = resultResp.Error

	return result, nil
}

func (repo userRepo) UpdateUser(sessionID string, param UpdateAccountParam) (*Error, error) {
	var result *Error

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_ADMIN_UPDATE_ACCOUNT
	urlStr := u.String()

	bodyReq, _ := json.Marshal(param)

	resp, err := repo.doRequest(bodyReq, sessionID, urlStr, http.MethodGet)
	if err != nil {
		log.Printf("func Update User error send request: err: %v\n", err)
		return result, err
	}

	var resultResp = struct {
		Error             *Error `json:"error", omitempty`
		ServerProcessTime string `json:"server_process_time"`
	}{}

	errUnMarshal := json.Unmarshal(resp, &resultResp)
	if errUnMarshal != nil {
		log.Println(errUnMarshal)
		return nil, errUnMarshal
	}

	result = resultResp.Error

	return result, nil
}

func (repo userRepo) ChangePassword(sessionID string, param ChangePasswordParam) (*Error, error) {
	var result *Error

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_ADMIN_CHANGE_PASSWORD
	urlStr := u.String()

	bodyReq, _ := json.Marshal(param)

	resp, err := repo.doRequest(bodyReq, sessionID, urlStr, http.MethodGet)
	if err != nil {
		log.Printf("func Update User error send request: err: %v\n", err)
		return result, err
	}

	var resultResp = struct {
		Error             *Error `json:"error", omitempty`
		ServerProcessTime string `json:"server_process_time"`
	}{}

	errUnMarshal := json.Unmarshal(resp, &resultResp)
	if errUnMarshal != nil {
		log.Println(errUnMarshal)
		return nil, errUnMarshal
	}

	result = resultResp.Error

	return result, nil
}

func (repo userRepo) CreateCompany(sessionID string, param CreateCompanyParam) (*Error, error) {
	var result *Error

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_ADMIN_CREATE_COMPANY
	urlStr := u.String()

	bodyReq, _ := json.Marshal(param)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer(bodyReq))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// set header
	req.Header.Add(headerConst.AUTHORIZATION, sessionID)
	req.Header.Add(headerConst.CONTENT_TYPE, "application/json")

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

	var resultResp = struct {
		Error             *Error `json:"error", omitempty`
		ServerProcessTime string `json:"server_process_time"`
	}{}

	errUnMarshal := json.Unmarshal(body, &resultResp)
	if errUnMarshal != nil {
		log.Println(errUnMarshal)
		return nil, errUnMarshal
	}

	result = resultResp.Error

	return result, nil
}

func (repo userRepo) UpdateCompany(sessionID string, param UpdateCompanyParam) (*Error, error) {
	var result *Error

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_ADMIN_UPDATE_COMPANY
	urlStr := u.String()

	bodyReq, _ := json.Marshal(param)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer(bodyReq))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// set header
	req.Header.Add(headerConst.AUTHORIZATION, sessionID)
	req.Header.Add(headerConst.CONTENT_TYPE, "application/json")

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

	var resultResp = struct {
		Error             *Error `json:"error", omitempty`
		ServerProcessTime string `json:"server_process_time"`
	}{}

	errUnMarshal := json.Unmarshal(body, &resultResp)
	if errUnMarshal != nil {
		log.Println(errUnMarshal)
		return nil, errUnMarshal
	}

	result = resultResp.Error

	return result, nil
}

func (repo userRepo) doRequest(param []byte, sessionID, url, method string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(param))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// set header
	req.Header.Add(headerConst.AUTHORIZATION, sessionID)
	req.Header.Add(headerConst.CONTENT_TYPE, "application/json")

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

	return body, nil
}
