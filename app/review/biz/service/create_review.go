package service

import (
	"context"
	"errors"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yzc/orange-review/app/review/biz/dal/mysql"
	"github.com/yzc/orange-review/app/review/biz/model"
	"github.com/yzc/orange-review/app/review/infra/rpc"
	"github.com/yzc/orange-review/app/review/pkg/consts"
	"github.com/yzc/orange-review/app/review/pkg/errorno"
	"github.com/yzc/orange-review/rpc_gen/kitex_gen/generator"
	review "github.com/yzc/orange-review/rpc_gen/kitex_gen/review"
)

type CreateReviewService struct {
	ctx context.Context
} // NewCreateReviewService new CreateReviewService
func NewCreateReviewService(ctx context.Context) *CreateReviewService {
	return &CreateReviewService{ctx: ctx}
}

// Run create note info
func (s *CreateReviewService) Run(req *review.CreateReviewRequest) (resp *review.CreateReviewResponse, err error) {
	// Finish your business logic.
	klog.Info("CreateReview ... ")

	resp = &review.CreateReviewResponse{}

	reviews, err := model.GetReviewByOrderId(mysql.DB, s.ctx, req.OrderId)
	if err != nil {
		resp.BaseResp = errorno.BuildDBQueryErrorResp(err)
		return resp, nil
	}
	if len(*reviews) > 0 {
		resp.BaseResp = errorno.BuildServiceErrorResp(errors.New("订单已评价"))
		return resp, nil
	}

	idResp, err := rpc.GeneratorClient.GenID(s.ctx, &generator.GenIDRequest{})
	if err != nil {
		resp.BaseResp = errorno.BuildServiceErrorResp(err)
		return resp, nil
	}

	resp.ReviewId = idResp.Id

	currentTime := time.Now()
	anonymous := 0
	if req.Anonymous {
		anonymous = 1
	}

	review := &model.ReviewInfo{
		CreateAt:     currentTime,
		UpdateAt:     currentTime,
		ReviewID:     idResp.Id,
		OrderID:      req.OrderId,
		UserID:       req.UserId,
		SkuID:        req.SkuId,
		SpuID:        req.SpuId,
		StoreID:      req.StoreId,
		Content:      req.Content,
		Score:        req.ItemScore,
		ServiceScore: req.ServiceScore,
		ExpressScore: req.ExpressScore,
		Anonymous:    int32(anonymous),
		PicInfo:      req.PicInfo,
		VideoInfo:    req.VideoInfo,
		Status:       consts.ReviewStatusInit,
	}

	_, err = model.CreateReview(mysql.DB, s.ctx, review)
	if err != nil {
		resp.BaseResp = errorno.BuildDBQueryErrorResp(err)
		return resp, nil
	}

	return
}
