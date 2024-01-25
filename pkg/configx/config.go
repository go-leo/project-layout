package configx

import (
	"github.com/go-leo/ddd-layout/pkg/actuatorx"
	"github.com/go-leo/ddd-layout/pkg/amqpx"
	"github.com/go-leo/ddd-layout/pkg/databasex"
	"github.com/go-leo/ddd-layout/pkg/ginx"
	"github.com/go-leo/ddd-layout/pkg/grpcx"
	"github.com/go-leo/ddd-layout/pkg/kafkax"
	"github.com/go-leo/ddd-layout/pkg/nacosx"
	"github.com/go-leo/ddd-layout/pkg/otelx"
	"github.com/go-leo/ddd-layout/pkg/redisx"
)

var AppConf Configuration

type EnvConfig struct {
	Name      string `mapstructure:"SERVICE_NAME" json:"SERVICE_NAME" yaml:"SERVICE_NAME"`
	Namespace string `mapstructure:"SERVICE_NAMESPACE" json:"SERVICE_NAMESPACE" yaml:"SERVICE_NAMESPACE"`
	ID        string `mapstructure:"SERVICE_ID" json:"SERVICE_ID" yaml:"SERVICE_ID"`
	Version   string `mapstructure:"SERVICE_VERSION" json:"SERVICE_VERSION" yaml:"SERVICE_VERSION"`
}

type Configuration struct {
	Env        *EnvConfig                `mapstructure:"env" json:"env" yaml:"env"`
	Actuator   *actuatorx.Config         `mapstructure:"actuator" json:"actuator" yaml:"actuator"`
	Nacos      map[string]*nacosx.Config `mapstructure:"nacos" json:"nacos" yaml:"nacos"`
	Gin        *ginx.Config              `mapstructure:"gin" json:"gin" yaml:"gin"`
	GrpcServer *grpcx.ServerConfig       `mapstructure:"grpc_server" json:"grpc_server" yaml:"grpc_server"`
	GrpcClient grpcx.ClientConfigs       `mapstructure:"grpc_client" json:"grpc_client" yaml:"grpc_client"`
	Trace      *otelx.TraceConfig        `mapstructure:"trace" json:"trace" yaml:"trace"`
	Metric     *otelx.MetricConfig       `mapstructure:"metric" json:"metric" yaml:"metric"`
	AMQP       amqpx.Configs             `mapstructure:"amqp" json:"amqp" yaml:"amqp"`
	Kafka      kafkax.Configs            `mapstructure:"kafka" json:"kafka" yaml:"kafka"`
	Redis      redisx.Configs            `mapstructure:"redis" json:"redis" yaml:"redis"`
	Database   databasex.Configs         `mapstructure:"database" json:"database" yaml:"database"`
}

func Env() *EnvConfig {
	return AppConf.Env
}

func Nacos() nacosx.Configs {
	return AppConf.Nacos
}

func Gin() *ginx.Config {
	return AppConf.Gin
}

func GrpcServer() *grpcx.ServerConfig {
	return AppConf.GrpcServer
}

func GrpcClient() grpcx.ClientConfigs {
	return AppConf.GrpcClient
}

func Actuator() *actuatorx.Config {
	return AppConf.Actuator
}

func Kafka() kafkax.Configs {
	return AppConf.Kafka
}

func AMQP() amqpx.Configs {
	return AppConf.AMQP
}

func Trace() *otelx.TraceConfig {
	return AppConf.Trace
}

func Metric() *otelx.MetricConfig {
	return AppConf.Metric
}

func Redis() redisx.Configs {
	return AppConf.Redis
}

func Database() databasex.Configs {
	return AppConf.Database
}
