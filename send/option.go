package send

import (
	confluentKafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/kybsa/async-replication-golang/send/application"
	"github.com/kybsa/async-replication-golang/send/domain/repository"
	appGorm "github.com/kybsa/async-replication-golang/send/infrastructure/repository/gorm"
	"github.com/kybsa/async-replication-golang/send/infrastructure/repository/kafka"
	"gorm.io/gorm"
)

type Config struct {
	sendMessageRepository       repository.SendMessage
	sendMessageStatusRepository repository.SendMessageStatus
	queue                       repository.Queue
	database                    repository.Database
}

type Option func(*Config)

func New(options ...Option) application.SendMessage {
	config := &Config{}

	for _, opt := range options {
		opt(config)
	}

	return nil
}

func WithGorm(db *gorm.DB) Option {
	return func(config *Config) {
		config.sendMessageRepository = appGorm.NewSendMessage()
		config.sendMessageStatusRepository = appGorm.NewSendMessageStatus()
		config.database = appGorm.NewDatabase(db)
	}
}

func WithKafka(kafkaProducer *confluentKafka.ConfigMap, topic string) Option {
	return func(config *Config) {
		config.queue = kafka.NewProducer(kafkaProducer, topic)
	}
}
