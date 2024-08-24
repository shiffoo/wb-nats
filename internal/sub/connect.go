package sub

import (
	"encoding/json"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/shiffoo/wb-nats-streaming/internal/cache"
	"github.com/shiffoo/wb-nats-streaming/internal/config"
	"github.com/shiffoo/wb-nats-streaming/internal/db"
	"github.com/shiffoo/wb-nats-streaming/internal/logger"
	"github.com/shiffoo/wb-nats-streaming/internal/models"
)

func Connect() {
	confNuts := config.CONFIG.Nats
	logg := logger.GetLogger()

	nc, _ := nats.Connect(confNuts.URL)
	js, err := nc.JetStream()

	if err != nil {
		logg.Fatal(err)
	}

	js.Subscribe(confNuts.StreamSubjects, func(msg *nats.Msg) {
		logg.Info("Received a message from NATS")
		msg.Ack()

		var order models.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			logg.Error("Failed to unmarshal message: ", err)
			return
		}
		logg.Infof("Processing order with ID: %s", order.ID)

		if err := db.AddDbOrder(order); err != nil {
			logg.Error("Failed to add order to DB: ", err)
		} else {
			cache.AddCacheOrder(order)
		}
	}, nats.Durable(confNuts.Subscriber), nats.ManualAck())

	logg.Info("Subscribe successfully")
	runtime.Goexit()

}
