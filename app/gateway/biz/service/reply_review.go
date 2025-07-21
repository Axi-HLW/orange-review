package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	gateway "github.com/yzc/orange-review/app/gateway/hertz_gen/gateway"
)

type ReplyReviewService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewReplyReviewService(Context context.Context, RequestContext *app.RequestContext) *ReplyReviewService {
	return &ReplyReviewService{RequestContext: RequestContext, Context: Context}
}

func (h *ReplyReviewService) Run(req *gateway.ReplyReviewReq) (resp *gateway.ReplyReviewResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
