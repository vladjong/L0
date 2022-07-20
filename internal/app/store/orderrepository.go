package store

import (
	"log"
	"testing"

	"github.com/vladjong/L0/internal/app/model"
)

const (
	INSERT_ORDER         = "insert into orders values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) returning order_uid, track_number;"
	INSERT_DELIVERIES    = "insert into deliveries (name, phone, zip, city, address, region, email) values ($1, $2, $3, $4, $5, $6, $7) returning id;"
	INSERT_PAYMENT       = "insert into payments values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning transaction;"
	INSERT_ITEM          = "insert into items values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) returning track_number;"
	SELECT_ORDER_ID      = "select * from orders where order_uid = $1;"
	SELECT_DELIVERIES_ID = "select * from deliveries where id = $1;"
	SELECT_PAYMENTS_ID   = "select * from payments where transaction = $1;"
	SELECT_ITEMS_ID      = "select * from items where track_number = $1;"
	SELECT_ORDER         = "select * from orders;"
	SELECT_DELIVERIES    = "select * from deliveries;"
	SELECT_PAYMENTS      = "select * from payments;"
	SELECT_ITEMS         = "select * from items;"
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

func (r *OrderRepository) SelectAll() ([]model.Order, error) {
	var orders []model.Order
	orders, err := r.SelectOrder(orders)
	if err != nil {
		return nil, err
	}
	orders, err = r.SelectDelivery(orders)
	if err != nil {
		return nil, err
	}
	orders, err = r.SelectPayment(orders)
	if err != nil {
		return nil, err
	}
	orders, err = r.SelectItem(orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) SelectOrder(orders []model.Order) ([]model.Order, error) {
	order := model.TestOrder(&testing.T{})
	tmp, err := r.store.db.Query(SELECT_ORDER)
	if err != nil {
		return nil, err
	}
	for tmp.Next() {
		if err := tmp.Scan(
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
			return nil, err
		}
		orders = append(orders, *order)
	}
	return orders, nil
}

func (r *OrderRepository) SelectDelivery(orders []model.Order) ([]model.Order, error) {
	tmp, err := r.store.db.Query(SELECT_DELIVERIES)
	if err != nil {
		return nil, err
	}
	for i := 0; tmp.Next() && i < len(orders); i++ {
		if err := tmp.Scan(
			&orders[i].Delivery.Id,
			&orders[i].Delivery.Name,
			&orders[i].Delivery.Phone,
			&orders[i].Delivery.Zip,
			&orders[i].Delivery.City,
			&orders[i].Delivery.Address,
			&orders[i].Delivery.Region,
			&orders[i].Delivery.Email,
		); err != nil {
			return nil, err
		}
	}
	return orders, nil
}

func (r *OrderRepository) SelectPayment(orders []model.Order) ([]model.Order, error) {
	tmp, err := r.store.db.Query(SELECT_PAYMENTS)
	if err != nil {
		return nil, err
	}
	for i := 0; tmp.Next() && i < len(orders); i++ {
		if err := tmp.Scan(
			&orders[i].Payment.Transaction,
			&orders[i].Payment.RequestId,
			&orders[i].Payment.Currency,
			&orders[i].Payment.Provider,
			&orders[i].Payment.Amount,
			&orders[i].Payment.PaymentDt,
			&orders[i].Payment.Bank,
			&orders[i].Payment.DeliveryCost,
			&orders[i].Payment.GoodsTotal,
			&orders[i].Payment.CustomFee,
		); err != nil {
			return nil, err
		}
	}
	return orders, nil
}

func (r *OrderRepository) SelectItem(orders []model.Order) ([]model.Order, error) {
	tmp, err := r.store.db.Query(SELECT_ITEMS)
	if err != nil {
		return nil, err
	}
	for i := 0; tmp.Next() && i < len(orders); i++ {
		if err := tmp.Scan(
			&orders[i].Items[0].Rid,
			&orders[i].Items[0].ChrtId,
			&orders[i].Items[0].TrackNumber,
			&orders[i].Items[0].Price,
			&orders[i].Items[0].Name,
			&orders[i].Items[0].Sale,
			&orders[i].Items[0].Size,
			&orders[i].Items[0].TotalPrice,
			&orders[i].Items[0].NmId,
			&orders[i].Items[0].Brand,
			&orders[i].Items[0].Status,
		); err != nil {
			return nil, err
		}
	}
	return orders, nil
}

func (r *OrderRepository) FindOrderId(id string) (*model.Order, error) {
	order := model.TestOrder(&testing.T{})
	err := r.SelectOrderID(id, order)
	if err != nil {
		return nil, err
	}
	err = r.SelectDeliveryID(order)
	if err != nil {
		return nil, err
	}
	err = r.SelectPaymentID(order)
	if err != nil {
		return nil, err
	}
	err = r.SelectItemID(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepository) SelectOrderID(id string, order *model.Order) error {
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

func (r *OrderRepository) SelectDeliveryID(order *model.Order) error {
	if err := r.store.db.QueryRow(SELECT_DELIVERIES_ID,
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

func (r *OrderRepository) SelectPaymentID(order *model.Order) error {
	if err := r.store.db.QueryRow(SELECT_PAYMENTS_ID,
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

func (r *OrderRepository) SelectItemID(order *model.Order) error {
	if err := r.store.db.QueryRow(SELECT_ITEMS_ID,
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
