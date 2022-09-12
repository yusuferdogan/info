package main

import "info/pkg/api"

func options() *api.Options{
	return &api.Options{
		KafkaServers: "localhost:29092",
	}
}