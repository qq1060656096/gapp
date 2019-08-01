package util

import "math"

// ApiJsonResult 接口json返回值
type ApiJsonResult struct {
	Code string 		`json:"code"`
	Message string 		`json:"message"`
	Data interface{} 	`json:"data"`
}

type ApiPagingJsonResult struct {
	TotalCount int 		`json:"totalCount"`
	TotalPageCount int 	`json:"totalPageCount"`
	Page int 			`json:"page"`
	PageSize int 		`json:"pageSize"`
	Lists interface{} 	`json:"lists"`
}
// NewApiJsonResult 创建 ApiJsonResult
func NewApiJsonResult(code string, message string) *ApiJsonResult {
	return &ApiJsonResult{
		Code: code,
		Message: message,
	}
}

// Simple 普通数据返回
func (o *ApiJsonResult) Simple (data interface{}) *ApiJsonResult {
	o.Data = data
	return o
}

// Paging 分页
func (o *ApiJsonResult) Paging(lists interface{}, totalCount, page int, pageSize int) *ApiJsonResult {
	totalPageCount := 0
	if pageSize > 0 {
		totalPageCount = int(math.Ceil(float64(totalCount/pageSize)))
	}
	o.Data = ApiPagingJsonResult{
		TotalCount: totalCount,
		TotalPageCount: totalPageCount,
		Page: page,
		PageSize: pageSize,
		Lists: lists,
	}
	return o
}

// GetApiJsonResult 获取接口json返回值
func GetApiJsonResult(code string, message string, data interface{}) *ApiJsonResult {
	return &ApiJsonResult{
		Code: code,
		Data: data,
		Message: message,
	}
}

// GetApiJsonPagingResult 获取接口分页json返回值
func GetApiJsonPagingResult(code string, message string, lists interface{}, totalCount, page int, pageSize int) *ApiJsonResult {
	totalPageCount := 0
	if pageSize > 0 {
		totalPageCount = int(math.Ceil(float64(totalCount/pageSize)))
	}
	return &ApiJsonResult{
		Code: code,
		Data: ApiPagingJsonResult {
			TotalCount: totalCount,
			TotalPageCount: totalPageCount,
			Page: page,
			PageSize: pageSize,
			Lists: lists,
		},
		Message: message,
	}
}