package is

import (
	"orca/pkg/validation"
	"regexp"
	"unicode"
)

var (
	// ErrUsername is the error that returns in case of an invalid username
	ErrUsername = validation.NewError("validation_is_username", "must be a valid username")
	// ErrPassword is the error that returns in case of an invalid password
	ErrPassword = validation.NewError("validation_is_password", "must be a valid password")
	// ErrCode is used to handle cases where role codes or dictionary type codes do not conform to the format.
	ErrCode = validation.NewError("validation_is_code", "must be a valid code")
	// ErrPath is the error that returns in case of an invalid resource path
	ErrPath = validation.NewError("validation_is_path", "must be a valid resource path")
)

var (
	// reSafeCharacterSet Considered a safe character set.
	reSafeCharacterSet = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	// rePath Resource path.
	rePath = regexp.MustCompile(`^/[a-zA-Z0-9\-/_]+(\?[a-zA-Z0-9=&]*)?$`)
)

func isSafeString(value string) bool {
	return reSafeCharacterSet.MatchString(value)
}

func isPassword(value string) bool {
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	for _, char := range value {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	// 至少需要两种不同的字符类型
	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasDigit {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 2
}

func isPath(value string) bool {
	return rePath.MatchString(value)
}

var Username = validation.NewStringRuleWithError(isSafeString, ErrUsername)
var Password = validation.NewStringRuleWithError(isPassword, ErrPassword)
var Code = validation.NewStringRuleWithError(isSafeString, ErrCode)
var Path = validation.NewStringRuleWithError(isPath, ErrPath)
