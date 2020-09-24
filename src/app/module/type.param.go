package module

// ListModuleParam This struct is used as parameter input for method GetListModule
type ListModuleParam struct {
	Page      int
	Rows      int
	OrderType string
}

// ListFeatureParam This struct is used as parameter input for method GetListFeature
type ListFeatureParam struct {
	Page      int
	Rows      int
	OrderType string
}
