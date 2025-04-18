package pagination

func TotalPageFromCount(count, pageSize int) int {
	return (count + pageSize - 1) / pageSize
}
