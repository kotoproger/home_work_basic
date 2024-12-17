package app

import (
	"github.com/kotoproger/home_work_basic/hw15_go_sql/internal/repositorywrapper"
)

type App struct {
	Repository repositorywrapper.TransactionalRepository
}
