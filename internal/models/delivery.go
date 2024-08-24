package models

type Delivery struct {
	ID      string `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
	OrderId string `json:"order_id"`
	Order   Order  `json:"-"`
}
