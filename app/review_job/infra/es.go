package infra

import (
	"context"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/yzc/orange-review/app/review_job/conf"
)

var (
	ESClient *elasticsearch.TypedClient
	Index    string
)

func InitES() {
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
	Index = conf.GetConf().Elasticsearch.Index
	log.Printf("init es client success, addr:%s, index:%s\n", conf.GetConf().Elasticsearch.Addr, Index)
}

// IndexDocument 索引文档
func IndexDocument(d map[string]interface{}) {
	reviewID := d["review_id"].(string)
	// 添加文档
	resp, err := ESClient.Index(Index).
		Id(reviewID).
		Document(d).
		Do(context.Background())
	if err != nil {
		log.Fatalf("indexing document failed, err:%v\n", err)
		return
	}
	log.Printf("result:%#v\n", resp.Result)
}

// updateDocument 更新文档
func UpdateDocument(d map[string]interface{}) {
	reviewID := d["review_id"].(string)
	resp, err := ESClient.Update(Index, reviewID).
		Doc(d). // 使用结构体变量更新
		Do(context.Background())
	if err != nil {
		log.Fatalf("update document failed, err:%v\n", err)
		return
	}
	log.Fatalf("result:%v\n", resp.Result)
}
