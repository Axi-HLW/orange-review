package generator

import (
	"context"
	generator "github.com/yzc/orange-review/rpc_gen/kitex_gen/generator"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GenID(ctx context.Context, req *generator.GenIDRequest, callOptions ...callopt.Option) (resp *generator.GenIDResponse, err error) {
	resp, err = defaultClient.GenID(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GenID call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
