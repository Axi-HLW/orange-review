package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	gateway "github.com/yzc/orange-review/app/gateway/hertz_gen/gateway"
)

type ListReviewBySpuIDService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListReviewBySpuIDService(Context context.Context, RequestContext *app.RequestContext) *ListReviewBySpuIDService {
	return &ListReviewBySpuIDService{RequestContext: RequestContext, Context: Context}
}

func (h *ListReviewBySpuIDService) Run(req *gateway.ListReviewBySpuIDReq) (resp *gateway.ListReviewBySpuIDResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
