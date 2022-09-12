package kafka

import (
	"context"
	"fmt"
	"info/pkg/domain/user"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type producer struct {
	producer *kafka.Producer
}

func NewProducer(servers string) *producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
	})
	if err != nil {
		panic(err)
	}
	return &producer{producer: p}
}

func (c producer) Produce(ctx context.Context, topic string, msg []byte) error {
	u := user.FromContext(ctx)
	formattedTopic := topic + "." + u.Id() + "." + u.Firstname() + "." + u.Lastname()
	deliveryChan := make(chan kafka.Event)
	err := c.producer.Produce(&kafka.Message{
		Headers: newHeaders(u),
		TopicPartition: kafka.TopicPartition{
			Topic:     &formattedTopic,
			Partition: kafka.PartitionAny,
		},
		Value: msg,
	}, deliveryChan)
	if err != nil {
		return err
	}
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		fmt.Printf("%v", m.TopicPartition.Error)
	}
	close(deliveryChan)
	c.producer.Flush(0)
	return nil
}

func newHeaders(u *user.User) []kafka.Header {
	m := make([]kafka.Header, 0)
	m = append(m, kafka.Header{
		Key:   "userId",
		Value: []byte(u.Id()),
	})
	m = append(m, kafka.Header{
		Key:   "userName",
		Value: []byte(u.Firstname()),
	})
	m = append(m, kafka.Header{
		Key:   "userLastname",
		Value: []byte(u.Lastname()),
	})
	m = append(m, kafka.Header{
		Key:   "userTitle",
		Value: []byte(u.Title()),
	})

	return m
}
