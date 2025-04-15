package pagination

func GetSQLArgs(page, pageSize int) (offset, limit int) {
	offset = (page-1)*pageSize + 1
	limit = pageSize
	return
}

