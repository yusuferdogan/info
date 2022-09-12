package serializer

import (
	"info/pkg/domain"
	"info/pkg/domain/action"
	"info/pkg/infrastucture/proto/pb"

	"google.golang.org/protobuf/proto"
)

type StreamSerializer struct {
}

func NewStreamSerializer() *StreamSerializer {
	return &StreamSerializer{}
}

func (s StreamSerializer) Serialize(e *domain.Stream) ([]byte, error) {
	p := NewStreamProto(e)
	b, err := proto.Marshal(p)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func NewStreamProto(e *domain.Stream) *pb.Stream {
	if e == nil {
		return nil
	}
	return &pb.Stream{

		Action: NewActionProto(e.Action()),
	}
}

func NewActionProto(e *action.Action) *pb.Action {
	if e == nil {
		return nil
	}

	return &pb.Action{
		Id:         e.Id(),
		Category:   e.Category(),
		ActionType: e.ActionType(),
	}
}
