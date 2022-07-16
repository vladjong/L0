package store

import (
	"github.com/vladjong/L0/internal/app/model"
)

const (
	INSERT_ORDER      = "insert into orders values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) returning order_uid, track_number;"
	INSERT_DELIVERIES = "insert into deliveries (name, phone, zip, city, address, region, email) values ($1, $2, $3, $4, $5, $6, $7) returning id;"
	INSERT_PAYMENT    = "insert into payments values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning transaction;"
	INSERT_ITEM       = "insert into items values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) returning track_number;"
)

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) Create(order *model.Order) (*model.Order, error) {
	var deliveryId int
	if err := r.store.db.QueryRow(INSERT_DELIVERIES,
		order.Delivery.Name,
		order.Delivery.Phone,
		order.Delivery.Zip,
		order.Delivery.City,
		order.Delivery.Address,
		order.Delivery.Region,
		order.Delivery.Email,
	).Scan(&deliveryId); err != nil {
		return nil, err
	}

	if err := r.store.db.QueryRow(INSERT_ORDER,
		order.OrderId,
		order.TrackNumber,
		order.Entry,
		order.Locale,
		deliveryId,
		order.Signature,
		order.Customer,
		order.DeliveryService,
		order.Shardkey,
		order.SmId,
		order.DateOf,
		order.OofShard,
	).Scan(&order.OrderId, &order.TrackNumber); err != nil {
		return nil, err
	}

	if err := r.store.db.QueryRow(INSERT_PAYMENT,
		order.Payment.Transaction,
		order.Payment.RequestId,
		order.Payment.Currency,
		order.Payment.Provider,
		order.Payment.Amount,
		order.Payment.PaymentDt,
		order.Payment.Bank,
		order.Payment.DeliveryCost,
		order.Payment.GoodsTotal,
		order.Payment.CustomFee,
	).Scan(&order.Payment.Transaction); err != nil {
		return nil, err
	}

	for _, item := range order.Items {
		if err := r.store.db.QueryRow(INSERT_ITEM,
			item.Rid,
			item.ChrtId,
			item.TrackNumber,
			item.Price,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmId,
			item.Brand,
			item.Status,
		).Scan(&item.TrackNumber); err != nil {
			return nil, err
		}
	}

	return order, nil
}

func (r *OrderRepository) FindOrderId(id string) (*model.Order, error) {
	return nil, nil
}
