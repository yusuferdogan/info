package kafka

import (
	"context"
	"info/pkg/domain"
	"info/pkg/infrastucture/proto/serializer"
)

const eventTopicName = "event"

type StreamService struct {
	p   *producer
	ser *serializer.StreamSerializer
}

func NewStreamService(p *producer, ser *serializer.StreamSerializer) *StreamService {
	return &StreamService{
		p:   p,
		ser: ser,
	}
}

func (s StreamService) Send(ctx context.Context, stream *domain.Stream) error {
	b, err := s.ser.Serialize(stream)
	if err != nil {
		return err
	}
	err = s.p.Produce(ctx, eventTopicName, b)
	if err != nil {
		return err
	}
	return nil
}
