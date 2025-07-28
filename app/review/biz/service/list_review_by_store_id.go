package service

import (
	"context"
	"errors"

	"github.com/yzc/orange-review/app/review/biz/model"
	"github.com/yzc/orange-review/app/review/pkg/errorno"
	review "github.com/yzc/orange-review/rpc_gen/kitex_gen/review"
)

type ListReviewByStoreIDService struct {
	ctx context.Context
} // NewListReviewByStoreIDService new ListReviewByStoreIDService
func NewListReviewByStoreIDService(ctx context.Context) *ListReviewByStoreIDService {
	return &ListReviewByStoreIDService{ctx: ctx}
}

// 根据StoreID查询评论列表
func (s *ListReviewByStoreIDService) Run(req *review.ListReviewByStoreIDReq) (resp *review.ListReviewByStoreIDResp, err error) {
	// Finish your business logic.
	resp = &review.ListReviewByStoreIDResp{}

	// 1. 校验参数
	if req.Page < 0 {
		resp.BaseResp = errorno.BuildServiceErrorResp(errors.New("page must be greater than 0"))
		return resp, nil
	}

	if req.PageSize <= 0 {
		resp.BaseResp = errorno.BuildServiceErrorResp(errors.New("size must be greater or equal than 0"))
		return resp, nil
	}

	// 2. 查询数据库
	reviewList, err := model.ListReviewByStoreID(s.ctx, req.StoreID, int(req.Page), int(req.PageSize))
	if err != nil {
		resp.BaseResp = errorno.BuildDBQueryErrorResp(err)
		return resp, nil
	}

	// 3. 组装数据
	for _, ri := range reviewList {
		reviewInfo := &review.ReviewInfo{
			ReviewID:     ri.ReviewID,
			UserID:       ri.UserID,
			OrderID:      ri.OrderID,
			Score:        ri.Score,
			ServiceScore: ri.ServiceScore,
			ExpressScore: ri.ExpressScore,
			Content:      ri.Content,
			PicInfo:      ri.PicInfo,
			VideoInfo:    ri.VideoInfo,
			Status:       ri.Status,
		}
		resp.List = append(resp.List, reviewInfo)
	}

	return resp, nil
}
