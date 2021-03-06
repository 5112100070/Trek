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
	// Using Public Endpoint Version
	IsInternal bool
}

// DetailCompanyParam This struct is used as parameter input for method GetDetailCompany
type DetailCompanyParam struct {
	CompanyID int64
	// Using Public Endpoint Version
	IsInternal bool
}

// CreateAccountParam This struct is used as parameter to create new account
type CreateAccountParam struct {
	Fullname     string `json:"fullname"`
	Email        string `json:"email"`
	Phone        string `json:"phone_number"`
	Role         int    `json:"role"`
	CompanyID    int64  `json:"company_id"`
	ProfileImage string `json:"profile_image"`
}

// UpdateAccountParam This struct is used as parameter to update account
type UpdateAccountParam struct {
	ID           int64  `json:"id"`
	Fullname     string `json:"fullname"`
	Phone        string `json:"phone_number"`
	Role         int    `json:"role"`
	ProfileImage string `json:"profile_image"`
}

// CreateCompanyParam This struct is used as parameter to create new company
type CreateCompanyParam struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Phone     string `json:"phone_number"`
	Role      int    `json:"role"`
	ImageLogo string `json:"image_logo"`
	ParentID  int64  `json:"parent_id"`
}

// UpdateCompanyParam This struct is used as parameter to update company
type UpdateCompanyParam struct {
	ID        int64  `json:"id"`
	Name      string `json:"company_name"`
	Address   string `json:"address"`
	Phone     string `json:"phone_number"`
	Role      int    `json:"role"`
	ImageLogo string `json:"image_logo"`
}

// ChangePasswordParam This struct is used as parameter to change password user accounts
type ChangePasswordParam struct {
	UserID      int64  `json:"user_id"`
	NewPassword string `json:"new_password"`
}

// ChangeStatusAccParam This struct is used as parameter to change status user accounts
type ChangeStatusAccParam struct {
	UserID    int64 `json:"user_id"`
	CompanyID int64 `json:"company_id"`
	IsEnabled bool  `json:"is_enabled"`
}
