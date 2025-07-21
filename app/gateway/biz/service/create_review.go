package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	gateway "github.com/yzc/orange-review/app/gateway/hertz_gen/gateway"
)

type CreateReviewService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateReviewService(Context context.Context, RequestContext *app.RequestContext) *CreateReviewService {
	return &CreateReviewService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateReviewService) Run(req *gateway.CreateReviewReq) (resp *gateway.CreateReviewResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
