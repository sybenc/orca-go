package code

//go:generate codegen -type=Code -output register.go
//go:generate codegen -type=Code -doc -output ../../docs/error_code.md
const (
	// ErrMenuAlreadyExist - 409: 菜单已存在。
	ErrMenuAlreadyExist Code = iota + 100101

	// ErrMenuNotFound - 404: 菜单未找到。
	ErrMenuNotFound Code = iota + 100101
)
