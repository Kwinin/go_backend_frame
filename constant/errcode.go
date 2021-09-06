package constant

import (
	"fmt"
	"strings"
)

const (
	UnknownError = -1
	Success      = 0
	Failure      = 1

	// DB
	DBError = 1000
	DataErr = 1001

	//login
	AnotherClientLoginIn = 2000
	StatusUserLoginError = 2001
	UserAccountForbid    = 2002
	UserLoginError       = 2003
	//param
	ParamError     = 4000
	ParamTypeError = 4001
	ParamJsonError = 4002
	//resource
	FileUploadError = 5001
	FileTypeError   = 5002
	FileStatusError = 5003
	// permission
	NoPermission = 6001

	//user
	CanNotFindAuth     = 10001
	UserForbiddenError = 10002
)

type InitError struct {
	Err  error
	Code int
	Desc string
}

func (i InitError) Error() string {
	str := fmt.Sprintf("initError: {code: %d, desc: %s, err: %s}", i.Code, i.Desc, i.Err)
	return str
}

func TranslateErrCode(code int, extra ...string) string {
	var msg string
	switch code {
	case UnknownError:
		msg = "Unknown error?"
	case Success:
		msg = "success!"
	case Failure:
		msg = "failure!"
	case ParamError:
		msg = "param error!"
	case ParamTypeError:
		msg = "param type error!"
	case CanNotFindAuth:
		msg = "user not found!"
	case DBError:
		msg = "user not found!"
	case AnotherClientLoginIn:
		msg = "your account login from another client"
	case StatusUserLoginError:
		msg = "user do not exit or password wrong!"
	case UserAccountForbid:
		msg = "User account is disabled"
	case UserForbiddenError:
		msg = "account forbidden,please contact admin!"
	case DataErr:
		msg = "data err!"
	case UserLoginError:
		msg = "user login error"
	case FileUploadError:
		msg = "file upload error"
	case FileTypeError:
		msg = "file type error"
	case FileStatusError:
		msg = "Resource status is not allowed to be used"
	case ParamJsonError:
		msg = "param json type error"
	case NoPermission:
		msg = "Permission denied"
	default:
	}

	if len(extra) > 0 {
		msg = msg + ": " + strings.Join(extra, ",")
	}
	return msg
}
