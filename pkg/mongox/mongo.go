package mongox

import (
	"context"
	"fmt"
	"github.com/go-leo/gox/syncx/lazyloadx"
	"go.mongodb.org/mongo-driver/mongo"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

func NewClients(ctx context.Context, config *Config) *lazyloadx.Group[*mongo.Client] {
	return &lazyloadx.Group[*mongo.Client]{
		New: func(key string) (*mongo.Client, error) {
			configs := config.GetConfigs()
			options, ok := configs[key]
			if !ok {
				return nil, fmt.Errorf("nacos %s not found", key)
			}
			return NewClient(ctx, options)
		},
	}
}

func NewClient(ctx context.Context, options *Options) (*mongo.Client, error) {
	return mongo.Connect(ctx, mongooptions.Client().ApplyURI(options.GetUri().GetValue()))
}
