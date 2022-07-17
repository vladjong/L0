package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vladjong/L0/internal/app/model"
	"github.com/vladjong/L0/internal/app/store/sqlstore"
)

func TestOrderRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("deliveries", "orders", "items", "payments")
	s := sqlstore.New(db)
	o := model.TestOrder(t)
	assert.NoError(t, s.Order().Create(o))
	assert.NotNil(t, o)
}

func TestOrderRepository_FindByID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("deliveries", "orders", "items", "payments")
	id := "b563feb7b2b84b6test"
	s := sqlstore.New(db)
	o := model.TestOrder(t)
	s.Order().Create(o)
	o, err := s.Order().FindOrderId(id)
	assert.NoError(t, err)
	assert.Equal(t, o.Locale, "en")
	assert.NotNil(t, o)
}

func TestOrderRepository_FindByIDError(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("deliveries", "orders", "items", "payments")
	s := sqlstore.New(db)
	id := "b563feb7b2b84b6test"
	_, err := s.Order().FindOrderId(id)
	assert.Error(t, err)
}
