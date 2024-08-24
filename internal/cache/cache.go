package cache

import (
	"errors"
	"sync"

	"github.com/shiffoo/wb-nats-streaming/internal/db"
	"github.com/shiffoo/wb-nats-streaming/internal/logger"
	"github.com/shiffoo/wb-nats-streaming/internal/models"
)

var CACHE *sync.Map

func Restore() error {
	CACHE = nil
	CACHE = &sync.Map{}

	logg := logger.GetLogger()
	var orders []models.Order

	if err := db.DB.Preload("Delivery").Preload("Payment").Preload("Items").Find(&orders).Error; err != nil {
		logg.Error(err)
		return err
	}

	for _, order := range orders {
		CACHE.Store(order.ID, order)
	}

	logg.Info("Restore cache")

	return nil
}

func AddCacheOrder(order models.Order) {
	CACHE.Store(order.ID, order)
}

func GetOrderFromCache(id string) (any, error) {
	logg := logger.GetLogger()
	order, ok := CACHE.Load(id)

	if ok {
		logg.Info("Get data from cache")
		return order, nil
	}

	var orderFind models.Order
	if err := db.DB.Preload("Delivery").Preload("Payment").Preload("Items").First(&orderFind, "id = ?", id).Error; err != nil {
		logg.Info("Data not found")
		return nil, errors.New("data not found")
	}

	Restore()
	logg.Info("Data found from DB")

	return orderFind, nil
}
