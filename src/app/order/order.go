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

func (repo orderRepo) ApproveOrderForAdmin(sessionID string, orderID int64, awb string) (*CreateOrderForAdminResponse, error) {
	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = fmt.Sprintf(urlConst.URL_ADMIN_APPROVE_ORDER, orderID)
	urlStr := u.String()

	bodyReq, _ := json.Marshal(struct {
		AWB string `json:"awb"`
	}{
		AWB: awb,
	})

	resp, err := repo.doRequest(bodyReq, sessionID, urlStr, http.MethodPost)
	if err != nil {
		log.Printf("func CreateOrderForAdmin error send request: err: %v\n", err)
		return nil, err
	}

	var resultResp CreateOrderForAdminResponse
	errUnMarshal := json.Unmarshal(resp, &resultResp)
	if errUnMarshal != nil {
		log.Printf("func CreateOrderForAdmin error when unmarshal response: err: %v. Payload: %+v\n", err, string(resp))
		return nil, errUnMarshal
	}

	return &resultResp, nil
}

func (repo orderRepo) RejectOrderForAdmin(sessionID string, orderID int64) (*CreateOrderForAdminResponse, error) {
	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = fmt.Sprintf(urlConst.URL_ADMIN_REJECT_ORDER, orderID)
	urlStr := u.String()

	resp, err := repo.doRequest(nil, sessionID, urlStr, http.MethodPost)
	if err != nil {
		log.Printf("func RejectOrderForAdmin error send request: err: %v\n", err)
		return nil, err
	}

	var resultResp CreateOrderForAdminResponse
	errUnMarshal := json.Unmarshal(resp, &resultResp)
	if errUnMarshal != nil {
		log.Printf("func RejectOrderForAdmin error when unmarshal response: err: %v. Payload: %+v\n", err, string(resp))
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
		log.Printf("func GetListOrders error create request. error: %v", err)
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

func (repo orderRepo) DispatchOrderToFulfilmentCenter(sessionID string, orderID int64) (*SuccessCRUDResponse, error) {
	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = fmt.Sprintf(urlConst.URL_ADMIN_DISPATCH_ORDER_FULFILMENT_CENTER, orderID)
	urlStr := u.String()

	bodyReq, _ := json.Marshal(struct{}{})

	resp, err := repo.doRequest(bodyReq, sessionID, urlStr, http.MethodPost)
	if err != nil {
		log.Printf("func DispatchOrderToFulfilmentCenter error send request: err: %v\n", err)
		return nil, err
	}

	var resultResp SuccessCRUDResponse
	errUnMarshal := json.Unmarshal(resp, &resultResp)
	if errUnMarshal != nil {
		log.Printf("func DispatchOrderToFulfilmentCenter error when unmarshal response: err: %v. Payload: %+v\n", err, string(resp))
		return nil, errUnMarshal
	}

	return &resultResp, nil
}

func (repo orderRepo) DispatchOrderToDriver(sessionID string, orderID int64) (*SuccessCRUDResponse, error) {
	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = fmt.Sprintf(urlConst.URL_ADMIN_DISPATCH_ORDER_PICK_UP, orderID)
	urlStr := u.String()

	bodyReq, _ := json.Marshal(struct{}{})

	resp, err := repo.doRequest(bodyReq, sessionID, urlStr, http.MethodPost)
	if err != nil {
		log.Printf("func DispatchOrderToDriver error send request: err: %v\n", err)
		return nil, err
	}

	var resultResp SuccessCRUDResponse
	errUnMarshal := json.Unmarshal(resp, &resultResp)
	if errUnMarshal != nil {
		log.Printf("func DispatchOrderToDriver error when unmarshal response: err: %v. Payload: %+v\n", err, string(resp))
		return nil, errUnMarshal
	}

	return &resultResp, nil
}

func (repo orderRepo) PickUpOrderToDriver(sessionID string, orderID int64, param PickUpParam) (*SuccessCRUDResponse, error) {
	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = fmt.Sprintf(urlConst.URL_ADMIN_PICKUP_DRIVER, orderID)
	urlStr := u.String()

	bodyReq, _ := json.Marshal(struct {
		DriverPickUP PickUpParam `json:"driver_pickup"`
	}{
		DriverPickUP: param,
	})

	resp, err := repo.doRequest(bodyReq, sessionID, urlStr, http.MethodPost)
	if err != nil {
		log.Printf("func PickUpOrderToDriver error send request: err: %v\n", err)
		return nil, err
	}

	var resultResp SuccessCRUDResponse
	errUnMarshal := json.Unmarshal(resp, &resultResp)
	if errUnMarshal != nil {
		log.Printf("func PickUpOrderToDriver error when unmarshal response: err: %v. Payload: %+v\n", err, string(resp))
		return nil, errUnMarshal
	}

	return &resultResp, nil
}

func (repo orderRepo) RejectPickUpOrder(sessionID string, orderID int64, pickupID ...int64) (*SuccessCRUDResponse, error) {
	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = fmt.Sprintf(urlConst.URL_ADMIN_PICKUP_REJECT, orderID)
	urlStr := u.String()

	bodyReq, _ := json.Marshal(struct {
		DriverPickUP struct {
			PickupIDs []int64 `json:"pickup_ids"`
		} `json:"driver_pickup"`
	}{
		DriverPickUP: struct {
			PickupIDs []int64 `json:"pickup_ids"`
		}{
			PickupIDs: pickupID,
		},
	})

	resp, err := repo.doRequest(bodyReq, sessionID, urlStr, http.MethodPost)
	if err != nil {
		log.Printf("func RejectPickUpOrder error send request: err: %v\n", err)
		return nil, err
	}

	var resultResp SuccessCRUDResponse
	errUnMarshal := json.Unmarshal(resp, &resultResp)
	if errUnMarshal != nil {
		log.Printf("func RejectPickUpOrder error when unmarshal response: err: %v. Payload: %+v\n", err, string(resp))
		return nil, errUnMarshal
	}

	return &resultResp, nil
}

func (repo orderRepo) FinishPickUpOrder(sessionID string, orderID int64, param FinishPickupParam) (*SuccessCRUDResponse, error) {
	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = fmt.Sprintf(urlConst.URL_ADMIN_PICKUP_FINISH, orderID)
	urlStr := u.String()

	bodyReq, _ := json.Marshal(struct {
		DriverPickup FinishPickupParam `json:"driver_pickup"`
	}{
		DriverPickup: param,
	})

	resp, err := repo.doRequest(bodyReq, sessionID, urlStr, http.MethodPost)
	if err != nil {
		log.Printf("func RejectPickUpOrder error send request: err: %v\n", err)
		return nil, err
	}

	var resultResp SuccessCRUDResponse
	errUnMarshal := json.Unmarshal(resp, &resultResp)
	if errUnMarshal != nil {
		log.Printf("func RejectPickUpOrder error when unmarshal response: err: %v. Payload: %+v\n", err, string(resp))
		return nil, errUnMarshal
	}

	return &resultResp, nil
}

func (repo orderRepo) DeliveryOrder(sessionID string, orderID int64, param DeliveryParam) (*SuccessCRUDResponse, error) {
	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = fmt.Sprintf(urlConst.URL_ADMIN_DELIVER_ORDER, orderID)
	urlStr := u.String()

	bodyReq, _ := json.Marshal(struct {
		Payload DeliveryParam `json:"driver_delivery"`
	}{
		Payload: param,
	})

	resp, err := repo.doRequest(bodyReq, sessionID, urlStr, http.MethodPost)
	if err != nil {
		log.Printf("func DeliveryOrder error send request: err: %v\n", err)
		return nil, err
	}

	var resultResp SuccessCRUDResponse
	errUnMarshal := json.Unmarshal(resp, &resultResp)
	if errUnMarshal != nil {
		log.Printf("func DeliveryOrder error when unmarshal response: err: %v. Payload: %+v\n", err, string(resp))
		return nil, errUnMarshal
	}

	return &resultResp, nil
}

func (repo orderRepo) TransitOnCGXOrder(sessionID string, orderID int64) (*SuccessCRUDResponse, error) {
	u, _ := url.ParseRequestURI(conf.GConfig.BaseUrlConfig.ProductDNS)
	u.Path = fmt.Sprintf(urlConst.URL_ADMIN_ON_CGX_ORDER, orderID)
	urlStr := u.String()

	bodyReq, _ := json.Marshal(struct{}{})

	resp, err := repo.doRequest(bodyReq, sessionID, urlStr, http.MethodPost)
	if err != nil {
		log.Printf("func TransitOnCGXOrder error send request: err: %v\n", err)
		return nil, err
	}

	var resultResp SuccessCRUDResponse
	errUnMarshal := json.Unmarshal(resp, &resultResp)
	if errUnMarshal != nil {
		log.Printf("func TransitOnCGXOrder error when unmarshal response: err: %v. Payload: %+v\n", err, string(resp))
		return nil, errUnMarshal
	}

	return &resultResp, nil
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
