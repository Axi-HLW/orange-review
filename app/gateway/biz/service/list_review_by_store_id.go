package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	gateway "github.com/yzc/orange-review/app/gateway/hertz_gen/gateway"
	"github.com/yzc/orange-review/app/gateway/infra/rpc"
	"github.com/yzc/orange-review/rpc_gen/kitex_gen/review"
)

type ListReviewByStoreIDService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListReviewByStoreIDService(Context context.Context, RequestContext *app.RequestContext) *ListReviewByStoreIDService {
	return &ListReviewByStoreIDService{RequestContext: RequestContext, Context: Context}
}

func (h *ListReviewByStoreIDService) Run(req *gateway.ListReviewByStoreIDReq) (resp *gateway.ListReviewByStoreIDResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code

	request := &review.ListReviewByStoreIDReq{
		StoreID:  req.StoreID,
		Page:     req.Page,
		PageSize: req.Size,
	}

	rpcResp, err := rpc.ReviewClient.ListReviewByStoreID(h.Context, request)

	reviewInfoList := make([]*gateway.ReviewInfo, 0)
	for _, ri := range rpcResp.List {
		reviewInfoList = append(reviewInfoList, &gateway.ReviewInfo{
			ReviewID:     ri.ReviewID,
			UserID:       ri.UserID,
			OrderID:      ri.OrderID,
			Score:        ri.Score,
			ServiceScore: ri.ServiceScore,
			ExpressScore: ri.ExpressScore,
			Content:      ri.Content,
			PicInfo:      ri.PicInfo,
			VideoInfo:    ri.VideoInfo,
		})
	}

	resp = &gateway.ListReviewByStoreIDResp{
		List: reviewInfoList,
	}
	return resp, err
}
