package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	gateway "github.com/yzc/orange-review/app/gateway/hertz_gen/gateway"
)

type AppealReviewService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAppealReviewService(Context context.Context, RequestContext *app.RequestContext) *AppealReviewService {
	return &AppealReviewService{RequestContext: RequestContext, Context: Context}
}

func (h *AppealReviewService) Run(req *gateway.AppealReviewReq) (resp *gateway.AppealReviewResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
