package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/app"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/internal/repository"
)

type Order struct {
	app app.App
}

type OrderItem struct {
	ID        string         `json:"id"`
	UserID    string         `json:"user_id"`    //nolint:tagliatelle
	OrderDate time.Time      `json:"order_date"` //nolint:tagliatelle
	Items     []OrderProduct `json:"products"`
}

type OrderProduct struct {
	Price     int64  `json:"price"`
	ProductID string `json:"product_id"` //nolint:tagliatelle
}

func (o *Order) GetByID(ctx context.Context, ID string) (*OrderItem, error) { //nolint:gocritic
	res, err := o.app.Repository.RunTransactional(ctx, func(repo repository.Querier) (any, error) {
		uuid := pgtype.UUID{}
		err := uuid.UnmarshalJSON([]byte(ID))
		if err != nil {
			return nil, fmt.Errorf("convert id to uuid: %w", err)
		}
		order, qError := repo.GetOrderById(ctx, uuid)
		if qError != nil {
			return nil, fmt.Errorf("get order by id: %w", qError)
		}

		if order == nil {
			return nil, nil
		}

		orderProducts, err := repo.GetOrderProducts(ctx, order.ID)
		if err != nil {
			return nil, fmt.Errorf("get order products: %w", err)
		}

		orderItem := OrderItem{
			ID:        string(order.ID.Bytes[0:]),
			UserID:    string(order.UserID.Bytes[0:]),
			OrderDate: order.OrderDate.Time,
		}

		if orderProducts != nil {
			orderItem.Items = make([]OrderProduct, len(orderProducts))
			for index, orderProduct := range orderProducts {
				orderItem.Items[index] = OrderProduct{
					Price:     int64(orderProduct.Price),
					ProductID: string(orderProduct.ProductID.Bytes[0:]),
				}
			}
		}

		return orderItem, nil
	})
	if err != nil {
		return nil, fmt.Errorf("get order by id: %w", err)
	}
	order, ok := res.(OrderItem)
	if !ok {
		return nil, fmt.Errorf("convert item to order")
	}

	return &order, nil
}
