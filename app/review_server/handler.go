package main

import (
	"context"
	review "github.com/yzc/orange-review/app/review_server/kitex_gen/review"
	"github.com/yzc/orange-review/app/review_server/biz/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateReview implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateReview(ctx context.Context, req *review.CreateReviewRequest) (resp *review.CreateReviewResponse, err error) {
	resp, err = service.NewCreateReviewService(ctx).Run(req)

	return resp, err
}
