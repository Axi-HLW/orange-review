package service

import (
	"context"
	review "github.com/yzc/orange-review/rpc_gen/kitex_gen/review"
)

type ListReviewBySpuIDService struct {
	ctx context.Context
} // NewListReviewBySpuIDService new ListReviewBySpuIDService
func NewListReviewBySpuIDService(ctx context.Context) *ListReviewBySpuIDService {
	return &ListReviewBySpuIDService{ctx: ctx}
}

// Run create note info
func (s *ListReviewBySpuIDService) Run(req *review.ListReviewBySpuIDReq) (resp *review.ListReviewBySpuIDResp, err error) {
	// Finish your business logic.

	return
}
