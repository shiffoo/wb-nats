package main

import (
	"github.com/shiffoo/wb-nats-streaming/internal/logger"
	"github.com/shiffoo/wb-nats-streaming/internal/sub"
)

func main() {
	logg := logger.GetLogger()
	if err := sub.Run(); err != nil {
		logg.Fatal(err)
	}
}
