package code

import (
	"log"
	"sync"
)

type Code int

var Codes = map[Code]Coder{}
var codeMutex = &sync.Mutex{}

var AllowHttpStatus = [6]int{200, 400, 401, 403, 404, 500}

type Coder interface {
	// Code 返回这个Coder的Code值
	Code() Code

	// HttpStatus 关联的Http状态值。
	HttpStatus() int

	// Message 返回Coder关联的消息
	Message() string

	// Reference 向用户返回更加详细的文档链接
	Reference() string
}

type defaultCoder struct {
	// C 是一个整形的Code代码
	C Code

	// Http 网络请求返回的状态码
	Http int

	// Msg 返回给用户的消息
	Msg string

	// Ref 返回更加详细的文档链接
	Ref string
}

func (d defaultCoder) Code() Code {
	return d.C
}

func (d defaultCoder) HttpStatus() int {
	if d.Http == 0 {
		return 500
	}
	return d.Http
}

func (d defaultCoder) Message() string {
	return d.Msg
}

func (d defaultCoder) Reference() string {
	return d.Ref
}

// register 系统默认使用的注册函数
func register(code Code, httpStatus int, message string, refs ...string) {
	var found = false
	for _, val := range AllowHttpStatus {
		if val == httpStatus {
			found = true
		}
	}
	if !found {
		log.Panicf("为了方便处理，Orca系统只使用200, 400, 401, 403, 404, 500六种HTTP状态码\n" +
			"200：请求成功\n" +
			"400：请求存在语法错误，服务器无法理解\n" +
			"401：客户端未通过身份验证\n" +
			"403：客户端无权访问指定资源\n" +
			"404：找不到指定资源\n" +
			"500：服务器内部发生错误\n")
	}
	var reference string
	if len(refs) > 0 {
		reference = refs[0]
	}
	Register(&defaultCoder{
		C:    code,
		Http: httpStatus,
		Msg:  message,
		Ref:  reference,
	})
}

// MustRegister 注册用户自定义的错误类型，他会覆盖已有的错误类型
func MustRegister(code Coder) {
	codeMutex.Lock()
	defer codeMutex.Unlock()

	Codes[code.Code()] = code
}

// Register 注册用户自定义的错误类型，如果已经存在该类型，会触发panic
func Register(code Coder) {
	codeMutex.Lock()
	defer codeMutex.Unlock()

	if _, ok := Codes[code.Code()]; ok {
		log.Panicf("代码 `%d` 已被注册", code.Code())
	}
	Codes[code.Code()] = code
}
