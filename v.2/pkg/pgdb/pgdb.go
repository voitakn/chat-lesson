package pgdb

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type DBTX interface {
	Ping(ctx context.Context) error
	Close() error
}

type Connection struct {
	dbc *pgxpool.Pool
}

var DB Connection

func Service() DBTX {
	return &DB
}

func New(ctx context.Context, DSN string) error {
	dbc, err := pgxpool.New(ctx, DSN)
	if err != nil {
		return errors.Errorf("failed to connect to DB: %v", err)
	}
	if err := dbc.Ping(ctx); err != nil {
		return errors.Errorf("failed to ping DB: %v", err)
	}
	DB = Connection{
		dbc: dbc,
	}
	return nil
}

func (p *Connection) Conn() *pgxpool.Pool {
	return p.dbc
}

func (p *Connection) Ping(ctx context.Context) error {
	fmt.Println("pgdb.DB.Ping")
	return p.dbc.Ping(ctx)
}

func (p *Connection) Close() error {
	fmt.Println("pgdb.DB.Close")
	p.dbc.Close()
	return nil
}
