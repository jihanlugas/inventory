package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jihanlugas/inventory/config"
	"github.com/jihanlugas/inventory/log"
	"time"
)

var pool *pgxpool.Pool

type ErrAcquireConn struct{}

type CloseConnection func()

func Initialize() *pgxpool.Pool {
	var err error

	dbConnCtx, cancelDbConnCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelDbConnCtx()

	pgxConfig, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DBInfo.Username, config.DBInfo.Password, config.DBInfo.Host, config.DBInfo.Port, config.DBInfo.DbName))
	if err != nil {
		panic(err)
	}

	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// AfterConnect is called after a connection is established, but before it is added to the pool
		// do something with every new connection
		// bisa untuk setup prepared statement atau yang lainnya
		// belum butuh untuk skrng
		return nil
	}

	pgxConfig.ConnConfig.Logger = zerologadapter.NewLogger(log.Sql)

	pool, err = pgxpool.ConnectConfig(dbConnCtx, pgxConfig)
	if err != nil {
		panic(err)
	}

	err = pool.Ping(dbConnCtx)
	if err != nil {
		panic(err)
	}

	return pool
}

func GetConnectionWithDuration(secondDuration int64) (*pgxpool.Conn, context.Context, CloseConnection) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(secondDuration)*time.Second)
	con, err := pool.Acquire(ctx)
	if err != nil {
		panic(&ErrAcquireConn{})
	}

	return con, ctx, closeConn(con, cancelCtx)
}

func closeConn(conn *pgxpool.Conn, cancelCtx context.CancelFunc) CloseConnection {
	return func() {
		conn.Release()
		cancelCtx()
	}
}

func GetConnection() (*pgxpool.Conn, context.Context, CloseConnection) {
	return GetConnectionWithDuration(15)
}

func DeferHandleTransaction(ctx context.Context, tx pgx.Tx) {
	if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
		log.System.Error().Err(err).Send()
	}
}
