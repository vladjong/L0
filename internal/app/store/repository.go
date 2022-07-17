package store

import "github.com/vladjong/L0/internal/app/model"

type OrderRepository interface {
	Create(*model.Order) error
	FindOrderId(string) (*model.Order, error)
}
