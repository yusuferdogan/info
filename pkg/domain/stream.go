package domain

import (
	"context"
	"info/pkg/domain/action"
)

type (
	StreamService interface {
		Send(ctx context.Context, stream *Stream) error
	}
	Stream struct {
		action *action.Action
	}
)

func NewStream(action *action.Action) *Stream {
	return &Stream{

		action: action,
	}
}

func (s Stream) Action() *action.Action {
	return s.action
}
