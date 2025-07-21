package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	gateway "github.com/yzc/orange-review/app/gateway/hertz_gen/gateway"
)

type AuditReviewService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAuditReviewService(Context context.Context, RequestContext *app.RequestContext) *AuditReviewService {
	return &AuditReviewService{RequestContext: RequestContext, Context: Context}
}

func (h *AuditReviewService) Run(req *gateway.AuditReviewReq) (resp *gateway.AuditReviewResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
