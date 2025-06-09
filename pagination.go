// 分页计算相关封装

package goutils

import (
	"math"
)

// Pagination Paginate return it
// 异常数据时分页总数为 0 ，当前页码、上下页码均不判断逻辑，只管数值增减
type Pagination struct {
	// 当前页面数据开始下标
	StartIndex int `json:"start_index"`
	// 当前页面数据结束下标
	EndIndex int `json:"end_index"`
	// 数据总数
	TotalCount int `json:"total_count"`
	// 分页总数
	PagesCount int `json:"pages_count"`
	// 当前页码
	PageNum int `json:"page_num"`
	// 当前 offset
	PageOffset int `json:"page_offset"`
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
func PaginateByPageNumSize(totalCount, pageNum, pageSize int) Pagination {
	pagi := Pagination{}
	if totalCount == 0 {
		return pagi
	}
	pagi.TotalCount = totalCount

	// 校验 pageSize
	if pageSize <= 0 {
		pageSize = totalCount // 如果 pageSize 小于等于 0，则返回所有数据
	}
	pagi.PageSize = pageSize

	// 计算总页数
	pagi.PagesCount = int(math.Ceil(float64(pagi.TotalCount) / float64(pagi.PageSize)))

	// 校验 pageNum
	if pageNum <= 0 {
		pageNum = 1 // 如果 pageNum 小于等于 0，则返回第一页
	} else if pageNum > pagi.PagesCount {
		pageNum = pagi.PagesCount // 如果 pageNum 超过总页数，则返回最后一页
	}
	pagi.PageNum = pageNum

	// 计算 offset
	offset := (pagi.PageNum - 1) * pagi.PageSize
	if offset < 0 {
		offset = 0
	} else if offset > pagi.TotalCount {
		offset = (pagi.TotalCount - 1) / pagi.PageSize * pagi.PageSize
	}
	pagi.PageOffset = offset

	// 计算开始和结束索引
	end := offset + pagi.PageSize
	if end > pagi.TotalCount {
		end = pagi.TotalCount
	}
	pagi.StartIndex = offset
	pagi.EndIndex = end

	// 计算上一页和下一页
	pagi.NextPageNum = pagi.PageNum + 1
	pagi.HasNext = pagi.NextPageNum <= pagi.PagesCount
	if !pagi.HasNext {
		pagi.NextPageNum = 1
	}

	pagi.PrevPageNum = pagi.PageNum - 1
	pagi.HasPrev = pagi.PrevPageNum >= 1
	if !pagi.HasPrev {
		pagi.PrevPageNum = pagi.PagesCount
	}

	return pagi
}

// PaginateByOffsetLimit 按 offset,limit 计算分页信息
func PaginateByOffsetLimit(totalCount, offset, limit int) Pagination {
	// 校验 limit
	if limit <= 0 {
		limit = totalCount // 如果 limit 小于等于 0，则返回所有数据
	}
	// 如果 offset 超出总数据量，将其调整为最后一页的起始位置
	if offset >= totalCount {
		offset = (totalCount - 1) / limit * limit
	}
	pageNum := offset/limit + 1
	pageSize := limit
	return PaginateByPageNumSize(totalCount, pageNum, pageSize)
}
