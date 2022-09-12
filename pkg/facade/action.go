package facade

import (
	"context"
	"info/pkg/domain"
	"info/pkg/domain/action"
	"info/pkg/domain/user"
	"log"
)

type ActionFacade interface {
	Action(ctx context.Context) (string, string)
}

type actionFacade struct {
	streamService domain.StreamService
}

func NewActionFacade(streamService domain.StreamService) *actionFacade {
	return &actionFacade{
		streamService: streamService,
	}
}

func (f actionFacade) Action(ctx context.Context) (string, string) {
	action := action.NewAction("id", "type", "category")

	newContext := user.NewContextWithValue(ctx)

	go func() {
		err := f.streamService.Send(
			newContext,
			domain.NewStream(action),
		)
		if err != nil {
			log.Println(err)
		}
	}()

	return "userId", "actionId"
}
