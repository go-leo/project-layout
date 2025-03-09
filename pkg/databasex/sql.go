package databasex

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-leo/gox/syncx/lazyloadx"
)

func NewDBs(ctx context.Context, config *Config) *lazyloadx.Group[*sql.DB] {
	return &lazyloadx.Group[*sql.DB]{
		New: func(key string) (*sql.DB, error) {
			configs := config.GetConfigs()
			options, ok := configs[key]
			if !ok {
				return nil, fmt.Errorf("database %s not found", key)
			}
			return NewDB(ctx, options)
		},
	}
}

func NewDB(ctx context.Context, options *Options) (*sql.DB, error) {
	db, err := sql.Open(options.GetDriverName().GetValue(), options.GetDsn().GetValue())
	if err != nil {
		return nil, err
	}
	if options.GetPingTimeout() != nil {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(ctx, options.GetPingTimeout().AsDuration())
		defer cancelFunc()
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	if options.GetMaxIdleConns() != nil {
		db.SetMaxIdleConns(int(options.GetMaxIdleConns().GetValue()))
	}
	if options.GetMaxOpenConns() != nil {
		db.SetMaxOpenConns(int(options.GetMaxOpenConns().GetValue()))
	}
	if options.GetConnMaxLifetime() != nil {
		db.SetConnMaxLifetime(options.GetConnMaxLifetime().AsDuration())
	}
	if options.GetConnMaxIdleTime() != nil {
		db.SetConnMaxIdleTime(options.GetConnMaxIdleTime().AsDuration())
	}
	return db, nil
}
