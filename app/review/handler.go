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
