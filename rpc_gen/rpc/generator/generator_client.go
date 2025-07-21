package generator

import (
	"context"
	generator "github.com/yzc/orange-review/rpc_gen/kitex_gen/generator"

	"github.com/yzc/orange-review/rpc_gen/kitex_gen/generator/idgenerateservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() idgenerateservice.Client
	Service() string
	GenID(ctx context.Context, Req *generator.GenIDRequest, callOptions ...callopt.Option) (r *generator.GenIDResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := idgenerateservice.NewClient(dstService, opts...)
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
	kitexClient idgenerateservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() idgenerateservice.Client {
	return c.kitexClient
}

func (c *clientImpl) GenID(ctx context.Context, Req *generator.GenIDRequest, callOptions ...callopt.Option) (r *generator.GenIDResponse, err error) {
	return c.kitexClient.GenID(ctx, Req, callOptions...)
}
