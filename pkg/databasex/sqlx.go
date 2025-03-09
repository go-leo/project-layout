package databasex

import (
	"context"
	"fmt"
	"github.com/go-leo/gox/syncx/lazyloadx"
	"github.com/jmoiron/sqlx"
)

func NewSqlxDBs(ctx context.Context, config *Config) *lazyloadx.Group[*sqlx.DB] {
	return &lazyloadx.Group[*sqlx.DB]{
		New: func(key string) (*sqlx.DB, error) {
			configs := config.GetConfigs()
			options, ok := configs[key]
			if !ok {
				return nil, fmt.Errorf("database %s not found", key)
			}
			return NewSqlxDB(ctx, options)
		},
	}
}

func NewSqlxDB(ctx context.Context, options *Options) (*sqlx.DB, error) {
	db, err := NewDB(ctx, options)
	if err != nil {
		return nil, err
	}
	return sqlx.NewDb(db, options.GetDriverName().GetValue()), nil
}
