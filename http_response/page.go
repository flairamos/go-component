package http_response

import "errors"

type Page struct {
	Data    interface{} `json:"data,omitempty"`
	Total   int64       `json:"total,omitempty"`
	HasNext bool        `json:"has_next,omitempty"`
	HasPrev bool        `json:"has_prev,omitempty"`
}

// CalculatePage  分页计算,根据currentPage 和 pageSize 计算出offset与limit
// 数据从数据库获取，如果data仍是经计算分页查询的的数据则该方案无意义
// 该方法结合分级计算，简化类offset limit，可以减少一次数据库查询
// 传入的数据只包含WHERE条件即可
// 不分页不可使用该方法
func CalculatePage(page int64, perPageSize int64, total int64) (offset, limit int64, err error) {
	if total <= 0 || perPageSize <= 0 || total < perPageSize {
		return 0, 0, errors.New("page number invalid")
	}
	if page <= 1 {
		// 第一页
		offset = 0
		limit = perPageSize - 1
	}
	offset = (page - 1) * perPageSize
	limit = page*perPageSize - 1
	return
}

func NewPageData(page int64, perPageSize int64, total int64, data interface{}) *Page {
	return &Page{
		Data:    data,
		Total:   total,
		HasNext: page < total/perPageSize,
		HasPrev: page > 1,
	}
}
