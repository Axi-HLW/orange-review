package service

import (
	"context"
	review "github.com/yzc/orange-review/rpc_gen/kitex_gen/review"
)

type ListReviewBySkuIDService struct {
	ctx context.Context
} // NewListReviewBySkuIDService new ListReviewBySkuIDService
func NewListReviewBySkuIDService(ctx context.Context) *ListReviewBySkuIDService {
	return &ListReviewBySkuIDService{ctx: ctx}
}

// Run create note info
func (s *ListReviewBySkuIDService) Run(req *review.ListReviewBySkuIDReq) (resp *review.ListReviewBySkuIDResp, err error) {
	// Finish your business logic.

	return
}
