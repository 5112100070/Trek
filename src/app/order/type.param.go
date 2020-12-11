package order

// CreateOrderParam This struct is used as parameter input for method Create Order
type CreateOrderParam struct {
	Receiver struct {
		OrderDefaultData
	} `json:"receiver"`
	Pickups []struct {
		PIC string
		OrderDefaultData
		Items []struct {
			Name     string  `json:"name"`
			Quantity float64 `json:"quantity"`
			Unit     int     `json:"unit"`
			Notes    string  `json:"notes"`
		} `json:"items"`
	} `json:"pickups"`
	Airwaybill   string `json:"airwaybill"`
	DeliveryType int    `json:"delivery_type"`
	CompanyID    int64  `json:"company_id"`
	Deadline     string `json:"deadline"`
}

// OrderDefaultData This struct is standart detail pick up data
type OrderDefaultData struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Notes       string `json:"notes"`
	Kecamatan   string `json:"kecamatan"`
	Kelurahan   string `json:"kelurahan"`
	Kota        string `json:"kota"`
	Provinsi    string `json:"provinsi"`
	ZIP         int    `json:"zip"`
}

// ListOrderParam This struct is used as parameter input for method GetListOrder
type ListOrderParam struct {
	Page      int
	Rows      int
	OrderType string
	// CompanyID used as filter based on company
	CompanyID int64
	// Status used as filter based on status
	Status int
	// CreateTimeFrom used as filter based on create time order
	CreateTimeFrom string
	CreateTimeTo   string
}
