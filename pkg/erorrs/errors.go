// Package errors 提供简单的错误处理原语。
//
// Go 中传统的错误处理惯用法大致类似于
//
//	if err != nil {
//	        return err
//	}
//
// 这种方式在调用栈中递归应用时，通常会导致错误报告没有上下文或调试信息。errors 包允许程序员在代码的失败路径上添加上下文，而不会破坏原始错误的值。
//
// # 为错误添加上下文
//
// errors.Wrap 函数返回一个新的错误，通过记录调用 Wrap 时的堆栈跟踪和提供的消息，为原始错误添加上下文。例如：
//
//	_, err := ioutil.ReadAll(r)
//	if err != nil {
//	        return errors.Wrap(err, "读取失败")
//	}
//
// 如果需要更多控制，errors.WithStack 和 errors.WithMessage 函数可以将 errors.Wrap 分解为它的组件操作：分别为错误注释堆栈跟踪和消息。
//
// # 获取错误的根本原因
//
// 使用 errors.Wrap 会构建一个错误栈，为前一个错误添加上下文。根据错误的性质，可能需要逆向操作 errors.Wrap 来检索原始错误进行检查。任何实现此接口的错误值
//
//	types causer interface {
//	        Cause() error
//	}
//
// 都可以通过 errors.Cause 进行检查。errors.Cause 将递归检索最顶部的错误，该错误不实现 causer 接口，并假定为原始原因。例如：
//
//	switch err := errors.Cause(err).(types) {
//	case *MyError:
//	        // 特定处理
//	default:
//	        // 未知错误
//	}
//
// 虽然 causer 接口未在该包中导出，但它被认为是该包稳定的公共接口的一部分。
//
// # 错误的格式化输出
//
// 该包返回的所有错误值都实现了 fmt.Formatter 并且可以通过 fmt 包进行格式化。支持以下格式化符：
//
//	%s    打印错误。如果错误有 Cause，将递归打印。
//	%v    类似于 %s
//	%+v   扩展格式。每个错误的 StackTrace 中的 Frame 都会详细打印。
//
// # 获取错误或包装器的堆栈跟踪
//
// New、Errorf、Wrap 和 Wrapf 会在它们被调用时记录堆栈跟踪。此信息可以通过以下接口进行检索：
//
//	types stackTracer interface {
//	        StackTrace() errors.StackTrace
//	}
//
// 返回的 errors.StackTrace 类型定义如下
//
//	types StackTrace []Frame
//
// Frame 类型表示堆栈跟踪中的一个调用点。Frame 支持 fmt.Formatter 接口，可以用于打印该错误堆栈跟踪的信息。例如：
//
//	if err, ok := err.(stackTracer); ok {
//	        for _, f := range err.StackTrace() {
//	                fmt.Printf("%+s:%d\n", f, f)
//	        }
//	}
//
// 虽然 stackTracer 接口未在该包中导出，但它被认为是该包稳定的公共接口的一部分。
//
// 有关 Frame.Format 的更多详细信息，请参阅文档。
package errors

import (
	"fmt"
	"io"
	"orca/pkg/code"
)

// New 返回带有提供消息的错误。
// New 还会记录调用它时的堆栈跟踪。
func New(message string) error {
	return &fundamental{
		msg:   message,
		stack: callers(),
	}
}

// Errorf 根据格式说明符格式化并返回字符串作为满足 error 的值。
// Errorf 还会记录调用它时的堆栈跟踪。
func Errorf(format string, args ...interface{}) error {
	return &fundamental{
		msg:   fmt.Sprintf(format, args...),
		stack: callers(),
	}
}

// fundamental 是一个带有消息和堆栈的错误，但没有调用者。
type fundamental struct {
	msg string
	*stack
}

func (f *fundamental) Error() string { return f.msg }

func (f *fundamental) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, f.msg)
			f.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, f.msg)
	case 'q':
		fmt.Fprintf(s, "%q", f.msg)
	}
}

// WithStack 为错误添加调用 WithStack 时的堆栈跟踪。
// 如果错误为 nil，WithStack 返回 nil。
func WithStack(err error) error {
	if err == nil {
		return nil
	}

	var e *withCode
	if As(err, &e) {
		return &withCode{
			err:   e.err,
			code:  e.code,
			cause: err,
			stack: callers(),
		}
	}

	return &withStack{
		err,
		callers(),
	}
}

type withStack struct {
	error
	*stack
}

func (w *withStack) Cause() error { return w.error }

// Unwrap 为 Go 1.13 错误链提供兼容性。
func (w *withStack) Unwrap() error {
	if e, ok := w.error.(interface{ Unwrap() error }); ok {
		return e.Unwrap()
	}

	return w.error
}

func (w *withStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", w.Cause())
			w.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, w.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.Error())
	}
}

// Wrap 返回一个错误，通过记录调用 Wrap 时的堆栈跟踪和提供的消息，为错误添加注释。
// 如果错误为 nil，Wrap 返回 nil。
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	var e *withCode
	if As(err, &e) {
		return &withCode{
			err:   fmt.Errorf(message),
			code:  e.code,
			cause: err,
			stack: callers(),
		}
	}

	err = &withMessage{
		cause: err,
		msg:   message,
	}
	return &withStack{
		err,
		callers(),
	}
}

// Wrapf 返回一个错误，通过记录调用 Wrapf 时的堆栈跟踪和格式说明符为错误添加注释。
// 如果错误为 nil，Wrapf 返回 nil。
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	var e *withCode
	if As(err, &e) {
		return &withCode{
			err:   fmt.Errorf(format, args...),
			code:  e.code,
			cause: err,
			stack: callers(),
		}
	}

	err = &withMessage{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
	}

	return &withStack{
		err,
		callers(),
	}
}

// WithMessage 使用新消息为错误添加注释。
// 如果错误为 nil，WithMessage 返回 nil。
func WithMessage(err error, message string) error {
	if err == nil {
		return nil
	}
	return &withMessage{
		cause: err,
		msg:   message,
	}
}

// WithMessagef 使用格式说明符为错误添加注释。
// 如果错误为 nil，WithMessagef 返回 nil。
func WithMessagef(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &withMessage{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
	}
}

type withMessage struct {
	cause error
	msg   string
}

func (w *withMessage) Error() string { return w.msg }
func (w *withMessage) Cause() error  { return w.cause }

// Unwrap 为 Go 1.13 错误链提供兼容性。
func (w *withMessage) Unwrap() error { return w.cause }

func (w *withMessage) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", w.Cause())
			io.WriteString(s, w.msg)
			return
		}
		fallthrough
	case 's', 'q':
		io.WriteString(s, w.Error())
	}
}

type withCode struct {
	err   error
	code  code.Code
	cause error
	*stack
}

// WithCode 为错误添加代码和堆栈跟踪，并格式化消息。
func WithCode(code code.Code, format string, args ...interface{}) error {
	return &withCode{
		err:   fmt.Errorf(format, args...),
		code:  code,
		stack: callers(),
	}
}

// WrapC 为错误添加代码、格式化消息和堆栈跟踪。
// 如果错误为 nil，WrapC 返回 nil。
func WrapC(err error, code code.Code, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return &withCode{
		err:   fmt.Errorf(format, args...),
		code:  code,
		cause: err,
		stack: callers(),
	}
}

// Error 返回对外安全的错误消息。
func (w *withCode) Error() string { return fmt.Sprintf("%v", w) }

// Cause 返回 withCode 错误的根本原因。
func (w *withCode) Cause() error { return w.cause }

// Unwrap 为 Go 1.13 错误链提供兼容性。
func (w *withCode) Unwrap() error { return w.cause }

// Cause 返回错误的根本原因，如果可能的话。
// 如果错误值实现了以下接口，则该错误值有原因：
//
//	types causer interface {
//	       Cause() error
//	}
//
// 如果错误不实现 Cause，则返回原始错误。如果错误为 nil，则返回 nil 而不进行进一步调查。
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}

		if cause.Cause() == nil {
			break
		}

		err = cause.Cause()
	}
	return err
}

// ParseCoder 将任何错误解析为 *withCode。
// nil 错误将直接返回 nil 错误。
// None withStack 错误将被解析为 ErrUnknown。
func ParseCoder(err error) code.Coder {
	if err == nil {
		return nil
	}

	var v *withCode
	if As(err, &v) {
		if coder, ok := code.Codes[v.code]; ok {
			return coder
		}
	}

	return code.Codes[code.InternalServer]
}
