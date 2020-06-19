package order

// CreateOrderParam This struct is used as parameter input for method Create Order
type CreateOrderParam struct {
	Receiver struct {
		CreateOrderDefaultData
	} `json:"receiver"`
	Pickups []struct {
		PIC string
		CreateOrderDefaultData
		Items []struct {
			Name     string  `json:"name"`
			Quantity float64 `json:"quantity"`
			Unit     int     `json:"unit"`
			Notes    string  `json:"notes"`
		} `json:"items"`
	} `json:"pickups"`
}

// CreateOrderDefaultData This struct is standart detail pick up data
type CreateOrderDefaultData struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Notes       string `json:"notes"`
}
