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

// NewRDBGetterRepository implements domain.RDBGetterRepository.
func NewRDBGetterRepository(config *domain.Config) (domain.RDBGetterRepository, func(), error) {
	l := logger.NewLogger("rdb")
	l.Info("rdb initializing...")

	datasource := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
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
		return nil, f, err
	}

	c, cancel := context.WithTimeout(context.Background(), config.GetRDBConnectionTimeout())
	defer cancel()

	if err := db.PingContext(c); err != nil {
		return nil, f, err
	}

	l.Info(
		"rdb initialized",
		zap.String("user", config.GetRDBUser()),
		zap.String("db_name", config.GetRDBName()),
		zap.String("address", config.GetRDBURL()),
	)

	return &rdb{
		db:     db,
		logger: l,
	}, f, nil
}

func (db *rdb) GetRDBOperator(c context.Context) domain.RDBOperator {
	return &dbOperator{
		db:     db.db,
		logger: db.logger,
	}
}

type dbOperator struct {
	db     *sql.DB
	logger *logger.Logger
}

func (db *dbOperator) Get(
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

func detectMysqlError(err error) error {
	if err == nil {
		return nil
	}

	driverErr, ok := err.(*mysql.MySQLError) // nolint:errorlint
	if !ok {
		return err
	}

	switch driverErr.Number {
	case mysqlerr.ER_DUP_ENTRY:
		return domain.NewConflictError(err.Error())
	default:
		return err
	}
}
