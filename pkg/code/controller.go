package code

//go:generate codegen -type=Code -output register.go
//go:generate codegen -type=Code -doc -output ../../docs/error_code.md
const (
	// ErrMenuAlreadyExist - 409: 菜单已存在。
	ErrMenuAlreadyExist Code = iota + 100101
)
