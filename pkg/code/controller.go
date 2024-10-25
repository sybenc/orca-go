package code

//go:generate codegen -type=Code -output register.go
//go:generate codegen -type=Code -doc -output ../../docs/error_code.md

const (
	// ErrUserAlreadyExist - 409: 用户已存在。
	ErrUserAlreadyExist Code = iota + 100101
	// ErrUserNotFound - 404: 用户不存在。
	ErrUserNotFound
)
const (
	// ErrMenuAlreadyExist - 409: 菜单已存在。
	ErrMenuAlreadyExist Code = iota + 100201

	// ErrMenuNotFound - 404: 菜单不存在。
	ErrMenuNotFound
)
