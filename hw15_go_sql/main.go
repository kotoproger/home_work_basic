package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/kotoproger/home_work_basic/configapp"
	"github.com/pressly/goose/v3"
)

func main() {
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

}
