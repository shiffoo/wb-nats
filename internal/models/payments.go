package models

type Payment struct {
	ID           string `json:"id" gorm:"primaryKey"`
	Transaction  string `json:"transaction"`
	Request      string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
	OrderId      string `json:"order_id"`
	Order        Order  `json:"-"`
}
