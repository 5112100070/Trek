package global

// GenerateListPage contain list number page [1,2,3]
func GenerateListPage(rows int) []int {
	var result []int
	for i := 1; i <= rows; i++ {
		result = append(result, i)
	}

	return result
}
