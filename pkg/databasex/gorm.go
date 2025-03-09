package databasex

import (
	"context"
	"fmt"
	"github.com/go-leo/gox/syncx/lazyloadx"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDBs(ctx context.Context, config *Config) *lazyloadx.Group[*gorm.DB] {
	return &lazyloadx.Group[*gorm.DB]{
		New: func(key string) (*gorm.DB, error) {
			configs := config.GetConfigs()
			options, ok := configs[key]
			if !ok {
				return nil, fmt.Errorf("database %s not found", key)
			}
			return NewGormDB(ctx, options)
		},
	}
}

func NewGormDB(ctx context.Context, options *Options) (*gorm.DB, error) {
	var opts []gorm.Option
	var dialector gorm.Dialector
	switch options.GetDriverName().GetValue() {
	case "mysql":
		dialector = mysql.Open(options.GetDsn().GetValue())
	case "clickhouse":
		dialector = clickhouse.Open(options.GetDsn().GetValue())
	}
	db, err := gorm.Open(dialector, opts...)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if options.GetPingTimeout() != nil {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(ctx, options.GetPingTimeout().AsDuration())
		defer cancelFunc()
	}
	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, err
	}
	if options.GetMaxIdleConns() != nil {
		sqlDB.SetMaxIdleConns(int(options.GetMaxIdleConns().GetValue()))
	}
	if options.GetMaxOpenConns() != nil {
		sqlDB.SetMaxOpenConns(int(options.GetMaxOpenConns().GetValue()))
	}
	if options.GetConnMaxLifetime() != nil {
		sqlDB.SetConnMaxLifetime(options.GetConnMaxLifetime().AsDuration())
	}
	if options.GetConnMaxIdleTime() != nil {
		sqlDB.SetConnMaxIdleTime(options.GetConnMaxIdleTime().AsDuration())
	}
	return db, err
}
