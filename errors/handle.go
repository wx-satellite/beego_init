package errors

type BynError struct {
	Code    int
	Message string
}

func (e *BynError) Error() string {
	return e.Message
}

func ParseError(err error) (code int, message string) {
	switch v := err.(type) {
	case *BynError:
		return v.Code, v.Message
	default:
		return ServiceErrorCode, "服务器开了个小差"
	}
}
