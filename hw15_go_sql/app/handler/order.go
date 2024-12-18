package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/app"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/internal/repository"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/internal/repositorywrapper"
)

type Order struct {
	app app.App
}

func NewOrder(a app.App) *Order {
	return &Order{app: a}
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

func (o *Order) GetByID(ctx context.Context, orderID string, userID string) (*OrderItem, error) {
	orderUUID := pgtype.UUID{}
	ouErr := orderUUID.Scan(orderID)
	if ouErr != nil {
		return nil, fmt.Errorf("convert order id to uuid: %w", ouErr)
	}
	userUUID := pgtype.UUID{}
	uuErr := userUUID.Scan(userID)
	if uuErr != nil {
		return nil, fmt.Errorf("convert user id to uuid: %w", uuErr)
	}
	res, err := o.app.Repository.RunTransactional(ctx, func(repo repository.Querier) (any, error) {
		order, qError := repo.GetOrderById(ctx, repository.GetOrderByIdParams{
			ID:     orderUUID,
			UserID: userUUID,
		})
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
			ID:        orderID,
			UserID:    userID,
			OrderDate: order.OrderDate.Time,
		}

		if orderProducts != nil {
			orderItem.Items = make([]OrderProduct, len(orderProducts))
			for index, orderProduct := range orderProducts {
				prID, pUerr := repositorywrapper.UUIDToString(orderProduct.ProductID)
				if pUerr != nil {
					return nil, fmt.Errorf("convert product id to string: %w", pUerr)
				}
				orderItem.Items[index] = OrderProduct{
					Price:     int64(orderProduct.Price),
					ProductID: prID,
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

func (o *Order) DeleteOrder(ctx context.Context, orderID string) error {
	orderUUID := pgtype.UUID{}
	ouErr := orderUUID.Scan(orderID)
	if ouErr != nil {
		return fmt.Errorf("convert order id to uuid: %w", ouErr)
	}

	_, err := o.app.Repository.RunTransactional(ctx, func(repo repository.Querier) (any, error) {
		err := repo.RemoveAllProductsFromOrder(ctx, orderUUID)
		if err != nil {
			return nil, err
		}
		err = repo.DeleteOrderById(ctx, orderUUID)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	return err
}
