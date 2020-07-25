package order

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

func (repo orderRepo) GetOrderDetailForAdmin(sessionID string, orderID int64) (OrderReponse, *ErrorOrder, error) {
	var result OrderReponse
	var respOrder MainListOrderResponse

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_ADMIN_GET_ORDER_DETAIL
	urlStr := u.String()

	bodyPayload := struct {
		OrderID int64 `json:"order_id"`
	}{
		OrderID: orderID,
	}

	payload, _ := json.Marshal(bodyPayload)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, urlStr, bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err)
		return result, nil, err
	}

	// set header
	req.Header.Add(headerConst.AUTHORIZATION, sessionID)
	// set query param
	q := req.URL.Query()
	q.Add("order_id", strconv.FormatInt(orderID, 10))

	req.URL.RawQuery = q.Encode()

	resp, errGetResp := client.Do(req)
	if err != nil {
		log.Println(errGetResp)
		return result, nil, errGetResp
	}

	if resp == nil || resp.Body == nil {
		log.Println("no response from cgx service")
		return result, nil, errors.New("no response from cgx service")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return result, nil, err
	}
	defer resp.Body.Close()

	errUnMarshal := json.Unmarshal(body, &respOrder)
	if errUnMarshal != nil {
		log.Println(errUnMarshal)
		return result, nil, errUnMarshal
	}

	if respOrder.Error != nil {
		return result, respOrder.Error, errors.New("no result")
	}

	if len(respOrder.Data.Orders) <= 0 {
		return result, nil, nil
	}

	// generate display data
	result = respOrder.Data.Orders[0]
	result.TotalPickUp = len(result.Pickups)
	result.CreateTimeStr = fmt.Sprintf("%s, %s", result.CreateTime.Weekday(), result.CreateTime.Format("02 Jan 2006 - 15:04:05"))
	result.UpdateTimeStr = result.UpdateTime.Format("02 Jan 2006 - 15:04:05")

	for i, _ := range result.Pickups {
		result.Pickups[i].CreateTimeStr = result.Pickups[i].CreateTime.Format("02 Jan 2006 - 15:04:05")
		result.Pickups[i].UpdateTimeStr = result.Pickups[i].UpdateTime.Format("02 Jan 2006 - 15:04:05")
		result.Pickups[i].TotalItems = len(result.Pickups[i].Items)

		for j, _ := range result.Pickups[i].Items {
			result.Pickups[i].Items[j].CreateTimeStr = result.Pickups[i].Items[j].CreateTime.Format("02 Jan 2006 - 15:04:05")
			result.Pickups[i].Items[j].UpdateTimeStr = result.Pickups[i].Items[j].UpdateTime.Format("02 Jan 2006 - 15:04:05")
			result.Pickups[i].Items[j].PickupTimeStr = result.Pickups[i].Items[j].PickUpTime.Format("02 Jan 2006 - 15:04:05")
			result.Pickups[i].Items[j].DeadlineStr = result.Pickups[i].Items[j].DeadlineTime.Format("02 Jan 2006 - 15:04:05")
		}
	}

	return result, nil, nil
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

	for i, _ := range result.Data.Orders {
		result.Data.Orders[i].UpdateTimeStr = result.Data.Orders[i].UpdateTime.Format("02 Jan 2006")
		result.Data.Orders[i].TotalPickUp = len(result.Data.Orders[i].Pickups)
	}

	return result, nil
}

func (repo orderRepo) GetListUnitInOrder(sessionID string) (MainListUnitResponse, error) {
	var result MainListUnitResponse

	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_ADMIN_GET_UNIT_ORDER
	urlStr := u.String()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, urlStr, strings.NewReader(""))
	if err != nil {
		log.Println(err)
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
		log.Println(errUnMarshal)
		return result, errUnMarshal
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
