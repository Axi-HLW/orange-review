package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/yzc/orange-review/app/gateway/biz/utils"
	"github.com/yzc/orange-review/app/gateway/conf"
	"github.com/yzc/orange-review/rpc_gen/kitex_gen/review/reviewservice"
)

var (
	once sync.Once
	ReviewClient reviewservice.Client
)

// InitAllClient 初始化所有 RPC 客户端
func InitAllClient() {
	once.Do(func() {
		initReviewClient()

		// 初始化其他客户端
	})
}

func initReviewClient() {
	// 服务发现
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	utils.MustHandleError(err)
	ReviewClient, err = reviewservice.NewClient("review", client.WithResolver(r))
	utils.MustHandleError(err)
}
