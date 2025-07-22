package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	gateway "github.com/yzc/orange-review/app/gateway/hertz_gen/gateway"
	"github.com/yzc/orange-review/app/gateway/infra/rpc"
	"github.com/yzc/orange-review/rpc_gen/kitex_gen/review"
)

type CreateReviewService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateReviewService(Context context.Context, RequestContext *app.RequestContext) *CreateReviewService {
	return &CreateReviewService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateReviewService) Run(req *gateway.CreateReviewReq) (resp *gateway.CreateReviewResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code

	request := &review.CreateReviewRequest{
		UserId:       req.UserID,
		OrderId:      req.OrderID,
		ItemScore:    req.Score,
		ServiceScore: req.ServiceScore,
		ExpressScore: req.ExpressScore,
		Content:      req.Content,
		PicInfo:      req.PicInfo,
		VideoInfo:    req.VideoInfo,
		Anonymous:    req.Anonymous,
	}
	rpcResp, err := rpc.ReviewClient.CreateReview(h.Context, request)
	resp = &gateway.CreateReviewResp{
		ReviewID: rpcResp.ReviewId,
	}
	return resp, err
}
