package publisher

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
)

type Publisher struct {
	writer *kafka.Writer
}

func NewPublisher(brokers []string, topic string) *Publisher {
	return &Publisher{
		writer: &kafka.Writer{
			Addr:                   kafka.TCP(brokers...),
			Balancer:               &kafka.LeastBytes{},
			Topic:                  topic,
			RequiredAcks:           kafka.RequireAll,
			MaxAttempts:            3,
			ErrorLogger:            kafka.LoggerFunc(func(msg string, args ...interface{}) { fmt.Println(msg) }),
			Compression:            compress.None,
			ReadTimeout:            10 * time.Second,
			WriteTimeout:           10 * time.Second,
			Async:                  false,
			BatchTimeout:           2 * time.Millisecond,
			BatchSize:              10,
			AllowAutoTopicCreation: false,
		},
	}
}

func (p *Publisher) Publish(message []byte) error {
	return p.writer.WriteMessages(context.Background(), kafka.Message{Value: message})
}
