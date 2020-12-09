package order

import "time"

type orderRepo struct {
}

type ErrorOrder struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

type CreateOrderForAdminResponse struct {
	Error *ErrorOrder `json:"error, omitempty"`
	Data  *struct {
		Success bool  `json:"success"`
		OrderID int64 `json:"order_id"`
	} `json:"data, omitempty"`
	ServerProcessTime string `json:"server_process_time"`
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
	ActualReceiverName  string           `json:"actual_receiver_name"`
	ReceiverName        string           `json:"receiver_name""`
	ReceiverAddress     string           `json:"receiver_address"`
	ReceiverKecamatan   string           `json:"receiver_kecamatan"`
	ReceiverKelurahan   string           `json:"receiver_kelurahan"`
	ReceiverCity        string           `json:"receiver_kota"`
	ReceiverProv        string           `json:"receiver_provinsi"`
	ReceiverRT          int64            `json:"receiver_rt"`
	ReceiverRW          int64            `json:"receiver_rw"`
	ReceiverZIP         int64            `json:"receiver_zip"`
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
	ArrivedTime         time.Time        `json:"arrived_time"`
	PickUpDeadline      time.Time        `json:"pickup_deadline"`
	Pickups             []PickUpResponse `json:"pickups"`

	// This variable only used for displaying to user
	TotalPickUp         int
	ArrivedTimeStr      string
	CreateTimeStr       string
	UpdateTimeStr       string
	PickUpDeadlineStr   string
	ReceiverAddrDisplay string
	StatusBadge         string
}

type PickUpResponse struct {
	ID            int64          `json:"id"`
	Name          string         `json:"name"`
	PIC           string         `json:"pic"`
	Address       string         `json:"address"`
	AddrRT        int64          `json:"rt"`
	AddrRW        int64          `json:"rw"`
	AddrKecamatan string         `json:"kecamatan"`
	AddrKelurahan string         `json:"kelurahan"`
	AddrCity      string         `json:"kota"`
	AddrProv      string         `json:"provinsi"`
	ZIP           int64          `json:"zip"`
	PhoneNumber   string         `json:"phone_number"`
	DriverName    string         `json:"driver_name"`
	DriverPhone   string         `json:"driver_phone"`
	Status        int            `json:"status"`
	StatusName    string         `json:"status_name"`
	UpdateBy      int64          `json:"update_by"`
	UpdateTime    time.Time      `json:"update_time"`
	CreateTime    time.Time      `json:"create_time"`
	Notes         string         `json:"notes"`
	Items         []ItemResponse `json:"items"`

	// This variable only used for displaying to user
	TotalItems    int
	CreateTimeStr string
	UpdateTimeStr string
	FullAddress   string
}

type ItemResponse struct {
	ID           int64     `json:"id"`
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
