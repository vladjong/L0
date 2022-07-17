package model

import (
	"testing"
	"time"
)

func TestOrder(t *testing.T) *Order {
	return &Order{
		OrderId:         "b563feb7b2b84b6test",
		TrackNumber:     "WBILMTESTTRACK",
		Entry:           "WBIL",
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
