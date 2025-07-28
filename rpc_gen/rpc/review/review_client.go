package review

import (
	"context"
	review "github.com/yzc/orange-review/rpc_gen/kitex_gen/review"

	"github.com/yzc/orange-review/rpc_gen/kitex_gen/review/reviewservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() reviewservice.Client
	Service() string
	CreateReview(ctx context.Context, Req *review.CreateReviewRequest, callOptions ...callopt.Option) (r *review.CreateReviewResponse, err error)
	ListReviewByStoreID(ctx context.Context, Req *review.ListReviewByStoreIDReq, callOptions ...callopt.Option) (r *review.ListReviewByStoreIDResp, err error)
	ListReviewBySpuID(ctx context.Context, Req *review.ListReviewBySpuIDReq, callOptions ...callopt.Option) (r *review.ListReviewBySpuIDResp, err error)
	ListReviewBySkuID(ctx context.Context, Req *review.ListReviewBySkuIDReq, callOptions ...callopt.Option) (r *review.ListReviewBySkuIDResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := reviewservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient reviewservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() reviewservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateReview(ctx context.Context, Req *review.CreateReviewRequest, callOptions ...callopt.Option) (r *review.CreateReviewResponse, err error) {
	return c.kitexClient.CreateReview(ctx, Req, callOptions...)
}

func (c *clientImpl) ListReviewByStoreID(ctx context.Context, Req *review.ListReviewByStoreIDReq, callOptions ...callopt.Option) (r *review.ListReviewByStoreIDResp, err error) {
	return c.kitexClient.ListReviewByStoreID(ctx, Req, callOptions...)
}

func (c *clientImpl) ListReviewBySpuID(ctx context.Context, Req *review.ListReviewBySpuIDReq, callOptions ...callopt.Option) (r *review.ListReviewBySpuIDResp, err error) {
	return c.kitexClient.ListReviewBySpuID(ctx, Req, callOptions...)
}

func (c *clientImpl) ListReviewBySkuID(ctx context.Context, Req *review.ListReviewBySkuIDReq, callOptions ...callopt.Option) (r *review.ListReviewBySkuIDResp, err error) {
	return c.kitexClient.ListReviewBySkuID(ctx, Req, callOptions...)
}
