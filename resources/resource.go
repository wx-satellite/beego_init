package resources

import "byn/errors"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetSuccessResponse(message string, data interface{}) Response {
	return Response{
		Code:    0,
		Message: message,
		Data:    data,
	}
}

func GetFailResponse(error error) Response {
	code, message := errors.ParseError(error)
	return Response{Code: code, Message: message}
}
