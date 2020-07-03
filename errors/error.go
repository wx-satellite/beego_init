package errors

/*
 错误码为5个数字： 1 表示错误类别（1系统2普通） 00 服务模块代码  01具体的错误代码
参考：https://open.weibo.com/wiki/Error_code
*/

const (
	ServiceErrorCode  = 10000 // 服务器内部错误
	ValidateErrorCode = 20000
)

// 淘客模块
var (
	ErrorTkNotFound = &BynError{Code: 20101, Message: "淘客不存在"}
)
