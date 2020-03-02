package entity

import "time"

type Order struct {
	ID               int64   `json:"id" db:"id"`
	InvoiceID        string  `json:"invoice_id" db:"invoice_id"`
	OrderStatus      int     `json:"order_status" db:"order_status"`
	UserID           int64   `json:"user_id" db:"user_id"`
	ContactName      string  `json:"contact_name" db:"contact_name"`
	ContactEmail     string  `json:"contact_email" db:"contact_email"`
	CompanyID        int64   `json:"company_id" db:"company_id"`
	ContactPhone     string  `json:"contact_phone" db:"contact_phone"`
	OrderPrice       float64 `json:"order_price" db:"order_price"`
	TransactionPrice float64 `json:"tx_price" db:"tx_price"`
	PickUp           string  `json:"pick_up" db:"pick_up"`
	Destination      string  `json:"destination" db:"destination"`

	PickUpTime time.Time `json:"pick_up_time" db:"pick_up_time"`
	ArriveTime time.Time `json:"arrive_time" db:"arrive_time"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}

type OrderDetailItem struct {
	ID       int64  `json:"id" db:"id"`
	OrderID  int64  `json:"order_id" db:"order_id"`
	Name     string `json:"name" db:"name"`
	Quantity int    `json:"quantity" db:"quantity"`
	Unit     int    `json:"unit" db:"unit"`
	UnitName string `json:"unit_name" db:"unit_name"`
}
