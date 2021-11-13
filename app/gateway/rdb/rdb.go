package rdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/VividCortex/mysqlerr"
	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/logger"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type rdb struct {
	db     *sql.DB
	logger *logger.Logger
}

// NewRDBDetectorRepository implements domain.RDBDetectorRepository.
func NewRDBDetectorRepository(
	config *domain.Config,
) (domain.RDBStatementSetRepository, domain.RDBDetectorRepository, func(), error) {
	l := logger.NewLogger("rdb")
	l.Info("rdb initializing...")

	datasource := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		config.GetRDBUser(),
		config.GetRDBPass(),
		config.GetRDBURL(),
		config.GetRDBName(),
	)

	db, err := sql.Open("mysql", datasource)
	f := func() {
		if db == nil {
			return
		}

		if err := db.Close(); err != nil {
			l.Error("fail to close RDB", zap.Error(err))
		}
	}

	if err != nil {
		return nil, nil, f, err
	}

	c, cancel := context.WithTimeout(context.Background(), config.GetRDBConnectionTimeout())
	defer cancel()

	if err := db.PingContext(c); err != nil {
		return nil, nil, f, err
	}

	l.Info(
		"rdb initialized",
		zap.String("user", config.GetRDBUser()),
		zap.String("db_name", config.GetRDBName()),
		zap.String("address", config.GetRDBURL()),
	)

	rdb := &rdb{
		db:     db,
		logger: l,
	}

	return rdb, rdb, f, nil
}

type dbStructor struct{}

func (db *rdb) Transaction(c context.Context, f func(c context.Context) error) error {
	db.logger.Info("start transaction")

	tx, err := db.db.BeginTx(c, &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false})
	if err != nil {
		return err
	}

	ctxInTx := context.WithValue(c, &dbStructor{}, tx)

	if err := f(ctxInTx); err != nil {
		db.logger.Warn("rollback")

		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			db.logger.Error("fail to rollback", zap.Error(err))
		}

		return err
	}

	if err := tx.Commit(); err != nil {
		db.logger.Error("fail to commit", zap.Error(err))

		return err
	}

	db.logger.Info("committed transaction")

	return nil
}

func (db *rdb) GetRDBOperator(c context.Context) domain.RDBOperator {
	tx, ok := c.Value(&dbStructor{}).(*sql.Tx)
	if ok && tx != nil {
		return &txOperator{
			tx:     tx,
			logger: db.logger,
		}
	}

	return &dbOperator{
		db:     db.db,
		logger: db.logger,
	}
}

type dbOperator struct {
	db     *sql.DB
	logger *logger.Logger
}

func (db *dbOperator) Get( // nolint:dupl
	c context.Context, query string, args []interface{}, dest ...interface{},
) error {
	db.logger.Info(
		"get data",
		zap.String("query", query),
		zap.Any("args", args),
	)

	stmt, err := db.db.PrepareContext(c, query)
	if err != nil {
		return err
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			db.logger.Error(
				"fail to close stmt",
				zap.String("query", query),
				zap.Any("args", args),
				zap.Error(err),
			)
		}
	}()

	row := stmt.QueryRowContext(c, args...)
	if err := row.Err(); err != nil {
		return err
	}

	if err := row.Scan(dest...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewNotFoundError(err.Error())
		}

		return err
	}

	db.logger.Info(
		"got data",
		zap.String("query", query),
		zap.Any("args", args),
	)

	return nil
}

func (db *dbOperator) Select(
	c context.Context, query string, args []interface{}, scanFunc func(*sql.Rows) error,
) error {
	db.logger.Info(
		"select",
		zap.String("query", query),
		zap.Any("args", args),
	)

	stmt, err := db.db.PrepareContext(c, query)
	if err != nil {
		return err
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			db.logger.Error(
				"fail to close stmt",
				zap.String("query", query),
				zap.Any("args", args),
				zap.Error(err),
			)
		}
	}()

	rows, err := stmt.QueryContext(c, args...)
	if err != nil {
		return err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			db.logger.Error(
				"fail to close rows",
				zap.String("query", query),
				zap.Any("args", args),
				zap.Error(err),
			)
		}
	}()

	if err := rows.Err(); err != nil {
		return err
	}

	db.logger.Info(
		"selected",
		zap.String("query", query),
		zap.Any("args", args),
	)

	if err := scanFunc(rows); err != nil {
		return err
	}

	return nil
}

func (db *dbOperator) Exec(
	c context.Context, query string, args ...interface{},
) (sql.Result, error) {
	db.logger.Info(
		"execute",
		zap.String("query", query),
		zap.Any("args", args),
	)

	stmt, err := db.db.PrepareContext(c, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			db.logger.Error(
				"fail to close stmt",
				zap.String("query", query),
				zap.Any("args", args),
				zap.Error(err),
			)
		}
	}()

	result, err := stmt.ExecContext(c, args...)
	if err != nil {
		return nil, detectMysqlError(err)
	}

	db.logger.Info(
		"executed",
		zap.String("query", query),
		zap.Any("args", args),
	)

	return result, nil
}

type txOperator struct {
	tx     *sql.Tx
	logger *logger.Logger
}

func (tx *txOperator) Get( // nolint:dupl
	c context.Context, query string, args []interface{}, dest ...interface{},
) error {
	tx.logger.Info(
		"get data (transaction)",
		zap.String("query", query),
		zap.Any("args", args),
	)

	stmt, err := tx.tx.PrepareContext(c, query)
	if err != nil {
		return err
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			tx.logger.Error(
				"fail to close stmt (transaction)",
				zap.String("query", query),
				zap.Any("args", args),
				zap.Error(err),
			)
		}
	}()

	row := stmt.QueryRowContext(c, args...)
	if err := row.Err(); err != nil {
		return err
	}

	if err := row.Scan(dest...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewNotFoundError(err.Error())
		}

		return err
	}

	tx.logger.Info(
		"got data (transaction)",
		zap.String("query", query),
		zap.Any("args", args),
	)

	return nil
}

func (tx *txOperator) Select(
	c context.Context, query string, args []interface{}, scanFunc func(*sql.Rows) error,
) error {
	tx.logger.Info(
		"select (transaction)",
		zap.String("query", query),
		zap.Any("args", args),
	)

	stmt, err := tx.tx.PrepareContext(c, query)
	if err != nil {
		return err
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			tx.logger.Error(
				"fail to close stmt (transaction)",
				zap.String("query", query),
				zap.Any("args", args),
				zap.Error(err),
			)
		}
	}()

	rows, err := stmt.QueryContext(c, args...)
	if err != nil {
		return err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			tx.logger.Error(
				"fail to close rows (transaction)",
				zap.String("query", query),
				zap.Any("args", args),
				zap.Error(err),
			)
		}
	}()

	if err := rows.Err(); err != nil {
		return err
	}

	tx.logger.Info(
		"selected (transaction)",
		zap.String("query", query),
		zap.Any("args", args),
	)

	if err := scanFunc(rows); err != nil {
		return err
	}

	return nil
}

func (tx *txOperator) Exec(
	c context.Context, query string, args ...interface{},
) (sql.Result, error) {
	tx.logger.Info(
		"execute (transaction)",
		zap.String("query", query),
		zap.Any("args", args),
	)

	stmt, err := tx.tx.PrepareContext(c, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			tx.logger.Error(
				"fail to close stmt (transaction)",
				zap.String("query", query),
				zap.Any("args", args),
				zap.Error(err),
			)
		}
	}()

	result, err := stmt.ExecContext(c, args...)
	if err != nil {
		return nil, detectMysqlError(err)
	}

	tx.logger.Info(
		"executed (transaction)",
		zap.String("query", query),
		zap.Any("args", args),
	)

	return result, nil
}

func detectMysqlError(err error) error {
	if err == nil {
		return nil
	}

	driverErr, ok := err.(*mysql.MySQLError) // nolint:errorlint
	if !ok {
		return err
	}

	switch driverErr.Number {
	// duplicate
	case mysqlerr.ER_DUP_ENTRY:
		return domain.NewConflictError(err.Error())
	// illegal foreign key constraint
	case mysqlerr.ER_ROW_IS_REFERENCED_2, mysqlerr.ER_NO_REFERENCED_ROW_2:
		return domain.NewIllegalForeignKeyConstraintError(err.Error())
	default:
		return err
	}
}
