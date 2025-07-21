package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	gateway "github.com/yzc/orange-review/app/gateway/hertz_gen/gateway"
)

type ListReviewByUserIDService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListReviewByUserIDService(Context context.Context, RequestContext *app.RequestContext) *ListReviewByUserIDService {
	return &ListReviewByUserIDService{RequestContext: RequestContext, Context: Context}
}

func (h *ListReviewByUserIDService) Run(req *gateway.ListReviewByUserIDReq) (resp *gateway.ListReviewByUserIDResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
