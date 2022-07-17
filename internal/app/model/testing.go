package model

import (
	"testing"
	"time"
)

func TestOrder(t *testing.T) *Order {
	return &Order{
		OrderId:     "b563feb7b2b84b6test",
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: Delivery{
			Name:    "Test",
			Phone:   "9720000000",
			Zip:     "2639809",
			City:    "Kiryat",
			Address: "Ploshad",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: Payment{
			Transaction:  "b563feb7b2b84b6test",
			RequestId:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       1817,
			PaymentDt:    1637907727,
			Bank:         "alpha",
			DeliveryCost: 1500,
			GoodsTotal:   317,
			CustomFee:    0,
		},
		Items: []Item{
			{
				ChrtId:      9934930,
				TrackNumber: "WBILMTESTTRACK",
				Price:       453,
				Rid:         "ab4219087a764ae0btest",
				Name:        "Mascaras",
				Sale:        30,
				Size:        "0",
				TotalPrice:  317,
				NmId:        2389212,
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
		},
		Locale:          "en",
		Signature:       "",
		Customer:        "test",
		DeliveryService: "meest",
		Shardkey:        "9",
		SmId:            9,
		DateOf:          time.Now(),
		OofShard:        "1",
	}
}
