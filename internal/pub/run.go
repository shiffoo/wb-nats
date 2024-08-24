package pub

import (
	"github.com/shiffoo/wb-nats-streaming/internal/config"
	"github.com/shiffoo/wb-nats-streaming/internal/logger"
)

func Run() error {
	logg := logger.GetLogger()
	if err := config.InitConfig(); err != nil {
		logg.Panic(err)
		return err
	}
	if _, err := Connect(); err != nil {
		logg.Panic(err)
		return err
	}
	return nil
}
