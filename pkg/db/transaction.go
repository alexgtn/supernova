package db

import (
	"context"

	"github.com/pkg/errors"

	ent "github.com/alexgtn/supernova/internal/infra/ent/gen"
)

type ContextKey string

const ContextKeyTx ContextKey = "transaction"

type Transaction struct {
	tx *ent.Tx
}

type TxContext struct {
	ctx    context.Context
	cancel context.CancelFunc
	tx     *Transaction
}

// repo is expected to set tx
func NewTxContext(ctx context.Context) *TxContext {
	tx := &Transaction{tx: nil}

	ctxVal := context.WithValue(ctx, ContextKeyTx, tx)

	ctxCancelVal, cancel := context.WithCancel(ctxVal)

	return &TxContext{
		ctx:    ctxCancelVal,
		cancel: cancel,
		tx:     tx,
	}
}

// ErrRollback rolls back tx, cancels context and returns error
func (t *TxContext) ErrRollback(error error) error {
	defer t.cancel()

	if t.tx.tx == nil {
		return error
	}

	if err := t.tx.tx.Rollback(); err != nil {
		return error
	}

	return error
}

func (t *TxContext) Commit() error {
	if t.tx.tx == nil {
		return errors.New("failed to commit tx, tx is nil")
	}

	if err := t.tx.tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit tx")
	}

	return nil
}

func (t *TxContext) GetContext() context.Context {
	return t.ctx
}

// return existing tx, new tx or client if tx is not setup
func OrClient(ctx context.Context, client *ent.Client) *ent.Client {
	tx, ok := ctx.Value(ContextKeyTx).(*Transaction)
	if !ok {
		return client
	}

	if tx.tx == nil {
		// begin new tx
		newTx, err := client.Tx(ctx)
		if err != nil {
			panic("could not start transaction")
		}
		// set new tx
		tx.tx = newTx
	}

	return tx.tx.Client()
}
