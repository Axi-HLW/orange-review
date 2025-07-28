package main

import (
	"context"
	review "github.com/yzc/orange-review/rpc_gen/kitex_gen/review"
	"github.com/yzc/orange-review/app/review/biz/service"
)

// ReviewServiceImpl implements the last service interface defined in the IDL.
type ReviewServiceImpl struct{}

// CreateReview implements the ReviewServiceImpl interface.
func (s *ReviewServiceImpl) CreateReview(ctx context.Context, req *review.CreateReviewRequest) (resp *review.CreateReviewResponse, err error) {
	resp, err = service.NewCreateReviewService(ctx).Run(req)

	return resp, err
}

// ListReviewByStoreID implements the ReviewServiceImpl interface.
func (s *ReviewServiceImpl) ListReviewByStoreID(ctx context.Context, req *review.ListReviewByStoreIDReq) (resp *review.ListReviewByStoreIDResp, err error) {
	resp, err = service.NewListReviewByStoreIDService(ctx).Run(req)

	return resp, err
}

// ListReviewBySpuID implements the ReviewServiceImpl interface.
func (s *ReviewServiceImpl) ListReviewBySpuID(ctx context.Context, req *review.ListReviewBySpuIDReq) (resp *review.ListReviewBySpuIDResp, err error) {
	resp, err = service.NewListReviewBySpuIDService(ctx).Run(req)

	return resp, err
}

// ListReviewBySkuID implements the ReviewServiceImpl interface.
func (s *ReviewServiceImpl) ListReviewBySkuID(ctx context.Context, req *review.ListReviewBySkuIDReq) (resp *review.ListReviewBySkuIDResp, err error) {
	resp, err = service.NewListReviewBySkuIDService(ctx).Run(req)

	return resp, err
}
