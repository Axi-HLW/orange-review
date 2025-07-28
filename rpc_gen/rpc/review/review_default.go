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

func ListReviewByStoreID(ctx context.Context, req *review.ListReviewByStoreIDReq, callOptions ...callopt.Option) (resp *review.ListReviewByStoreIDResp, err error) {
	resp, err = defaultClient.ListReviewByStoreID(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListReviewByStoreID call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListReviewBySpuID(ctx context.Context, req *review.ListReviewBySpuIDReq, callOptions ...callopt.Option) (resp *review.ListReviewBySpuIDResp, err error) {
	resp, err = defaultClient.ListReviewBySpuID(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListReviewBySpuID call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListReviewBySkuID(ctx context.Context, req *review.ListReviewBySkuIDReq, callOptions ...callopt.Option) (resp *review.ListReviewBySkuIDResp, err error) {
	resp, err = defaultClient.ListReviewBySkuID(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListReviewBySkuID call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
