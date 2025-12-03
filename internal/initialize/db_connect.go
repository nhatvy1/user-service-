package initialize

import (
	"context"
	"fmt"
	"log"
	"time"
	"user-service/pkg/globals"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DbConnect() (*pgxpool.Pool, error) {
	dbConnectionStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", globals.Config.Database.User, globals.Config.Database.Password, globals.Config.Database.Host, globals.Config.Database.Port, globals.Config.Database.DbName)

	poolCfg, err := pgxpool.ParseConfig(dbConnectionStr)

	if err != nil {
		panic("database connect failed")
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
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	log.Printf("🚀 Successfully connected to PostgreSQL: %s:%d/%s", globals.Config.Database.Host, globals.Config.Database.Port, globals.Config.Database.DbName)

	return pool, nil
}
