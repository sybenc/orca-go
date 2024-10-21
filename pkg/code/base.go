package code

const (
	// Success - 200: 请求成功。
	Success Code = iota + 100001

	// ErrInternalServer - 500: 服务器内部错误。
	ErrInternalServer

	// ErrBadRequest - 400: 请求存在错误。
	ErrBadRequest

	// ErrNotFound - 404: 资源未找到。
	ErrNotFound

	// ErrValidation - 400: 字段验证错误。
	ErrValidation

	// ErrBind - 400: 参数绑定错误。
	ErrBind
)
