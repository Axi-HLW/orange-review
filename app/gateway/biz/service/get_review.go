package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	gateway "github.com/yzc/orange-review/app/gateway/hertz_gen/gateway"
)

type GetReviewService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetReviewService(Context context.Context, RequestContext *app.RequestContext) *GetReviewService {
	return &GetReviewService{RequestContext: RequestContext, Context: Context}
}

func (h *GetReviewService) Run(req *gateway.GetReviewReq) (resp *gateway.GetReviewResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
