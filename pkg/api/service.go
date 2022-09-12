package api

import (
	"info/pkg/domain"
	"info/pkg/facade"
	"info/pkg/infrastucture/kafka"
	"info/pkg/infrastucture/proto/serializer"
)

type (
	Options struct {
		KafkaServers string
	}
	ServiceProvider struct {
		actionFacade facade.ActionFacade
	}
)

func NewServiceProvider(actionFacade facade.ActionFacade) *ServiceProvider {
	return &ServiceProvider{actionFacade: actionFacade}
}

func (s ServiceProvider) ActionFacade() facade.ActionFacade {
	return s.actionFacade
}

func RegisterServices(opt *Options) *ServiceProvider {
	kafkaClient := kafka.NewProducer(opt.KafkaServers)

	var streamService domain.StreamService = kafka.NewStreamService(kafkaClient, serializer.NewStreamSerializer())
	var actionFacade facade.ActionFacade = facade.NewActionFacade(streamService)

	return NewServiceProvider(actionFacade)
}
