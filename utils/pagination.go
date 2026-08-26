package utils

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Page   int
	Limit  int
	Offset int
}

func GetPagination(request *http.Request) Pagination {
	page, err := strconv.Atoi(request.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(request.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit
	return Pagination{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}

}
