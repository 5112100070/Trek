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
	"time"

	"github.com/5112100070/Trek/src/conf"
	headerConst "github.com/5112100070/Trek/src/constants/header"
	urlConst "github.com/5112100070/Trek/src/constants/url"
	"github.com/5112100070/publib/encoding"
	publibTime "github.com/5112100070/publib/time"
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
		log.Printf("func GetDetailAccount error when unmarshal: %v, payload: %v \n", errUnMarshal, string(body))
		return result, errUnMarshal
	}

	return result, nil
}

func (repo userRepo) GetDetailCompany(sessionID string, param DetailCompanyParam) (MainDetailCompanyResponse, error) {
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
	headerTime := publibTime.GetTimeWIB().Format(time.RFC822)
	req.Header.Add(headerConst.DATE, headerTime)

	if param.IsInternal {
		req.Header.Add(headerConst.USER_AGENT, "cgx")

		stdHeaderTime, _ := time.Parse(time.RFC822, headerTime)
		hmacPayload := encoding.HMACAuthData{
			Method: http.MethodGet,
			Date:   stdHeaderTime.Unix(),
			Path:   urlConst.URL_ADMIN_GET_DETAIL_COMPANY,
		}

		req.Header.Add(headerConst.PROXY_AUTHORIZATION, hmacPayload.GenerateHMACHash(""))
	}

	// set query param
	q := req.URL.Query()
	q.Add("id", strconv.FormatInt(param.CompanyID, 10))

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

	if param.CompanyID != 0 {
		q.Add("company_id", strconv.FormatInt(param.CompanyID, 10))
	}

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
	headerTime := publibTime.GetTimeWIB().Format(time.RFC822)
	req.Header.Add(headerConst.DATE, headerTime)

	if param.IsInternal {
		req.Header.Add(headerConst.USER_AGENT, "cgx")

		stdHeaderTime, _ := time.Parse(time.RFC822, headerTime)
		hmacPayload := encoding.HMACAuthData{
			Method: http.MethodGet,
			Date:   stdHeaderTime.Unix(),
			Path:   urlConst.URL_ADMIN_GET_LIST_COMPANY,
		}

		req.Header.Add(headerConst.PROXY_AUTHORIZATION, hmacPayload.GenerateHMACHash(""))
	}

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
		log.Printf("func UpdateUser error send request: err: %v\n", err)
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

func (repo userRepo) ChangePassword(sessionID string, newPassword string) (*Error, error) {
	var result *Error

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_CHANGE_PASSWORD
	urlStr := u.String()

	bodyReq, _ := json.Marshal(struct {
		NewPassword string `json:"new_password"`
	}{
		NewPassword: newPassword,
	})

	resp, err := repo.doRequest(bodyReq, sessionID, urlStr, http.MethodGet)
	if err != nil {
		log.Printf("func ChangePassword error send request: err: %v\n", err)
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

func (repo userRepo) ChangePasswordAdmin(sessionID string, param ChangePasswordParam) (*Error, error) {
	var result *Error

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_ADMIN_CHANGE_PASSWORD
	urlStr := u.String()

	bodyReq, _ := json.Marshal(param)

	resp, err := repo.doRequest(bodyReq, sessionID, urlStr, http.MethodGet)
	if err != nil {
		log.Printf("func ChangePasswordAdmin error send request: err: %v\n", err)
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

func (repo userRepo) ChangeStatusAccount(sessionID string, param ChangeStatusAccParam) (*Error, error) {
	var result *Error

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_ADMIN_CHANGE_CHANGE_STATUS_ACTIVATION
	urlStr := u.String()

	bodyReq, _ := json.Marshal(param)

	resp, err := repo.doRequest(bodyReq, sessionID, urlStr, http.MethodGet)
	if err != nil {
		log.Printf("func Change Status Account error send request: err: %v\n", err)
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
