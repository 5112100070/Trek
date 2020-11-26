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
	statusConst "github.com/5112100070/Trek/src/constants/status"
	urlConst "github.com/5112100070/Trek/src/constants/url"
)

// InitOrderRepo - initialization for orderRepo
func InitOrderRepo() *orderRepo {
	return &orderRepo{}
}

func (repo orderRepo) GetListOrderStatusCGX() (map[string]string, error) {
	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_DESC_GET_LIST_ORDER_STATUS
	urlStr := u.String()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		log.Println("func GetListOrderStatusCGX error when create new request. Error: ", err)
		return nil, err
	}

	resp, errGetResp := client.Do(req)
	if err != nil {
		log.Println("func GetListOrderStatusCGX error error when do resp. Error: ", errGetResp)
		return nil, errGetResp
	}

	if resp == nil || resp.Body == nil {
		log.Println("func GetListOrderStatusCGX no response from cgx service")
		return nil, errors.New("func GetListOrderStatusCGX no response from cgx service")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("func GetListOrderStatusCGX error read response. Error: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	var resultResp struct {
		Data  map[string]string `json:"data", omitempty`
		Error *ErrorOrder       `json:"error", omitempty`
	}
	errUnMarshal := json.Unmarshal(body, &resultResp)
	if errUnMarshal != nil {
		log.Printf("func GetListOrderStatusCGX error when unmarshal response: err: %v. Payload: %+v\n", errUnMarshal, string(body))
		return nil, errUnMarshal
	}

	if resultResp.Error != nil {
		log.Printf("func GetListOrderStatusCGX error from CGX. Code: %d, error: %v", resultResp.Error.Code, resultResp.Error.Detail)
		return nil, errors.New(resultResp.Error.Detail)
	}

	// if not have result from cgx. must return error
	if resultResp.Data == nil {
		log.Printf("func GetListOrderStatusCGX not have response from CGX")
		return nil, errors.New("GetListOrderStatusCGX not have response from CGX")
	}

	return resultResp.Data, nil
}

func (repo orderRepo) GetListPickupStatusCGX() (map[string]string, error) {
	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = urlConst.URL_DESC_GET_LIST_PICKUP_STATUS
	urlStr := u.String()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		log.Println("func GetListPickupStatusCGX error when create new request. Error: ", err)
		return nil, err
	}

	resp, errGetResp := client.Do(req)
	if err != nil {
		log.Println("func GetListPickupStatusCGX error error when do resp. Error: ", errGetResp)
		return nil, errGetResp
	}

	if resp == nil || resp.Body == nil {
		log.Println("func GetListPickupStatusCGX no response from cgx service")
		return nil, errors.New("func GetListPickupStatusCGX no response from cgx service")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("func GetListPickupStatusCGX error read response. Error: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	var resultResp struct {
		Data  map[string]string `json:"data", omitempty`
		Error *ErrorOrder       `json:"error", omitempty`
	}
	errUnMarshal := json.Unmarshal(body, &resultResp)
	if errUnMarshal != nil {
		log.Printf("func GetListPickupStatusCGX error when unmarshal response: err: %v. Payload: %+v\n", errUnMarshal, string(body))
		return nil, errUnMarshal
	}

	if resultResp.Error != nil {
		log.Printf("func GetListPickupStatusCGX error from CGX. Code: %d, error: %v", resultResp.Error.Code, resultResp.Error.Detail)
		return nil, errors.New(resultResp.Error.Detail)
	}

	// if not have result from cgx. must return error
	if resultResp.Data == nil {
		log.Printf("func GetListPickupStatusCGX not have response from CGX")
		return nil, errors.New("GetListPickupStatusCGX not have response from CGX")
	}

	return resultResp.Data, nil
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

	// Generate Fully Address
	result.ReceiverAddrDisplay = result.ReceiverAddress

	// Generate RT
	if result.ReceiverRT > 0 {
		result.ReceiverAddrDisplay = fmt.Sprintf("%s, RT: %d", result.ReceiverAddrDisplay, result.ReceiverRT)
	}

	// Generate RW
	if result.ReceiverRW > 0 {
		result.ReceiverAddrDisplay = fmt.Sprintf("%s, RW: %d", result.ReceiverAddrDisplay, result.ReceiverRW)
	}

	result.ReceiverAddrDisplay = fmt.Sprintf("%s, Kec. %s, Kel. %s, %s, %s, %d", result.ReceiverAddrDisplay, result.ReceiverKecamatan,
		result.ReceiverKelurahan, result.ReceiverCity, result.ReceiverProv, result.ReceiverZIP)

	for i, _ := range result.Pickups {
		result.Pickups[i].CreateTimeStr = result.Pickups[i].CreateTime.Format("02 Jan 2006 - 15:04:05")
		result.Pickups[i].UpdateTimeStr = result.Pickups[i].UpdateTime.Format("02 Jan 2006 - 15:04:05")
		result.Pickups[i].TotalItems = len(result.Pickups[i].Items)

		result.Pickups[i].FullAddress = result.Pickups[i].Address

		// Generate RT for pickup
		if result.Pickups[i].AddrRT > 0 {
			result.Pickups[i].FullAddress = fmt.Sprintf("%s, RT: %d", result.Pickups[i].FullAddress, result.Pickups[i].AddrRT)
		}

		// Generate RW for pickup
		if result.Pickups[i].AddrRW > 0 {
			result.Pickups[i].FullAddress = fmt.Sprintf("%s, RW: %d", result.Pickups[i].FullAddress, result.Pickups[i].AddrRW)
		}

		result.Pickups[i].FullAddress = fmt.Sprintf("%s, Kec. %s, Kel. %s, %s, %s, %d", result.Pickups[i].FullAddress, result.Pickups[i].AddrKecamatan,
			result.Pickups[i].AddrKelurahan, result.Pickups[i].AddrCity, result.Pickups[i].AddrProv, result.Pickups[i].ZIP)

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

	// build param
	limit := int64(param.Rows)
	offset := int64(((param.Page - 1) * param.Rows) + 1)

	var option struct {
		CompanyID *int64  `json:"company_id"`
		Offset    *int64  `json:"offset"`
		Limit     *int64  `json:"limit"`
		SortType  *string `json:"sort_type"`
	}

	option.Limit = &limit
	option.Offset = &offset
	option.SortType = &param.OrderType

	log.Println(urlStr)
	log.Println(limit)
	log.Println(offset)
	log.Println(param.OrderType)

	if param.CompanyID > 0 {
		companyID := int64(param.CompanyID)
		option.CompanyID = &companyID
		log.Println(companyID)
	}

	payload, _ := json.Marshal(option)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, urlStr, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("func GetListOrders error create request. error: %v", err)
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
		log.Printf("func GetListOrders error when readALL response: %v. error: %v", resp.Body, err)
		return result, err
	}
	defer resp.Body.Close()

	errUnMarshal := json.Unmarshal(body, &result)
	if errUnMarshal != nil {
		log.Printf("func GetListOrders error hit:%v when unmarshal: %v. error: %v", u.String(), string(body), errUnMarshal)
		return result, errUnMarshal
	}

	for i, _ := range result.Data.Orders {
		result.Data.Orders[i].ArrivedTimeStr = result.Data.Orders[i].ArrivedTime.Format("02 Jan 2006")
		result.Data.Orders[i].UpdateTimeStr = result.Data.Orders[i].UpdateTime.Format("02 Jan 2006")
		result.Data.Orders[i].TotalPickUp = len(result.Data.Orders[i].Pickups)
		result.Data.Orders[i].StatusBadge = statusConst.MAP_BADGE_BY_STATUS_ORDER[result.Data.Orders[i].Status]
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
		log.Println("func GetListUnitInOrder error when create request: ", err)
		return result, err
	}

	// set header
	req.Header.Add(headerConst.AUTHORIZATION, sessionID)

	resp, errGetResp := client.Do(req)
	if err != nil {
		log.Printf("func GetListUnitInOrder when request: %v, error: %v", urlStr, errGetResp)
		return result, errGetResp
	}

	if resp == nil || resp.Body == nil {
		log.Println("func GetListUnitInOrder no response from cgx service")
		return result, errors.New("no response from cgx service")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("func GetListUnitInOrder when read request", err)
		return result, err
	}
	defer resp.Body.Close()

	errUnMarshal := json.Unmarshal(body, &result)
	if errUnMarshal != nil {
		log.Println("func GetListUnitInOrder when failed unmarshal: ", errUnMarshal)
		return result, errUnMarshal
	}

	return result, nil
}

func (repo orderRepo) doRequest(param []byte, sessionID, url, method string) ([]byte, error) {
	client := &http.Client{}
	if param == nil {
		param = []byte("{}")
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(param))
	if err != nil {
		log.Println("func doRequest error when dorequest: ", err)
		return nil, err
	}

	// set header
	if sessionID != "" {
		req.Header.Add(headerConst.AUTHORIZATION, sessionID)
	}
	req.Header.Add(headerConst.CONTENT_TYPE, "application/json")

	resp, errGetResp := client.Do(req)
	if err != nil {
		log.Println(errGetResp)
		return nil, errGetResp
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("func doRequest error when read response: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	return body, nil
}
