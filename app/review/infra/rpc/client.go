package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/yzc/orange-review/app/review/conf"
	generator "github.com/yzc/orange-review/rpc_gen/kitex_gen/generator/idgenerateservice"
)

var (
	once            sync.Once
	GeneratorClient generator.Client
)

// InitAllClient 初始化所有 RPC 客户端
func InitAllClient() {
	once.Do(func() {
		initGeneratorClient()

		// 初始化其他客户端
	})
}

func initGeneratorClient() {
	// 服务发现
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		klog.Fatal(err)
	}
	GeneratorClient, err = generator.NewClient("generator", client.WithResolver(r))
	if err != nil {
		klog.Fatal(err)
	}
}
