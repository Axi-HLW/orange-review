package review

import (
	"context"
	review "github.com/yzc/orange-review/rpc_gen/kitex_gen/review"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func CreateReview(ctx context.Context, req *review.CreateReviewRequest, callOptions ...callopt.Option) (resp *review.CreateReviewResponse, err error) {
	resp, err = defaultClient.CreateReview(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateReview call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
