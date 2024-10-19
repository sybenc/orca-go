package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"orca/pkg/code"
	"strings"
)

// formatInfo 包含所有的错误信息。
type formatInfo struct {
	code    code.Code
	message string
	err     string
	stack   *stack
}

// Format 实现了 fmt.Formatter 接口。https://golang.org/pkg/fmt/#hdr-Printing
//
// 支持的格式化符：
//
//	%s  - 返回映射到错误代码的用户安全错误字符串，或者如果未指定错误信息则返回错误消息。
//	%v  - %s 的别名
//
// 支持的标志：
//
//	#  - 以 JSON 格式输出，适用于日志记录
//	-  - 输出调用方的详细信息，适用于故障排查
//	+  - 输出完整的错误堆栈信息，适用于调试
//
// 示例：
//
//	%s:   内部读取 B 错误
//	%v:   内部读取 B 错误
//	%-v:  内部读取 B 错误 - #0 [/home/lk/workspace/golang/src/github.com/sybenc/main.go:12 (main.main)] (#100102) 内部服务器错误
//	%+v:  内部读取 B 错误 - #0 [/home/lk/workspace/golang/src/github.com/sybenc/main.go:12 (main.main)] (#100102) 内部服务器错误; 内部读取 A 错误 - #1 [/home/lk/workspace/golang/src/github.com/sybenc/main.go:35 (main.newErrorB)] (#100104) 验证失败
//	%#v:  [{"error":"内部读取 B 错误"}]
//	%#-v: [{"caller":"#0 /home/lk/workspace/golang/src/github.com/sybenc/main.go:12 (main.main)","error":"内部读取 B 错误","message":"(#100102) 内部服务器错误"}]
//	%#+v: [{"caller":"#0 /home/lk/workspace/golang/src/github.com/sybenc/main.go:12 (main.main)","error":"内部读取 B 错误","message":"(#100102) 内部服务器错误"},{"caller":"#1 /home/lk/workspace/golang/src/github.com/sybenc/main.go:35 (main.newErrorB)","error":"内部读取 A 错误","message":"(#100104) 验证失败"}]
func (w *withCode) Format(state fmt.State, verb rune) {
	switch verb {
	case 'v':
		str := bytes.NewBuffer([]byte{})
		var jsonData []map[string]interface{}

		var (
			flagDetail bool
			flagTrace  bool
			modeJSON   bool
		)

		if state.Flag('#') {
			modeJSON = true
		}

		if state.Flag('-') {
			flagDetail = true
		}
		if state.Flag('+') {
			flagTrace = true
		}

		sep := ""
		errs := list(w)
		length := len(errs)
		for k, e := range errs {
			finfo := buildFormatInfo(e)
			jsonData, str = format(length-k-1, jsonData, str, finfo, sep, flagDetail, flagTrace, modeJSON)
			sep = "; \n"

			if !flagTrace {
				break
			}

			if !flagDetail && !flagTrace && !modeJSON {
				break
			}
		}
		if modeJSON {
			var byts []byte
			byts, _ = json.Marshal(jsonData)

			str.Write(byts)
		}

		fmt.Fprintf(state, "%s", strings.Trim(str.String(), "\r\n\t"))
	default:
		finfo := buildFormatInfo(w)
		// Externally-safe error message
		fmt.Fprintf(state, finfo.message)
	}
}

func format(k int, jsonData []map[string]interface{}, str *bytes.Buffer, finfo *formatInfo,
	sep string, flagDetail, flagTrace, modeJSON bool) ([]map[string]interface{}, *bytes.Buffer) {
	if modeJSON {
		data := map[string]interface{}{}
		if flagDetail || flagTrace {
			data = map[string]interface{}{
				"message": finfo.message,
				"code":    finfo.code,
				"error":   finfo.err,
			}

			caller := fmt.Sprintf("#%d", k)
			if finfo.stack != nil {
				f := Frame((*finfo.stack)[0])
				caller = fmt.Sprintf("%s %s:%d (%s)",
					caller,
					f.file(),
					f.line(),
					f.name(),
				)
			}
			data["caller"] = caller
		} else {
			data["error"] = finfo.message
		}
		jsonData = append(jsonData, data)
	} else {
		if flagDetail || flagTrace {
			if finfo.stack != nil {
				f := Frame((*finfo.stack)[0])
				fmt.Fprintf(str, "%s%s - #%d [%s:%d (%s)] (%d) %s",
					sep,
					finfo.err,
					k,
					f.file(),
					f.line(),
					f.name(),
					finfo.code,
					finfo.message,
				)
			} else {
				fmt.Fprintf(str, "%s%s - #%d %s", sep, finfo.err, k, finfo.message)
			}

		} else {
			fmt.Fprintf(str, finfo.message)
		}
	}

	return jsonData, str
}

// list 将会转换错误栈为普通的数组。
func list(e error) []error {
	ret := []error{}

	if e != nil {
		if w, ok := e.(interface{ Unwrap() error }); ok {
			ret = append(ret, e)
			ret = append(ret, list(w.Unwrap())...)
		} else {
			ret = append(ret, e)
		}
	}

	return ret
}

func buildFormatInfo(e error) *formatInfo {
	var finfo *formatInfo

	switch err := e.(type) {
	case *fundamental:
		finfo = &formatInfo{
			code:    code.Codes[code.InternalServer].Code(),
			message: err.msg,
			err:     err.msg,
			stack:   err.stack,
		}
	case *withStack:
		finfo = &formatInfo{
			code:    code.Codes[code.InternalServer].Code(),
			message: err.Error(),
			err:     err.Error(),
			stack:   err.stack,
		}
	case *withCode:
		coder, ok := code.Codes[err.code]
		if !ok {
			coder = code.Codes[code.InternalServer]
		}

		extMsg := coder.Message()
		if extMsg == "" {
			extMsg = err.err.Error()
		}

		finfo = &formatInfo{
			code:    coder.Code(),
			message: extMsg,
			err:     err.err.Error(),
			stack:   err.stack,
		}
	default:
		finfo = &formatInfo{
			code:    code.Codes[code.InternalServer].Code(),
			message: err.Error(),
			err:     err.Error(),
		}
	}

	return finfo
}
