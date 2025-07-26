package infra

import (
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/yzc/orange-review/app/review_job/conf"
)

var KafkaReader *kafka.Reader

func InitKafka() {
	KafkaReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: conf.GetConf().Kafka.Addrs,
		GroupID: conf.GetConf().Kafka.GroupID,
		Topic:   conf.GetConf().Kafka.Topic,
	})
	log.Printf("init kafka reader success, addr:%s, group_id:%s, topic:%s\n", conf.GetConf().Kafka.Addrs, conf.GetConf().Kafka.GroupID, conf.GetConf().Kafka.Topic)
}
