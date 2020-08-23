package order

import "time"

type orderRepo struct {
}

type ErrorOrder struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

type CreateOrderForAdminResponse struct {
	Error             *ErrorOrder `json:"error", omitempty`
	Status            string      `json:"status"`
	ServerProcessTime string      `json:"server_process_time"`
}

type ApproveOrderForAdminResponse struct {
	Error             *ErrorOrder `json:"error", omitempty`
	Success           bool        `json:"success"`
	ServerProcessTime string      `json:"server_process_time"`
}

type SuccessCRUDResponse struct {
	Error             *ErrorOrder `json:"error", omitempty`
	Success           bool        `json:"success"`
	ServerProcessTime string      `json:"server_process_time"`
}

type MainListOrderResponse struct {
	Data struct {
		Orders []OrderReponse `json:"order"`
		Next   bool           `json: "next"`
	} `json:"data", omitempty`
	Error             *ErrorOrder `json:"error", omitempty`
	ServerProcessTime string      `json:"server_process_time"`
}

type MainListUnitResponse struct {
	Data              map[int]string `json:"data", omitempty`
	Error             *ErrorOrder    `json:"error", omitempty`
	ServerProcessTime string         `json:"server_process_time"`
}

type OrderReponse struct {
	ID                  int64            `json:"id"`
	AWB                 string           `json:"awb"`
	RequestorID         int64            `json:"requestor_id"`
	ReceiverName        string           `json:"receiver_name""`
	ReceiverAddress     string           `json:"receiver_address"`
	ReceiverPhoneNumber string           `json:"receiver_phone_number"`
	ReceiverNotes       string           `json:"receiver_notes"`
	CompanyID           int64            `json:"company_id"`
	CompanyName         string           `json:"company_name"`
	Status              int              `json:"status"`
	StatusName          string           `json:"status_name"`
	UpdateBy            int64            `json:"update_by"`
	CreateTime          time.Time        `json:"create_time"`
	UpdateTime          time.Time        `json:"update_time"`
	DeliveryType        int64            `json:"delivery_type"`
	DeliveryName        string           `json:"delivery_name"`
	Pickups             []PickUpResponse `json:"pickups"`

	// This variable only used for displaying to user
	TotalPickUp   int
	CreateTimeStr string
	UpdateTimeStr string
	StatusBadge   string
}

type PickUpResponse struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Address     string         `json:"address"`
	PhoneNumber string         `json:"phone_number"`
	DriverName  string         `json:"driver_name"`
	DriverPhone string         `json:"driver_phone"`
	Status      int            `json:"status"`
	UpdateBy    int64          `json:"update_by"`
	UpdateTime  time.Time      `json:"update_time"`
	CreateTime  time.Time      `json:"create_time"`
	Notes       string         `json:"notes"`
	Items       []ItemResponse `json:"items"`

	// This variable only used for displaying to user
	TotalItems    int
	CreateTimeStr string
	UpdateTimeStr string
}

type ItemResponse struct { 
	ID 	int64 `json:"id"`
	Name         string    `json:"name"`
	Quantity     int64     `json:"quantity"`
	Unit         int64     `json:"unit"`
	UnitName     string    `json:"unit_name"`
	Notes        string    `json:"notes"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
	PickUpTime   time.Time `json:"pickup_time"`
	DeadlineTime time.Time `json:"deadline"`

	// This variable only used for displaying to user
	CreateTimeStr string
	UpdateTimeStr string
	DeadlineStr   string
	PickupTimeStr string
}
 
// Handling standart unit data
type Unit struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
}
