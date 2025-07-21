package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	gateway "github.com/yzc/orange-review/app/gateway/hertz_gen/gateway"
)

type AuditAppealService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAuditAppealService(Context context.Context, RequestContext *app.RequestContext) *AuditAppealService {
	return &AuditAppealService{RequestContext: RequestContext, Context: Context}
}

func (h *AuditAppealService) Run(req *gateway.AuditAppealReq) (resp *gateway.AuditAppealResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
