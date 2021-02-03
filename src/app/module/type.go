package module

type ModuleResponse struct {
	Data *struct {
		Modules []Module `json:"modules"`
		Total   int      `json:"total"`
	} `json:"data", omitempty`
	Error             *Error `json:"error", omitempty`
	ServerProcessTime string `json:"server_process_time"`
}

type FeatureResponse struct {
	Data *struct {
		Features []Features `json:"features"`
		Total    int        `json:"total"`
	} `json:"data", omitempty`
	Error             *Error `json:"error", omitempty`
	ServerProcessTime string `json:"server_process_time"`
}

type Module struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Features struct {
	ID              int64  `json:"id"`
	AccountModuleID int64  `json:"account_module_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	PathUrl         string `json:"path_url"`
	Method          string `json:"method"`
}

type Error struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

type moduleRepo struct {
}
