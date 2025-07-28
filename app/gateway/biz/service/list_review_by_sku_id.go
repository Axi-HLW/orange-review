package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	gateway "github.com/yzc/orange-review/app/gateway/hertz_gen/gateway"
)

type ListReviewBySkuIDService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListReviewBySkuIDService(Context context.Context, RequestContext *app.RequestContext) *ListReviewBySkuIDService {
	return &ListReviewBySkuIDService{RequestContext: RequestContext, Context: Context}
}

func (h *ListReviewBySkuIDService) Run(req *gateway.ListReviewBySkuIDReq) (resp *gateway.ListReviewBySkuIDResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
