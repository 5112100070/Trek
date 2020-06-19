package order

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

// InitOrderRepo - initialization for orderRepo
func InitOrderRepo() *orderRepo {
	return &orderRepo{}
}

func (repo orderRepo) CreateOrderForAdmin(sessionID string, payload CreateOrderParam) (*CreateOrderForAdminResponse, error) {
	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_ADMIN_CREATE_ORDER_FOR_ADMIN
	urlStr := u.String()

	bodyReq, _ := json.Marshal(payload)

	resp, err := repo.doRequest(bodyReq, sessionID, urlStr, http.MethodPost)
	if err != nil {
		log.Printf("func CreateOrderForAdmin error send request: err: %v\n", err)
		return nil, err
	}

	var resultResp CreateOrderForAdminResponse
	errUnMarshal := json.Unmarshal(resp, &resultResp)
	if errUnMarshal != nil {
		log.Println(errUnMarshal)
		return nil, errUnMarshal
	}

	return &resultResp, nil
}

func (repo orderRepo) GetListOrders(sessionID string, param ListOrderParam) (MainListOrderResponse, error) {
	var result MainListOrderResponse

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_ADMIN_GET_ORDER
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

	for i, _ := range result.Data {
		result.Data[i].UpdateTimeStr = result.Data[i].UpdateTime.Format("02 Jan 2006")
		result.Data[i].TotalPickUp = len(result.Data[i].Pickups)
	}

	return result, nil
}

func (repo orderRepo) doRequest(param []byte, sessionID, url, method string) ([]byte, error) {
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	return body, nil
}
