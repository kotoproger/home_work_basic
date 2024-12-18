package main

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/kotoproger/home_work_basic/configapp"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/app"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/internal/repository"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/internal/repositorywrapper"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/server"
	"github.com/pressly/goose/v3"
)

func main() {
	godotenv.Load()
	c := configapp.ConfigApp{}

	c.AddParam(configapp.ConfigParam{
		Name: "http_port", Description: "Port", ShortName: "p", Default: "8880",
	})
	c.AddParam(configapp.ConfigParam{
		Name: "database_url", Description: "database url", ShortName: "d", Default: "",
	})

	configapp.GetConfig(c)

	pool, err := pgxpool.New(context.Background(), c.GetString("database_url"))
	if err != nil {
		panic("cant create connection pool")
	}

	goose.SetDialect("postgres")

	db := stdlib.OpenDBFromPool(pool)
	if db == nil {
		panic("cannot open db")
	}

	println("Migrating")

	err = goose.Up(db, "./sql/migrations")
	if err != nil {
		panic(err.Error())
	}
	println("Done")

	app := app.App{
		Repository: &repositorywrapper.Wrapper{
			Pool: pool,
			Repo: &repository.Queries{},
		},
		Ctx: context.Background(),
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		server.Start(app, c)
	}(&wg)

	wg.Wait()
}
