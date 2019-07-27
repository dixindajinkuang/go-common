package mysql

import (
	"context"
	"database/sql"
	"github.com/dajinkuang/elog"
	"time"
)

func DBExecContext(ctx context.Context, __DB *sql.DB, query string, args ...interface{}) (result sql.Result, err error) {
	start := time.Now()
	defer func() {
		dur := int64(time.Since(start) / time.Millisecond)
		if err != nil {
			elog.Error(ctx, "mysql-DBExecContext-error", "query", query, "args", args, "error", err, "dur/ms", dur)
		} else {
			elog.Info(ctx, "mysql-DBExecContext-info", "query", query, "args", args, "dur/ms", dur)
		}
	}()
	result, err = __DB.ExecContext(ctx, query, args...)
	return
}

func DBQueryContext(ctx context.Context, __DB *sql.DB, query string, args ...interface{}) (rows *sql.Rows, err error) {
	start := time.Now()
	defer func() {
		dur := int64(time.Since(start) / time.Millisecond)
		if err != nil {
			elog.Error(ctx, "mysql-DBQueryContext-error", "query", query, "args", args, "error", err, "dur/ms", dur)
		} else {
			elog.Info(ctx, "mysql-DBQueryContext-info", "query", query, "args", args, "dur/ms", dur)
		}
	}()
	rows, err = __DB.QueryContext(ctx, query, args...)
	return
}

func DBQueryRowContext(ctx context.Context, __DB *sql.DB, query string, args ...interface{}) (row *sql.Row) {
	start := time.Now()
	defer func() {
		dur := int64(time.Since(start) / time.Millisecond)
		elog.Info(ctx, "mysql-DBQueryRowContext-info", "query", query, "args", args, "dur/ms", dur)
	}()
	row = __DB.QueryRowContext(ctx, query, args...)
	return
}

func TxExecContext(ctx context.Context, __Tx *sql.Tx, query string, args ...interface{}) (result sql.Result, err error) {
	start := time.Now()
	defer func() {
		dur := int64(time.Since(start) / time.Millisecond)
		if err != nil {
			elog.Error(ctx, "mysql-TxExecContext-error", "query", query, "args", args, "error", err, "dur/ms", dur)
		} else {
			elog.Info(ctx, "mysql-TxExecContext-info", "query", query, "args", args, "dur/ms", dur)
		}
	}()
	result, err = __Tx.ExecContext(ctx, query, args...)
	return
}

func TxQueryContext(ctx context.Context, __Tx *sql.Tx, query string, args ...interface{}) (rows *sql.Rows, err error) {
	start := time.Now()
	defer func() {
		dur := int64(time.Since(start) / time.Millisecond)
		if err != nil {
			elog.Error(ctx, "mysql-TxQueryContext-error", "query", query, "args", args, "error", err, "dur/ms", dur)
		} else {
			elog.Info(ctx, "mysql-TxQueryContext-info", "query", query, "args", args, "dur/ms", dur)
		}
	}()
	rows, err = __Tx.QueryContext(ctx, query, args...)
	return
}

func TxQueryRowContext(ctx context.Context, __Tx *sql.Tx, query string, args ...interface{}) (row *sql.Row) {
	start := time.Now()
	defer func() {
		dur := int64(time.Since(start) / time.Millisecond)
		elog.Info(ctx, "mysql-TxQueryRowContext-info", "query", query, "args", args, "dur/ms", dur)
	}()
	row = __Tx.QueryRowContext(ctx, query, args...)
	return
}
