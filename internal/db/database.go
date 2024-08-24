package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/shiffoo/wb-nats-streaming/internal/config"
	"github.com/shiffoo/wb-nats-streaming/internal/logger"
	"github.com/shiffoo/wb-nats-streaming/internal/models"
)

var DB *gorm.DB

func InitializeDB() error {
	var err error
	logg := logger.GetLogger()
	info := config.CONFIG.DB
	urlPostgres := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.Name)

	DB, err = gorm.Open("postgres", urlPostgres)
	if err != nil {
		logg.Panic("Failed to connect to database: ", err)
		return err
	}

	DB.AutoMigrate(&models.Order{}, &models.Delivery{}, &models.Item{}, &models.Payment{})

	err = DB.DB().Ping()
	if err != nil {
		logg.Panic(err)
		return err
	}
	logg.Info("Init database")

	return nil
}
func AddDbOrder(order models.Order) error {
	logg := logger.GetLogger()
	tx := DB.Begin()

	if err := tx.Create(&order).Error; err != nil {
		logg.Error("Failed to save order to DB: ", err)
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		logg.Error("Failed to commit transaction: ", err)
		return err
	}

	logg.Info("Order saved to DB successfully")
	return nil
}
