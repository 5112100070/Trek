package order

import "time"

type orderRepo struct {
}

type CreateOrderForAdminResponse struct {
	Error             *string `json:"error", omitempty`
	Status            string  `json:"status"`
	ServerProcessTime string  `json:"server_process_time"`
}

type MainListOrderResponse struct {
	Data              []OrderReponse `json:"data", omitempty`
	TotalOrder        int            `json: "total_order"`
	Error             *string        `json:"error", omitempty`
	ServerProcessTime string         `json:"server_process_time"`
}

type OrderReponse struct {
	ID              int64            `json:"id"`
	RequestorID     int64            `json:"requestor_id"`
	ReceiverName    string           `json:"receiver_name""`
	ReceiverAddress string           `json:"receiver_address"`
	ReceiverNotes   string           `json:"receiver_notes"`
	Status          int              `json:"status"`
	UpdateBy        int64            `json:"update_by"`
	CreateTime      time.Time        `json:"create_time"`
	UpdateTime      time.Time        `json:"update_time"`
	Pickups         []PickUpResponse `json:"pickups"`

	// This variable only used for displaying to user
	TotalPickUp   int
	UpdateTimeStr string
}

type PickUpResponse struct {
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
}

type ItemResponse struct {
	Name         string    `json:"name"`
	Quantity     int64     `json:"quantity"`
	Unit         int64     `json:"unit"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
	PickUpTime   time.Time `json:"pickup_time"`
	DeadlineTime time.Time `json:"deadline"`
}
