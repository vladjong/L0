package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

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

func (o *Order) Validate() error {
	return validation.ValidateStruct(o,
		validation.Field(&o.OrderId, validation.Required, validation.Length(1, 100)),
		validation.Field(&o.TrackNumber, validation.Required, validation.Length(1, 100)),
		validation.Field(&o.Entry, validation.Required, validation.Length(1, 100)),
		validation.Field(&o.Locale, validation.Required, validation.Length(1, 5)),
		validation.Field(&o.Customer, validation.Required, validation.Length(1, 100)),
		validation.Field(&o.DeliveryService, validation.Required, validation.Length(1, 100)),
		validation.Field(&o.Shardkey, validation.Required, validation.Length(1, 100)),
		validation.Field(&o.SmId, validation.Required),
		validation.Field(&o.DateOf, validation.Required),
		validation.Field(&o.OofShard, validation.Required),
	)
}
