package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/yzc/orange-review/app/review_job/infra"
)

// Msg 定义kafka中接收到的数据
type Msg struct {
	Type     string `json:"type"`
	Database string `json:"databse"`
	Table    string `json:"table"`
	IsDdl    bool   `json:"isDdl"`
	Data     []map[string]interface{}
}

func main() {
	infra.InitES()
	infra.InitKafka()

	for {
		m, err := infra.KafkaReader.ReadMessage(context.Background())
		if errors.Is(err, context.Canceled) {
			log.Printf("readMessage from kafka failed, err:%v", err)
			return
		}
		if err != nil {
			log.Fatalf("readMessage from kafka failed, err:%v", err)
			break
		}
		log.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		// 2. 将完整评价数据写入ES
		msg := new(Msg)
		if err := json.Unmarshal(m.Value, msg); err != nil {
			log.Printf("unmarshal msg from kafka failed, err:%v", err)
			continue
		}

		// 补充！
		// 实际的业务场景可能需要在这增加一个步骤：对数据做业务处理
		// 例如：把两张表的数据合成一个文档写入ES

		if msg.Type == "INSERT" {
			// 往ES中新增文档
			for idx := range msg.Data {
				infra.IndexDocument(msg.Data[idx])
			}
		} else {
			// 往ES中更新文档
			for idx := range msg.Data {
				infra.UpdateDocument(msg.Data[idx])
			}
		}
	}
}
