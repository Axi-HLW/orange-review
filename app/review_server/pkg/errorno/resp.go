package errorno

import (
	"errors"

	"github.com/yzc/orange-review/app/review_server/kitex_gen/base"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *base.BaseResp {
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

func baseResp(err ErrNo) *base.BaseResp {
	return &base.BaseResp{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}
