package base_response

type PageResponse struct {
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	Error       string      `json:"error"`
	Data        interface{} `json:"data"`
	CurrentPage uint        `json:"current"`
	PageSize    uint        `json:"pageSize"`
	Total       uint64      `json:"total"`
}

func NewPageResponse(currentPage, pageSize uint, total uint64) *PageResponse {
	return NewDataPageResponse(nil, currentPage, pageSize, total)
}

func NewDataPageResponse(data interface{}, currentPage, pageSize uint, total uint64) *PageResponse {
	return &PageResponse{
		Code:        SUCCESS.Code,
		Message:     SUCCESS.Message,
		Data:        data,
		CurrentPage: currentPage,
		PageSize:    pageSize,
		Total:       total,
	}
}

func NewErrorPageResponse(err error, retCode *RetCode, currentPage, pageSize uint, total uint64) *PageResponse {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	return &PageResponse{
		Code:        retCode.Code,
		Message:     retCode.Message,
		Error:       errMsg,
		CurrentPage: currentPage,
		PageSize:    pageSize,
		Total:       total,
	}
}
