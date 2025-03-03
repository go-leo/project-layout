package databasex

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewGormDBs(ctx context.Context, configs Configs) (map[string]*gorm.DB, error) {
	dbs := make(map[string]*gorm.DB)
	for key, config := range configs {
		db, err := NewGormDB(ctx, config)
		if err != nil {
			return nil, err
		}
		dbs[key] = db
	}
	return dbs, nil
}

func NewGormDB(ctx context.Context, config *Config) (*gorm.DB, error) {
	var opts []gorm.Option
	open := mysql.Open(config.DSN)
	db, err := gorm.Open(open, opts...)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if config.Timeout <= 0 {
		config.Timeout = 3 * time.Second
	}
	ctx, cancelFunc := context.WithTimeout(ctx, config.Timeout)
	defer cancelFunc()
	if err = sqlDB.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, err
}
