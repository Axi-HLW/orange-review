package errorno

import (
	"errors"
	"fmt"
)

const (
	// Review Service 20 xxx xx
	// Success xx 000 00
	ReviewServiceSuccessCode = 2000000

	// Param Error xx 001 xx
	UnknownFieldErrCode = 2000101
	MissingParamErrCode = 2000102
	InvalidParamErrCode = 2000103

	// Service Error xx 002 xx
	CommonServiceErrCode = 2000201

	// Authorization Error xx 003 xx

	// Business Error xx 004 xx

	// Database Error xx 005 xx
	DBQueryFailedCode = 2000501
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}

// WithMessage 给ErrNo对象的ErrMsg字段赋值
func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	ReviewServiceSuccess = NewErrNo(ReviewServiceSuccessCode, "Success")
	UnknownFieldErr      = NewErrNo(UnknownFieldErrCode, "Unknown field")
	MissingParamErr      = NewErrNo(MissingParamErrCode, "Missing parameter")
	InvalidParamErr      = NewErrNo(InvalidParamErrCode, "Invalid parameter")
	DBQueryFailedErr     = NewErrNo(DBQueryFailedCode, "DB query failed")
	CommonServiceErr     = NewErrNo(CommonServiceErrCode, "Common service error")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := CommonServiceErr
	s.ErrMsg = err.Error()
	return s
}
