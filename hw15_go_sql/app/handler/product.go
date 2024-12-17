package handler

import (
	"context"
	"fmt"

	"github.com/kotoproger/home_work_basic/hw15_go_sql/app"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/internal/repository"
)

type Product struct {
	app app.App
}

type ProductItem struct {
	ID    string
	Name  string
	Price int64
}

func (p *Product) GetList(ctx context.Context) ([]ProductItem, error) {
	res, err := p.app.Repository.RunTransactional(ctx, func(repo repository.Querier) (any, error) {
		products, err := repo.GetProducts(ctx)
		if err != nil {
			return nil, fmt.Errorf("get products from database: %w", err)
		}

		list := make([]ProductItem, len(products))
		for index, item := range products {
			list[index] = ProductItem{
				ID:    string(item.ID.Bytes[0:]),
				Name:  item.Name,
				Price: int64(item.Price),
			}
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	list, ok := res.([]ProductItem)
	if !ok {
		return nil, fmt.Errorf("convert repository result in products")
	}

	return list, nil
}
