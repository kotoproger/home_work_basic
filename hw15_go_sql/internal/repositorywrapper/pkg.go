package repositorywrapper

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/internal/repository"
)

type TransactionalRepository interface {
	RunTransactional(ctx context.Context, execute func(repo repository.Querier) (any, error)) (any, error)
}

type Wrapper struct {
	Pool *pgxpool.Pool
	Repo *repository.Queries
}

func (w *Wrapper) RunTransactional(
	ctx context.Context,
	execute func(repo repository.Querier) ([]any, error),
) ([]any, error) {
	repo, commit, rollback, release, err := w.getRepository(ctx)
	if err != nil {
		return nil, fmt.Errorf("get repository: %w", err)
	}
	defer release()

	result, runError := execute(repo)
	if runError != nil {
		rollback()
		return result, fmt.Errorf("execute error: %w", runError)
	}

	commitError := commit()
	if commitError != nil {
		return result, fmt.Errorf("commit error: %w", commitError)
	}

	return result, nil
}

func (w *Wrapper) getRepository(ctx context.Context) (
	repo repository.Querier,
	commit func() error,
	rollback func(),
	release func(),
	err error,
) {
	conn, aqerror := w.Pool.Acquire(ctx)
	if aqerror != nil {
		return nil, nil, nil, nil, fmt.Errorf("get repository: %w", aqerror)
	}
	transaction, err := conn.BeginTx(
		ctx,
		pgx.TxOptions{IsoLevel: pgx.ReadCommitted},
	)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("start transaction: %w", err)
	}

	release = func(conn *pgxpool.Conn) func() {
		return func() {
			defer conn.Release()
		}
	}(conn)

	commit = func(tr pgx.Tx, ctx context.Context) func() error {
		return func() error {
			return tr.Commit(ctx)
		}
	}(transaction, ctx)

	rollback = func(tr pgx.Tx, ctx context.Context) func() {
		return func() {
			defer tr.Rollback(ctx)
		}
	}(transaction, ctx)

	return w.Repo.WithTx(transaction), commit, rollback, release, nil
}
