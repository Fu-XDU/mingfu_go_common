package base_response

var (
	SUCCESS      = NewRetCode(0, "Success")
	UnknownError = NewRetCode(10000, "Unknown error")
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func NewResponse() *Response {
	return NewDataResponse(nil)
}

func NewDataResponse(data interface{}) *Response {
	return &Response{
		Code:    SUCCESS.Code,
		Message: SUCCESS.Message,
		Data:    data,
	}
}

func NewErrorResponse(err error, retcode *RetCode) *Response {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	if retcode == nil {
		retcode = UnknownError
	}

	return &Response{
		Code:    retcode.Code,
		Message: retcode.Message,
		Error:   errMsg,
	}
}
