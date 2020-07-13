package entity

// Pagination is default struct response for pagination configuration
type Pagination struct {
	Template  string
	Page      int
	NextPage  int
	PrevPage  int
	Rows      int
	TotalPage int
	// ListPage contain list number page ["1","2","3"]
	ListPage []int
}
