// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddProductToOrder(ctx context.Context, arg AddProductToOrderParams) (pgtype.UUID, error)
	CreateOrder(ctx context.Context, userID pgtype.UUID) (pgtype.UUID, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (pgtype.UUID, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (pgtype.UUID, error)
	FindUser(ctx context.Context, email string) (*GeneralUser, error)
	GetOrderById(ctx context.Context, id pgtype.UUID) (*GeneralOrder, error)
	GetOrderProducts(ctx context.Context, orderID pgtype.UUID) ([]*GeneralOrderProduct, error)
	GetOrdersByUser(ctx context.Context, userID pgtype.UUID) ([]*GeneralOrder, error)
	GetProductById(ctx context.Context, id pgtype.UUID) (*GeneralProduct, error)
	RemoveProductFromOrder(ctx context.Context, id pgtype.UUID) error
	UpdateOrderAmount(ctx context.Context, arg UpdateOrderAmountParams) (int32, error)
}

var _ Querier = (*Queries)(nil)