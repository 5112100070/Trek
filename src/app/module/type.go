package module

type ModuleResponse struct {
	Data *struct {
		Modules []Module `json:"modules"`
		Total   int      `json:"total"`
	} `json:"data", omitempty`
	Error             *Error `json:"error", omitempty`
	ServerProcessTime string `json:"server_process_time"`
}

type Module struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Error struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

type moduleRepo struct {
}
