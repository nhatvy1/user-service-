package initialize

import (
	"context"
	"fmt"
	"time"
	"user-service/pkg/globals"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

func NewDatabase() (*Database, error) {
	dbConnectionStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		globals.Config.Database.User,
		globals.Config.Database.Password,
		globals.Config.Database.Host,
		globals.Config.Database.Port,
		globals.Config.Database.DbName,
	)

	poolCfg, err := pgxpool.ParseConfig(dbConnectionStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	poolCfg.MinConns = int32(globals.Config.Database.MaxIdle)
	poolCfg.MaxConns = int32(globals.Config.Database.MaxOpen)
	if globals.Config.Database.MaxLife > 0 {
		poolCfg.MaxConnLifetime = time.Duration(globals.Config.Database.MaxLife) * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// log.Println("DATABASE CONNECTION ESTABLISHED")

	return &Database{Pool: pool}, nil
}

func (db *Database) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}
