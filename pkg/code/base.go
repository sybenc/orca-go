package code

//go:generate codegen -type=Code -output register.go
//go:generate codegen -type=Code -doc -output ../../docs/error_code.md
const (
	// Success - 200: 请求成功。
	Success Code = iota + 100001

	// ErrInternalServer - 500: 服务器内部错误。
	ErrInternalServer

	// ErrBadRequest - 400: 请求存在错误。
	ErrBadRequest

	// ErrNotFound - 404: 资源未找到。
	ErrNotFound

	// ErrValidate - 400: 字段验证错误。
	ErrValidate

	// ErrBind - 400: 参数绑定错误。
	ErrBind
)
