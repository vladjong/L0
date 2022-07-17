package model

import "time"

type Order struct {
	OrderId         string `json:"order_uid"`
	TrackNumber     string `json:"track_number"`
	Entry           string `json:"entry"`
	Delivery        Delivery
	Payment         Payment
	Items           []Item
	Locale          string    `json:"locale"`
	Signature       string    `json:"internal_signature"`
	Customer        string    `json:"customer_id"`
	DeliveryService string    `json:"delivery_service"`
	Shardkey        string    `json:"shardkey"`
	SmId            int       `json:"sm_id"`
	DateOf          time.Time `json:"date_created"`
	OofShard        string    `json:"oof_shard"`
}
