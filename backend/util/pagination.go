package util

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Current int
	Size    int
	SortKey string
	SortVal string
}

func NewPagination(c *gin.Context) Pagination {
	page := Pagination{}

	current, err := strconv.Atoi(c.Query("current"))
	if err != nil {
		current = 1
	}
	page.Current = current

	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		size = 10
	}
	page.Size = size

	sort := strings.Split(c.Query("sort"), ",")
	if len(sort) == 2 && len(sort[0]) > 0 {
		sortVal := sort[1]
		if sortVal != "asc" && sortVal != "desc" {
			sortVal = "desc"
		}
		page.SortKey = sort[0]
		page.SortVal = sortVal
	}
	return page
}
