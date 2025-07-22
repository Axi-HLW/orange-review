package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yzc/orange-review/app/review/infra/rpc"
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

	idResp, err := rpc.GeneratorClient.GenID(s.ctx, &generator.GenIDRequest{})
	if err != nil {
		resp.BaseResp = errorno.BuildBaseResp(err)
		return resp, nil
	}

	resp.ReviewId = idResp.Id

	return
}
