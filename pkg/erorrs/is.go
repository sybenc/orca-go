package errors

import (
	stderrors "errors"
	"orca/pkg/code"
)

// Is 报告 err 的链条中是否有任何错误与目标 target 匹配。
// 错误链包括 err 本身，后续错误是通过反复调用 Unwrap 获得的。
// 如果错误与目标相等，或实现了 Is(error) bool 方法且 Is(target) 返回 true，则认为该错误与目标匹配。
func Is(err, target error) bool { return stderrors.Is(err, target) }

// As 查找 err 的链条中第一个与目标 target 匹配的错误，如果找到，则将 target 设置为该错误值并返回 true。
// 错误链包括 err 本身，后续错误是通过反复调用 Unwrap 获得的。 如果错误的具体值可以赋给 target 指向的值，
// 或者错误有一个 As(interface{}) bool 方法且 As(target) 返回 true，则错误与目标匹配。
// 在后一种情况下，As 方法负责设置 target。 如果 target 不是一个非 nil 指针，指向实现 error 接口的类型或任意接口类型，As 将会引发 panic。如果 err 为 nil，As 返回 false。
func As(err error, target interface{}) bool { return stderrors.As(err, &target) }

// Unwrap 返回调用 err 的 Unwrap 方法的结果，如果 err 类型包含一个返回 error 的 Unwrap 方法。
// 否则，Unwrap 返回 nil。
func Unwrap(err error) error {
	return stderrors.Unwrap(err)
}

// IsCode 判断 err 链中的任何错误是否包含给定的错误代码。
func IsCode(err error, code code.Code) bool {
	var v *withCode
	if As(err, &v) {
		if v.code == code {
			return true
		}

		if v.cause != nil {
			return IsCode(v.cause, code)
		}

		return false
	}

	return false
}
