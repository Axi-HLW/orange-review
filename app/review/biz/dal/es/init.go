package es

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/yzc/orange-review/app/review/conf"
)

var (
	ESClient *elasticsearch.TypedClient
)

func Init() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			conf.GetConf().Elasticsearch.Addr,
		},
	}
	var err error
	ESClient, err = elasticsearch.NewTypedClient(cfg)
	if err != nil {
		panic(err)
	}
}
