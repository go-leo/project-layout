package redisx

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-leo/gox/syncx/lazyloadx"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

// New 创建redis客户端集合
func New(ctx context.Context, config *Config) *lazyloadx.Group[redis.UniversalClient] {
	return &lazyloadx.Group[redis.UniversalClient]{
		New: func(key string) (redis.UniversalClient, error) {
			configs := config.GetConfigs()
			options, ok := configs[key]
			if !ok {
				return nil, fmt.Errorf("redis %s not found", key)
			}
			var tlsConfig *tls.Config
			if info := options.GetTlsConfig(); info != nil {
				cert, err := tls.LoadX509KeyPair(info.GetCertFile(), info.GetKeyFile())
				if err != nil {
					return nil, err
				}
				tlsConfig = &tls.Config{Certificates: []tls.Certificate{cert}}
			}
			client := redis.NewUniversalClient(&redis.UniversalOptions{
				Addrs:                 options.Addrs,
				ClientName:            options.GetClientName().GetValue(),
				DB:                    int(options.GetDb().GetValue()),
				Protocol:              int(options.GetProtocol().GetValue()),
				Username:              options.GetUsername().GetValue(),
				Password:              options.GetPassword().GetValue(),
				SentinelUsername:      options.GetSentinelUsername().GetValue(),
				SentinelPassword:      options.GetSentinelPassword().GetValue(),
				MaxRetries:            int(options.GetMaxRetries().GetValue()),
				MinRetryBackoff:       options.GetMinRetryBackoff().AsDuration(),
				MaxRetryBackoff:       options.GetMaxRetryBackoff().AsDuration(),
				DialTimeout:           options.GetDialTimeout().AsDuration(),
				ReadTimeout:           options.GetReadTimeout().AsDuration(),
				WriteTimeout:          options.GetWriteTimeout().AsDuration(),
				ContextTimeoutEnabled: options.GetContextTimeoutEnabled().GetValue(),
				PoolFIFO:              options.GetPoolFifo().GetValue(),
				PoolSize:              int(options.GetPoolSize().GetValue()),
				PoolTimeout:           options.GetPoolTimeout().AsDuration(),
				MinIdleConns:          int(options.GetMinIdleConns().GetValue()),
				MaxIdleConns:          int(options.GetMaxIdleConns().GetValue()),
				MaxActiveConns:        int(options.GetMaxActiveConns().GetValue()),
				ConnMaxIdleTime:       options.GetConnMaxIdleTime().AsDuration(),
				ConnMaxLifetime:       options.GetConnMaxLifetime().AsDuration(),
				TLSConfig:             tlsConfig,
				MaxRedirects:          int(options.GetClusterOptions().GetMaxRedirects().GetValue()),
				ReadOnly:              options.GetClusterOptions().GetReadOnly().GetValue(),
				RouteByLatency:        options.GetClusterOptions().GetRouteByLatency().GetValue(),
				RouteRandomly:         options.GetClusterOptions().GetRouteRandomly().GetValue(),
				MasterName:            options.GetFailoverOptions().GetMasterName().GetValue(),
				DisableIndentity:      options.GetDisableIdentity().GetValue(),
				IdentitySuffix:        options.GetIdentitySuffix().GetValue(),
			})
			_, err := client.Ping(ctx).Result()
			if err != nil {
				return nil, fmt.Errorf("failed ping redis %v, %w", key, err)
			}
			if options.GetEnableTracing().GetValue() {
				if err := redisotel.InstrumentTracing(client); err != nil {
					return nil, err
				}
			}
			if options.GetEnableMetrics().GetValue() {
				if err := redisotel.InstrumentMetrics(client); err != nil {
					return nil, err
				}
			}
			return client, nil
		},
	}
}
