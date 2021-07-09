// 分页计算相关封装

package goutils

import (
	"math"
)

// Pagination Paginate return it
// 异常数据时分页总数为 0 ，当前页码、上下页码均不判断逻辑，只管数值增减
type Pagination struct {
	// 当前页面数据开始下标
	StartIndex int
	// 当前页面数据结束下标
	EndIndex int
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
	pagi := Pagination{}
	if totalCount == 0 {
		return pagi
	}
	pagi.TotalCount = totalCount

	pagi.PageNum = pageNum
	if pageNum <= 0 {
		pagi.PageNum = 1
	}

	// pageSize <= 0 返回全量数据
	pagi.PageSize = pageSize
	if pageSize <= 0 {
		pagi.PageSize = totalCount
	}

	pagi.PagesCount = int(math.Ceil(float64(pagi.TotalCount) / float64(pagi.PageSize)))

	pagi.NextPageNum = pagi.PageNum + 1
	pagi.HasNext = true
	// 下一页超过最大页数返回第一页
	if pagi.NextPageNum > pagi.PagesCount {
		pagi.NextPageNum = 1
		pagi.HasNext = false
	}

	// 上一页小于第一页返回最后一页
	pagi.PrevPageNum = pagi.PageNum - 1
	pagi.HasPrev = true
	if pagi.PrevPageNum < 1 {
		pagi.PrevPageNum = pagi.PagesCount
		pagi.HasPrev = false
	}

	offset := (pagi.PageNum - 1) * pagi.PageSize
	if offset < 0 {
		offset = 0
	} else if offset > pagi.TotalCount {
		offset = pagi.TotalCount
	}
	end := offset + pagi.PageSize
	if end > pagi.TotalCount {
		end = pagi.TotalCount
	}
	pagi.StartIndex = offset
	pagi.EndIndex = end
	return pagi
}

// PaginateByOffsetLimit 按 offset,limit 计算分页信息
func PaginateByOffsetLimit(totalCount, offset, limit int) Pagination {
	pageNum := offset/limit + 1
	pageSize := limit
	return PaginateByPageNumSize(totalCount, pageNum, pageSize)
}
