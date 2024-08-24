package pub

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/shiffoo/wb-nats-streaming/internal/config"
	"github.com/shiffoo/wb-nats-streaming/internal/helper"
	"github.com/shiffoo/wb-nats-streaming/internal/logger"
)

func Connect() (string, error) {
	confNuts := config.CONFIG.Nats
	logg := logger.GetLogger()

	nc, err := nats.Connect(confNuts.URL)

	if err != nil {
		logg.Panic(err)
		return "", err
	}

	defer nc.Close()

	js, err := nc.JetStream()

	if err != nil {
		logg.Panic(err)
		return "", err
	}

	if err = createStream(js); err != nil {
		logg.Panic(err)
		return "", err
	}

	order := helper.CreateOrder()
	orderJSON, err := json.Marshal(order)

	if err != nil {
		logg.Panic(err)
		return "", err
	}
	logg.Info("Stream created successfully or already exists")

	_, err = js.Publish(confNuts.SubjectName, orderJSON)
	if err != nil {
		logg.Panic("Failed to publish message to NATS: ", err)
		return "", err
	}
	logg.Infof("Order with ID:%s has been published", order.ID)

	return order.ID, nil
}

func createStream(js nats.JetStreamContext) error {
	logg := logger.GetLogger()
	confNuts := config.CONFIG.Nats

	stream, _ := js.StreamInfo(confNuts.StreamName)

	if stream == nil {
		logg.Printf("creating stream %q and subjects %q", confNuts.StreamName, confNuts.StreamSubjects)
		_, err := js.AddStream(&nats.StreamConfig{
			Name:     confNuts.StreamName,
			Subjects: []string{confNuts.StreamSubjects},
		})

		if err != nil {
			logg.Println(err)
			return err
		}
	}

	return nil
}
