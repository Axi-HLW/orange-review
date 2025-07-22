package main

import (
	"context"
	generator "github.com/yzc/orange-review/rpc_gen/kitex_gen/generator"
	"github.com/yzc/orange-review/app/generator/biz/service"
)

// IDGenerateServiceImpl implements the last service interface defined in the IDL.
type IDGenerateServiceImpl struct{}

// GenID implements the IDGenerateServiceImpl interface.
func (s *IDGenerateServiceImpl) GenID(ctx context.Context, req *generator.GenIDRequest) (resp *generator.GenIDResponse, err error) {
	resp, err = service.NewGenIDService(ctx).Run(req)

	return resp, err
}
