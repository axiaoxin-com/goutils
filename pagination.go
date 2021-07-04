// 分页计算相关封装

package goutils

import (
	"math"
)

// Pagination Paginate return it
// 异常数据时分页总数为 0 ，当前页码、上下页码均不判断逻辑，只管数值增减
type Pagination struct {
	// 数据总数
	TotalCount int `json:"total_count"`
	// 分页总数
	PagesCount int `json:"pages_count"`
	// 当前页码
	PageNum int `json:"page_num"`
	// 上一页页码
	PrevPageNum int `json:"prev_page_num"`
	// 下一页页码
	NextPageNum int `json:"next_page_num"`
	// 分页大小
	PageSize int `json:"page_size"`
	// 是否有上一页
	HasPrev bool `json:"has_prev"`
	// 是否有下一页
	HasNext bool `json:"has_next"`
}

// PaginateByPageNumSize 按 pagenum,pagesize 计算分页信息
// 参数必须全部大于 0
func PaginateByPageNumSize(totalCount, pageNum, pageSize int) Pagination {
	if totalCount < 0 {
		totalCount = 0
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	pagesCount := int(math.Ceil(float64(totalCount) / float64(pageSize)))
	nextPageNum := pageNum + 1
	hasNext := nextPageNum <= pagesCount
	if nextPageNum > pagesCount {
		nextPageNum = 1
	}
	prevPageNum := pageNum - 1
	hasPrev := prevPageNum > 0
	if prevPageNum < 1 {
		prevPageNum = 1
	}
	return Pagination{
		TotalCount:  totalCount,
		PagesCount:  pagesCount,
		PageNum:     pageNum,
		PrevPageNum: prevPageNum,
		NextPageNum: nextPageNum,
		PageSize:    pageSize,
		HasPrev:     hasPrev,
		HasNext:     hasNext,
	}
}

// PaginateByOffsetLimit 按 offset,limit 计算分页信息
func PaginateByOffsetLimit(totalCount, offset, limit int) Pagination {
	pageNum := offset/limit + 1
	pageSize := limit
	return PaginateByPageNumSize(totalCount, pageNum, pageSize)
}
