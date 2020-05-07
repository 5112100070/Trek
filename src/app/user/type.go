package user

type MainListAccountResponse struct {
	Data              *ListAccountsResponse `json:"data", omitempty`
	Error             *Error                `json:"error", omitempty`
	ServerProcessTime string                `json:"server_process_time"`
}

type ListAccountsResponse struct {
	Accounts       []User `json:"accounts", omitempty`
	IsHaveNextPage bool   `json:"have_next_page", omitempty`
	Total          int64  `json:"total", omitempty`
}

type User struct {
	ID         int64         `json:"user_id"`
	Fullname   string        `json:"fullname"`
	Email      string        `json:"email"`
	CreateTime string        `json:"create_time"`
	Role       int           `json:"role"`
	Attribute  UserAttribute `json:"attribute"`
}

type UserAttribute struct {
	IsEnabled bool `json:"is_enabled"`
}

type MainListCompanyResponse struct {
	Data              *ListCompanyResponse `json:"data", omitempty`
	Error             *Error               `json:"error", omitempty`
	ServerProcessTime string               `json:"server_process_time"`
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
	Code    int    `json:"code"`
	Massage string `json:"massage"`
}

type userRepo struct {
}
