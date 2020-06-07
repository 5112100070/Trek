package user

// ListUserParam This struct is used as parameter input for method GetListUsers
type ListUserParam struct {
	// CompanyID used as filter based on company
	CompanyID int64
	Page      int
	Rows      int
	OrderType string
}

// ListCompanyParam This struct is used as parameter input for method GetListCompany
type ListCompanyParam struct {
	Page             int
	Rows             int
	OrderType        string
	FilterByIsEnable string
}

// CreateAccountParam This struct is used as parameter to create new account
type CreateAccountParam struct {
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	Phone     string `json:"phone_number"`
	Role      int    `json:"role"`
	CompanyID int64  `json:"company_id"`
}

// UpdateAccountParam This struct is used as parameter to update account
type UpdateAccountParam struct {
	ID       int64  `json:"id"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone_number"`
	Role     int    `json:"role"`
}

// CreateCompanyParam This struct is used as parameter to create new company
type CreateCompanyParam struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone_number"`
	Role    int    `json:"role"`
}

// UpdateCompanyParam This struct is used as parameter to update company
type UpdateCompanyParam struct {
	ID      int64  `json:"id"`
	Name    string `json:"company_name"`
	Address string `json:"address"`
	Phone   string `json:"phone_number"`
	Role    int    `json:"role"`
}

// ChangePasswordParam This struct is used as parameter to change password user accounts
type ChangePasswordParam struct {
	UserID      int64  `json:"user_id"`
	NewPassword string `json:"new_password"`
}
