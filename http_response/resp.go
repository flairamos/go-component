package http_response

// Response 返回参数
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

// NewErrResponse 失败返回
func NewErrResponse(code int, err error, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: err.Error(),
		Data:    data,
		Success: false,
	}
}

// NewSuccessResponse 成功返回
func NewSuccessResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
		Success: true,
	}
}

// NewPageErrResp 分页失败结果
func NewPageErrResp(code int, err error, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: err.Error(),
		Data:    data,
		Success: false,
	}
}

// NewPageSuccessResp 分页成功结果
func NewPageSuccessResp(code int, message string, data interface{}, page int64, perPageSize int64, total int64) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    NewPageData(page, perPageSize, total, data),
		Success: true,
	}
}
