package library

type HError struct {
	Message string
	Code    int
}

func (e *HError) Error() string {
	return e.Message
}

// ErrorF 自定义错误格式
func ErrorF(code int, msg string) *HError {
	return &HError{
		Message: msg,
		Code:    code,
	}
}
