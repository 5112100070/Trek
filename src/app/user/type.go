package user

type MainListAccountResponse struct {
	Data              *ListAccountsResponse `json:"data", omitempty`
	Error             *Error                `json:"error", omitempty`
	ServerProcessTime string                `json:"server_process_time"`
}

type MainDetailAccountResponse struct {
	Data              User   `json:"data", omitempty`
	Error             *Error `json:"error", omitempty`
	ServerProcessTime string `json:"server_process_time"`
}

type ListAccountsResponse struct {
	Accounts       []User `json:"accounts", omitempty`
	IsHaveNextPage bool   `json:"have_next_page", omitempty`
	Total          int    `json:"total", omitempty`
}

type User struct {
	ID          int64         `json:"user_id"`
	Fullname    string        `json:"fullname"`
	CompanyID   int64         `json:"company_id"`
	Company     UserCompany   `json:"company"`
	PhoneNumber string        `json:"phone_number"`
	Email       string        `json:"email"`
	CreateTime  string        `json:"create_time"`
	Role        int           `json:"role"`
	RoleWording string        `json:"role_wording"`
	RoleColor   string        `json:"role_color"`
	Attribute   UserAttribute `json:"attribute"`
}

type UserCompany struct {
	ID          int64  `json:"company_id"`
	CompanyName string `json:"company_name"`
}

type UserAttribute struct {
	IsEnabled bool `json:"is_enabled"`
}

type MainListCompanyResponse struct {
	Data              *ListCompanyResponse `json:"data", omitempty`
	Error             *Error               `json:"error", omitempty`
	ServerProcessTime string               `json:"server_process_time"`
}

type MainDetailCompanyResponse struct {
	Data              CompanyProfile `json:"data", omitempty`
	Error             *Error         `json:"error", omitempty`
	ServerProcessTime string         `json:"server_process_time"`
}

type ListCompanyResponse struct {
	Companies      []CompanyProfile `json:"companies", omitempty`
	IsHaveNextPage bool             `json:"have_next_page", omitempty`
	Total          int              `json:"total", omitempty`
}

type CompanyProfile struct {
	ID               int64  `json:"company_id"`
	CompanyName      string `json:"company_name"`
	Address          string `json:"address"`
	PhoneNumber      string `json:"phone_number"`
	IsEnabled        bool   `json:"is_enabled"`
	StatusActivation string `json:"status_activation"`
	Role             int    `json:"role"`
}

type Error struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

type userRepo struct {
}
