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
	Page      int
	Rows      int
	OrderType string
}
