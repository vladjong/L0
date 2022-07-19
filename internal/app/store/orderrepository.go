package store

import (
	"log"
	"testing"

	"github.com/vladjong/L0/internal/app/model"
)

const (
	INSERT_ORDER      = "insert into orders values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) returning order_uid, track_number;"
	INSERT_DELIVERIES = "insert into deliveries (name, phone, zip, city, address, region, email) values ($1, $2, $3, $4, $5, $6, $7) returning id;"
	INSERT_PAYMENT    = "insert into payments values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning transaction;"
	INSERT_ITEM       = "insert into items values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) returning track_number;"
	SELECT_ORDER_ID   = "select * from orders where order_uid = $1;"
	SELECT_DELIVERIES = "select * from deliveries where id = $1;"
	SELECT_PAYMENTS   = "select * from payments where transaction = $1;"
	SELECT_ITEMS      = "select * from items where track_number = $1;"
)

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) Create(order *model.Order) error {
	if err := order.Validate(); err != nil {
		return err
	}
	deliveryId, err := r.InsertDelivery(order)
	if err != nil {
		return err
	}
	err = r.InsertOrder(order, deliveryId)
	if err != nil {
		return err
	}
	err = r.InsertPayment(order)
	if err != nil {
		return err
	}
	err = r.InsertItem(order)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) InsertDelivery(order *model.Order) (int, error) {
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
		return 0, err
	}
	return deliveryId, nil
}

func (r *OrderRepository) InsertOrder(order *model.Order, deliveryId int) error {
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
		return err
	}
	return nil
}

func (r *OrderRepository) InsertPayment(order *model.Order) error {
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
		return err
	}
	return nil
}

func (r *OrderRepository) InsertItem(order *model.Order) error {
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
			return err
		}
	}
	return nil
}

func (r *OrderRepository) FindOrderId(id string) (*model.Order, error) {
	order := model.TestOrder(&testing.T{})
	err := r.SelectOrder(id, order)
	if err != nil {
		return nil, err
	}
	err = r.SelectDelivery(order)
	if err != nil {
		return nil, err
	}
	err = r.SelectPayment(order)
	if err != nil {
		return nil, err
	}
	err = r.SelectItem(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepository) SelectOrder(id string, order *model.Order) error {
	if err := r.store.db.QueryRow(SELECT_ORDER_ID,
		id,
	).Scan(
		&order.OrderId,
		&order.TrackNumber,
		&order.Entry,
		&order.Locale,
		&order.Delivery.Id,
		&order.Signature,
		&order.Customer,
		&order.DeliveryService,
		&order.Shardkey,
		&order.SmId,
		&order.DateOf,
		&order.OofShard,
	); err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) SelectDelivery(order *model.Order) error {
	if err := r.store.db.QueryRow(SELECT_DELIVERIES,
		order.Delivery.Id,
	).Scan(
		&order.Delivery.Id,
		&order.Delivery.Name,
		&order.Delivery.Phone,
		&order.Delivery.Zip,
		&order.Delivery.City,
		&order.Delivery.Address,
		&order.Delivery.Region,
		&order.Delivery.Email,
	); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *OrderRepository) SelectPayment(order *model.Order) error {
	if err := r.store.db.QueryRow(SELECT_PAYMENTS,
		order.OrderId,
	).Scan(
		&order.Payment.Transaction,
		&order.Payment.RequestId,
		&order.Payment.Currency,
		&order.Payment.Provider,
		&order.Payment.Amount,
		&order.Payment.PaymentDt,
		&order.Payment.Bank,
		&order.Payment.DeliveryCost,
		&order.Payment.GoodsTotal,
		&order.Payment.CustomFee,
	); err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) SelectItem(order *model.Order) error {
	if err := r.store.db.QueryRow(SELECT_ITEMS,
		order.TrackNumber,
	).Scan(
		&order.Items[0].Rid,
		&order.Items[0].ChrtId,
		&order.Items[0].TrackNumber,
		&order.Items[0].Price,
		&order.Items[0].Name,
		&order.Items[0].Sale,
		&order.Items[0].Size,
		&order.Items[0].TotalPrice,
		&order.Items[0].NmId,
		&order.Items[0].Brand,
		&order.Items[0].Status,
	); err != nil {
		return err
	}
	return nil
}
