package helper

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shiffoo/wb-nats-streaming/internal/models"
)

func CreateOrder() models.Order {
	rand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyz"
	numbers := "1234567890"
	letAndNum := "abcdefghijklmnopqrstuvwxyz1234567890"

	name := randomStr(letters, 4)
	delivery := models.Delivery{
		ID:      uuid.NewString(),
		Name:    name + " " + name + "ov",
		Phone:   "+97" + randomStr(numbers, 9),
		Zip:     randomStr(numbers, 7),
		City:    randomStr(letters, 9),
		Address: randomStr(letters, 9) + randomStr(numbers, 3),
		Region:  randomStr(letters, 7),
		Email:   randomStr(letAndNum, 7) + "@gmail.com",
	}

	payment := models.Payment{
		ID:           uuid.NewString(),
		Transaction:  randomStr(letAndNum, 20),
		Request:      "",
		Currency:     randomStr(letters, 3),
		Provider:     randomStr(letters, 6),
		Amount:       rand.Intn(10000),
		PaymentDt:    1000 + rand.Intn(9000),
		Bank:         randomStr(letters, 7),
		DeliveryCost: 1000 + rand.Intn(9000),
		GoodsTotal:   100 + rand.Intn(900),
		CustomFee:    rand.Intn(10000),
	}

	item := models.Item{
		ID:          uuid.NewString(),
		ChrtId:      10000 + rand.Intn(90000),
		TrackNumber: randomStr(letters, 14),
		Price:       100 + rand.Intn(900),
		Rid:         randomStr(letAndNum, 20),
		Name:        randomStr(letters, 10),
		Sale:        rand.Intn(100),
		Size:        randomStr(numbers, 1),
		TotalPrice:  rand.Intn(100000),
		NmId:        10000 + rand.Intn(90000),
		Brand:       randomStr(letters, 10),
		Status:      rand.Intn(100),
	}

	timeOrder, _ := time.Parse(time.RFC3339, "2021-11-26T06:22:19Z")
	order := models.Order{
		ID:                uuid.NewString(),
		TrackNumber:       randomStr(letters, 14),
		Entry:             randomStr(letters, 4),
		Locale:            randomStr(letters, 2),
		InternalSignature: "",
		CustomerId:        randomStr(letters, 4),
		DeliveryService:   randomStr(letters, 4),
		ShardKey:          randomStr(numbers, 1),
		SmId:              uint(rand.Intn(100)),
		DateCreated:       timeOrder,
		OffShard:          randomStr(numbers, 1),
	}

	delivery.Order = order
	delivery.OrderId = order.ID
	item.Order = order
	item.OrderId = order.ID
	payment.Order = order
	payment.OrderId = order.ID

	order.Payment = &payment
	order.Delivery = &delivery
	order.Items = append(order.Items, item)

	return order
}

func randomStr(str string, lenStr int) string {
	var b strings.Builder
	for i := 0; i < lenStr; i++ {
		ch := string(str[rand.Intn(len(str))])
		b.WriteString(ch)
	}
	return b.String()
}
