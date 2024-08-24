package sub

import (
	"github.com/shiffoo/wb-nats-streaming/internal/cache"
	"github.com/shiffoo/wb-nats-streaming/internal/config"
	"github.com/shiffoo/wb-nats-streaming/internal/db"
	"github.com/shiffoo/wb-nats-streaming/internal/logger"
	"github.com/shiffoo/wb-nats-streaming/internal/transport"
)

func Run() error {
	logg := logger.GetLogger()

	if err := config.InitConfig(); err != nil {
		logg.Panic(err)
		return err
	}

	if err := db.InitializeDB(); err != nil {
		logg.Panic(err)
		return err
	}

	if err := cache.Restore(); err != nil {
		logg.Panic(err)
		return err
	}

	go Connect()

	if err := transport.InitRouter(); err != nil {
		logg.Panic(err)
		return err
	}

	return nil
}
