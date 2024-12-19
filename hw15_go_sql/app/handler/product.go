package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/kotoproger/home_work_basic/hw15_go_sql/app"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/internal/repository"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/internal/repositorywrapper"
)

type Product struct {
	app app.App
}

func NewProduct(a app.App) *Product {
	return &Product{app: a}
}

type ProductItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

func (p *Product) GetList(ctx context.Context) ([]ProductItem, error) {
	res, err := p.app.Repository.RunTransactional(ctx, func(repo repository.Querier) (any, error) {
		products, err := repo.GetProducts(ctx)
		if err != nil {
			return nil, fmt.Errorf("get products from database: %w", err)
		}

		list := make([]ProductItem, len(products))
		for index, item := range products {
			suuid, uErr := repositorywrapper.UUIDToString(item.ID)
			if uErr != nil {
				log.Println(uErr)
				continue
			}
			list[index] = ProductItem{
				ID:    suuid,
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
