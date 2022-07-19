package store_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vladjong/L0/internal/app/model"
	"github.com/vladjong/L0/internal/app/store"
)

func TestOrderRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("deliveries", "orders", "items", "payments")
	order := model.TestOrder(t)
	err := s.Order().Create(order)
	assert.NoError(t, err)
	o, err := s.Order().FindOrderId(order.OrderId)
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
