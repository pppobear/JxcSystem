package config

type pagination struct {
	DefaultPageSize uint
}

var p = pagination{8}

func GetPagination() *pagination {
	return &p
}
