package tools

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Index int
	Size  int
	SortKey  string
	SortVal  string
}

func (p *Pagination) SetIndex(index int){
	p.Index = index
}

func (p *Pagination) SetSize(size int){
	p.Size = size
}

func (p *Pagination) SetSort(key,val string){
	p.SortKey = key
	p.SortVal = val
}

func NewPagination(c *gin.Context) Pagination {
	p := Pagination{}

	index,err := strconv.Atoi(c.Query("pageIndex"))
	if err != nil{
		index = 0
	}
	p.SetIndex(index)

	size ,err:= strconv.Atoi(c.Query("pageSize"))
	if err!= nil{
		size = 10
	}
	p.SetSize(size)

	sort := strings.Split(c.Query("sort"),",")
	if len(sort)==2 && len(sort[0]) > 0{
		sortVal := sort[1]
		if sortVal != "asc" && sortVal != "desc"{
			sortVal = "desc"
		}
		p.SetSort(sort[0],sortVal)
	}
	return p
}