package store_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vladjong/L0/internal/app/model"
	"github.com/vladjong/L0/internal/app/store"
)

func TestOrderRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("deliveries", "orders", "items", "payments")
	d, err := s.Order().Create(&model.Order{
		OrderId:     "b563feb7b2b84b6test",
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: model.Delivery{
			Name:    "Test",
			Phone:   "9720000000",
			Zip:     "2639809",
			City:    "Kiryat",
			Address: "Ploshad",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: model.Payment{
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
		Items:           []model.Item{},
		Locale:          "en",
		Signature:       "",
		Customer:        "test",
		DeliveryService: "meest",
		Shardkey:        "9",
		SmId:            9,
		DateOf:          time.Now(),
		OofShard:        "1",
	})
	assert.NoError(t, err)
	assert.NotNil(t, d)
}

func TestOrderRepository_FindByID(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("deliveries", "orders", "items", "payments")
	id := "b563feb7b2b84b6test"
	s.Order().Create(&model.Order{
		OrderId:     "b563feb7b2b84b6test",
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: model.Delivery{
			Name:    "Test",
			Phone:   "9720000000",
			Zip:     "2639809",
			City:    "Kiryat",
			Address: "Ploshad",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: model.Payment{
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
		Items:           []model.Item{},
		Locale:          "en",
		Signature:       "",
		Customer:        "test",
		DeliveryService: "meest",
		Shardkey:        "9",
		SmId:            9,
		DateOf:          time.Now(),
		OofShard:        "1",
	})
	o, err := s.Order().FindOrderId(id)
	assert.NoError(t, err)
	assert.Equal(t, o.Locale, "en")
	assert.NotNil(t, o)
}

func TestOrderRepository_FindByIDError(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("deliveries", "orders", "items", "payments")
	id := "b563feb7b2b84b6test"
	_, err := s.Order().FindOrderId(id)
	assert.Error(t, err)
}
