package util

// ApiJsonResult 接口json返回值
type ApiJsonResult struct {
	Code string 		`json:"code"`
	Message string 		`json:"message"`
	Data interface{} 	`json:"data"`
}


// GetApiJsonResult 获取接口json返回值
func GetApiJsonResult(code string, data interface{}, message string) *ApiJsonResult {
	return &ApiJsonResult{
		Code: code,
		Data: data,
		Message: message,
	}
}
