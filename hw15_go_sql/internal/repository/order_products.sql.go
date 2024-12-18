// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: order_products.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const AddProductToOrder = `-- name: AddProductToOrder :one
insert into general.order_products (order_id, product_id, price)
values ($1, $2, $3)
returning id
`

type AddProductToOrderParams struct {
	OrderID   pgtype.UUID `db:"order_id" json:"order_id"`
	ProductID pgtype.UUID `db:"product_id" json:"product_id"`
	Price     int32       `db:"price" json:"price"`
}

func (q *Queries) AddProductToOrder(ctx context.Context, arg AddProductToOrderParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, AddProductToOrder, arg.OrderID, arg.ProductID, arg.Price)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}

const GetOrderProducts = `-- name: GetOrderProducts :many
select id, order_id, product_id, price from general.order_products where order_id = $1
`

func (q *Queries) GetOrderProducts(ctx context.Context, orderID pgtype.UUID) ([]*GeneralOrderProduct, error) {
	rows, err := q.db.Query(ctx, GetOrderProducts, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GeneralOrderProduct{}
	for rows.Next() {
		var i GeneralOrderProduct
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.ProductID,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const RemoveAllProductsFromOrder = `-- name: RemoveAllProductsFromOrder :exec
delete from general.order_products where order_id = $1
`

func (q *Queries) RemoveAllProductsFromOrder(ctx context.Context, orderID pgtype.UUID) error {
	_, err := q.db.Exec(ctx, RemoveAllProductsFromOrder, orderID)
	return err
}

const RemoveProductFromOrder = `-- name: RemoveProductFromOrder :exec
delete from general.order_products where id = $1
`

func (q *Queries) RemoveProductFromOrder(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, RemoveProductFromOrder, id)
	return err
}
