package kafka

import (
	"context"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/kybsa/async-replication-golang/send/domain/entity"
)

type producer struct {
	configMap *confluentKafka.ConfigMap
	target    string
}

func NewProducer(configMap *confluentKafka.ConfigMap, target string) *producer {
	return &producer{
		configMap: configMap,
		target:    target,
	}
}

func (_self *producer) Send(ctx context.Context, message entity.SendMessage) (err error) {
	copyConfig := make(confluentKafka.ConfigMap)
	for k, v := range *_self.configMap {
		copyConfig[k] = v
	}

	copyConfig["transactional.id"] = message.IdempotencyKey
	copyConfig["enable.idempotence"] = true

	kafkaProducer, err := confluentKafka.NewProducer(&copyConfig)

	if err != nil {
		return
	}
	defer kafkaProducer.Close()

	if err = kafkaProducer.InitTransactions(nil); err != nil {
		return
	}

	if err = kafkaProducer.BeginTransaction(); err != nil {
		return
	}

	err = kafkaProducer.Produce(&confluentKafka.Message{
		TopicPartition: confluentKafka.TopicPartition{Topic: &_self.target, Partition: confluentKafka.PartitionAny},
		Value:          message.Message,
	}, nil)

	if err == nil {
		return
	}

	return kafkaProducer.AbortTransaction(nil)
}
