package errs

import "fmt"

func GetErrMsg(code ErrCode) string {
	if msg, ok := errMsgMap[code]; ok {
		return msg
	}
	return fmt.Sprintf("unknown error code: %v", code)
}
