package service

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yzc/orange-review/app/review_server/infra/rpc"
	review "github.com/yzc/orange-review/app/review_server/kitex_gen/review"
	"github.com/yzc/orange-review/app/review_server/pkg/errorno"
	"github.com/yzc/orange-review/rpc_gen/kitex_gen/generator"
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
	resp = &review.CreateReviewResponse{}
	// 1. 参数校验
	// 1.1 基本参数校验：对于商品的打分应该在0-5之间，出现负数或者超过5的情况，都应该返回参数错误（用validate）
	// if err = req.Validate(); err != nil {
	// 	resp.BaseResp = errorno.BuildBaseResp(err)
	// 	return resp, nil
	// }
	// 1.2 逻辑校验：已经评价过的订单，不应该重复评价
	// 2. 生成评价ID
	genIDResp, err := rpc.GeneratorClient.GenID(s.ctx, &generator.GenIDRequest{})
	if genIDResp.Id == -1 {
		resp.BaseResp = errorno.BuildBaseResp(errors.New("GenID service failed"))
		return resp, nil
	}
	klog.Infof("id = %d", genIDResp.Id)
	// 3. 查询订单和商品快照信息
	// 4. 评价信息入库
	return resp, nil
}
