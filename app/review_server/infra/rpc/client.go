package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/yzc/orange-review/app/review_server/conf"
	generator "github.com/yzc/orange-review/rpc_gen/kitex_gen/generator/idgenerateservice"
)

var (
	GeneratorClient generator.Client

	once sync.Once
)

// InitAllClient 初始化所有客户端
func InitAllClient() {
	once.Do(func() {
		initGeneratorClient()

		// 初始化其他客户端
	})
}

func initGeneratorClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		klog.Fatalf("init consul resolver failed, err: %v", err)
	}
	GeneratorClient, err = generator.NewClient("generator", client.WithResolver(r))
	if err != nil {
		klog.Fatalf("init generator client failed, err: %v", err)
	}
}
