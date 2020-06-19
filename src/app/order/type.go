package order

type orderRepo struct {
}

type CreateOrderForAdminResponse struct {
	Error             *string `json:"error", omitempty`
	Status            string  `json:"status"`
	ServerProcessTime string  `json:"server_process_time"`
}
