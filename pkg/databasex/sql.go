package databasex

import (
	"context"
	"database/sql"
	"time"
)

func NewDBs(ctx context.Context, configs Configs) (map[string]*sql.DB, error) {
	dbs := make(map[string]*sql.DB)
	for key, config := range configs {
		db, err := NewDB(ctx, config)
		if err != nil {
			return nil, err
		}
		dbs[key] = db
	}
	return dbs, nil
}

func NewDB(ctx context.Context, config *Config) (*sql.DB, error) {
	db, err := sql.Open(config.DriverName, config.DSN)
	if err != nil {
		return nil, err
	}
	if config.Timeout <= 0 {
		config.Timeout = 3 * time.Second
	}
	ctx, cancelFunc := context.WithTimeout(ctx, config.Timeout)
	defer cancelFunc()
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
