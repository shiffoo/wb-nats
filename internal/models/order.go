package models

import (
	"time"
)

type Order struct {
	ID                string    `json:"order_uid" gorm:"primaryKey"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          *Delivery `json:"delivery"`
	Payment           *Payment  `json:"payment"`
	Items             []Item    `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          string    `json:"shardkey"`
	SmId              uint      `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OffShard          string    `json:"off_shard"`
}
