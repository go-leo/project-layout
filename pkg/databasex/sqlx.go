package databasex

import (
	"context"
	"github.com/jmoiron/sqlx"
)

func NewSqlxDBs(ctx context.Context, configs Configs) (map[string]*sqlx.DB, error) {
	dbs := make(map[string]*sqlx.DB)
	for key, config := range configs {
		db, err := NewSqlxDB(ctx, config)
		if err != nil {
			return nil, err
		}
		dbs[key] = db
	}
	return dbs, nil
}

func NewSqlxDB(ctx context.Context, config *Config) (*sqlx.DB, error) {
	db, err := NewDB(ctx, config)
	if err != nil {
		return nil, err
	}
	return sqlx.NewDb(db, config.DriverName), nil
}
