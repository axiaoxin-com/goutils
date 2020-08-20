// 分页计算相关封装

package goutils

import "math"

// Pagination Paginate return it
// 异常数据时分页总数为 0 ，当前页码、上下页码均不判断逻辑，只管数值增减
type Pagination struct {
	// 数据总数
	TotalCount int `json:"total_count"`
	// 分页总数
	PagesCount int `json:"pages_count"`
	// 当前页码
	PageNum int `json:"page_num"`
	// 分页大小
	PageSize int `json:"page_size"`
	// 是否有上一页
	HasPrev bool `json:"has_prev"`
	// 是否有下一页
	HasNext bool `json:"has_next"`
	// 上一页页码
	PrevPageNum int `json:"prev_page_num"`
	// 下一页页码
	NextPageNum int `json:"next_page_num"`
}

// PaginateByPageNumSize 按 pagenum,pagesize 计算分页信息
func PaginateByPageNumSize(totalCount, pageNum, pageSize int) Pagination {
	if totalCount <= 0 {
		return Pagination{}
	}
	pagesCount := 0
	if pageSize > 0 {
		pagesCount = int(math.Ceil(float64(totalCount) / float64(pageSize)))
	}
	hasNext := true
	nextPageNum := pageNum + 1
	if nextPageNum >= pagesCount {
		hasNext = false
	}
	hasPrev := true
	prevPageNum := pageNum - 1
	if prevPageNum <= 0 {
		hasPrev = false
	}
	return Pagination{
		TotalCount:  totalCount,
		PagesCount:  pagesCount,
		PageNum:     pageNum,
		PageSize:    pageSize,
		HasPrev:     hasPrev,
		HasNext:     hasNext,
		PrevPageNum: prevPageNum,
		NextPageNum: nextPageNum,
	}
}

// PaginateByOffsetLimit 按 offset,limit 计算分页信息
func PaginateByOffsetLimit(totalCount, offset, limit int) Pagination {
	if totalCount <= 0 {
		return Pagination{}
	}
	pageNum := 1
	if offset <= 0 {
		pageNum = 1
	} else {
		pageNum = offset/limit + 1
	}
	pagesCount := 0
	if limit > 0 {
		pagesCount = int(math.Ceil(float64(totalCount) / float64(limit)))
	}
	hasNext := true
	nextPageNum := pageNum + 1
	if limit == 0 || offset+limit >= totalCount || nextPageNum >= pagesCount {
		hasNext = false
	}
	hasPrev := true
	prevPageNum := pageNum - 1
	if limit == 0 || offset+limit <= 0 || pageNum == 1 {
		hasPrev = false
	}
	return Pagination{
		TotalCount:  totalCount,
		PagesCount:  pagesCount,
		PageNum:     pageNum,
		PageSize:    limit,
		HasPrev:     hasPrev,
		HasNext:     hasNext,
		PrevPageNum: prevPageNum,
		NextPageNum: nextPageNum,
	}
}
