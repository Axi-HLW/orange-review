package errorno

import (
	"errors"

	"github.com/yzc/orange-review/rpc_gen/kitex_gen/base"
)

// BuildServiceErrorResp 构建统一错误响应
func BuildServiceErrorResp(err error) *base.BaseResp {
	if err == nil {
		return baseResp(ReviewServiceSuccess)
	}

	e := ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := CommonServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

// BuildDBQueryErrorResp 构建数据库查询错误响应
func BuildDBQueryErrorResp(err error) *base.BaseResp {
	if err == nil {
		return baseResp(ReviewServiceSuccess)
	}

	e := ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := DBQueryFailedErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err ErrNo) *base.BaseResp {
	return &base.BaseResp{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}
