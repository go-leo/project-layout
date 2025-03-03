package otelx

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/go-leo/gox/stringx"
	prome "github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"time"
)

type MetricConfig struct {
	// PrometheusOptions
	PrometheusOptions *PrometheusMetricReader `mapstructure:"prometheus_options" json:"prometheus_options" yaml:"prometheus_options"`
	// HTTPOptions
	HTTPOptions *HTTPMetricReader `mapstructure:"http_options" json:"http_options" yaml:"http_options"`
	// GRPCOptions
	GRPCOptions *GRPCMetricReader `mapstructure:"grpc_options" json:"grpc_options" yaml:"grpc_options"`
	// WriterOptions
	WriterOptions *WriterMetricReader `mapstructure:"writer_options" json:"writer_options" yaml:"writer_options"`
	// Resources
	ResourceFlag string `mapstructure:"resource_flag" json:"resource_flag" yaml:"resource_flag"`
	// Attributes
	Attributes []attribute.KeyValue
	// ViewOptions
	ViewOptions []ViewOption
}

type ViewOption struct {
	Criteria sdkmetric.Instrument
	Mask     sdkmetric.Stream
}

type GRPCMetricReader struct {
	Endpoint               string            `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	Insecure               bool              `mapstructure:"insecure" json:"insecure" yaml:"insecure"`
	Headers                map[string]string `mapstructure:"headers" json:"headers" yaml:"headers"`
	Compressor             string            `mapstructure:"compressor" json:"compressor" yaml:"compressor"`
	ReconnectionPeriod     time.Duration     `mapstructure:"reconnection_period" json:"reconnection_period" yaml:"reconnection_period"`
	Timeout                time.Duration     `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
	ServiceConfig          string            `mapstructure:"service_config" json:"service_config" yaml:"service_config"`
	PeriodicReaderTimeout  time.Duration     `mapstructure:"periodic_reader_timeout" json:"periodic_reader_timeout" yaml:"periodic_reader_timeout"`
	PeriodicReaderInterval time.Duration     `mapstructure:"periodic_reader_interval" json:"periodic_reader_interval" yaml:"periodic_reader_interval"`
	TLSConfig              *tls.Config
	DialOptions            []grpc.DialOption
	GRPCConn               *grpc.ClientConn
	Retry                  *otlpmetricgrpc.RetryConfig
	TemporalitySelector    sdkmetric.TemporalitySelector
	AggregationSelector    sdkmetric.AggregationSelector
}

type HTTPMetricReader struct {
	Endpoint               string            `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	URLPath                string            `mapstructure:"url_path" json:"url_path" yaml:"url_path"`
	Insecure               bool              `mapstructure:"insecure" json:"insecure" yaml:"insecure"`
	Headers                map[string]string `mapstructure:"headers" json:"headers" yaml:"headers"`
	Timeout                time.Duration     `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
	PeriodicReaderTimeout  time.Duration     `mapstructure:"periodic_reader_timeout" json:"periodic_reader_timeout" yaml:"periodic_reader_timeout"`
	PeriodicReaderInterval time.Duration     `mapstructure:"periodic_reader_interval" json:"periodic_reader_interval" yaml:"periodic_reader_interval"`
	TLSConfig              *tls.Config
	Retry                  *otlpmetrichttp.RetryConfig
	TemporalitySelector    sdkmetric.TemporalitySelector
	AggregationSelector    sdkmetric.AggregationSelector
	Compression            otlpmetrichttp.Compression
}

type PrometheusMetricReader struct {
	WithoutUnits      bool `mapstructure:"without_units" json:"without_units" yaml:"without_units"`
	WithoutTargetInfo bool `mapstructure:"without_target_info" json:"without_target_info" yaml:"without_target_info"`
	WithoutScopeInfo  bool `mapstructure:"without_scope_info" json:"without_scope_info" yaml:"without_scope_info"`
	Registerer        prome.Registerer
	Aggregation       sdkmetric.AggregationSelector
}

type WriterMetricReader struct {
	// Writer 标准输入或者文件
	Writer                 io.Writer
	NewEncoder             func(writer io.Writer) stdoutmetric.Encoder
	TemporalitySelector    sdkmetric.TemporalitySelector
	AggregationSelector    sdkmetric.AggregationSelector
	PeriodicReaderTimeout  time.Duration `mapstructure:"periodic_reader_timeout" json:"periodic_reader_timeout" yaml:"periodic_reader_timeout"`
	PeriodicReaderInterval time.Duration `mapstructure:"periodic_reader_interval" json:"periodic_reader_interval" yaml:"periodic_reader_interval"`
}

func NewMeterProvider(ctx context.Context, o *MetricConfig) (metric.MeterProvider, error) {
	var reader sdkmetric.Reader
	var err error
	switch {
	case o.PrometheusOptions != nil:
		reader, err = prometheusMetricReader(ctx, o.PrometheusOptions)
	case o.GRPCOptions != nil:
		reader, err = grpcMetricReader(ctx, o.GRPCOptions)
	case o.HTTPOptions != nil:
		reader, err = httpMetricReader(ctx, o.HTTPOptions)
	case o.WriterOptions != nil:
		reader, err = writerMetricReader(ctx, o.WriterOptions)
	default:
		return nil, errors.New("not found a metric reader config")
	}
	if err != nil {
		return nil, err
	}
	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(NewResource(ctx, o.ResourceFlag, o.Attributes...)),
		sdkmetric.WithReader(reader),
		sdkmetric.WithView(newView(o.ViewOptions)...),
	)
	otel.SetMeterProvider(meterProvider)
	return meterProvider, nil
}

func prometheusMetricReader(ctx context.Context, o *PrometheusMetricReader) (sdkmetric.Reader, error) {
	var opts []prometheus.Option
	if o.Registerer != nil {
		opts = append(opts, prometheus.WithRegisterer(o.Registerer))
	}
	if o.Aggregation != nil {
		opts = append(opts, prometheus.WithAggregationSelector(o.Aggregation))
	}
	if o.WithoutTargetInfo {
		opts = append(opts, prometheus.WithoutTargetInfo())
	}
	if o.WithoutUnits {
		opts = append(opts, prometheus.WithoutUnits())
	}
	if o.WithoutScopeInfo {
		opts = append(opts, prometheus.WithoutScopeInfo())
	}
	exporter, err := prometheus.New(opts...)
	if err != nil {
		return nil, err
	}
	return exporter, nil
}

func writerMetricReader(ctx context.Context, o *WriterMetricReader) (sdkmetric.Reader, error) {
	var opts []stdoutmetric.Option
	if o.Writer != nil && o.NewEncoder != nil {
		enc := o.NewEncoder(o.Writer)
		opts = append(opts, stdoutmetric.WithEncoder(enc))
	}
	if o.TemporalitySelector != nil {
		opts = append(opts, stdoutmetric.WithTemporalitySelector(o.TemporalitySelector))
	}
	if o.AggregationSelector != nil {
		opts = append(opts, stdoutmetric.WithAggregationSelector(o.AggregationSelector))
	}
	exporter, err := stdoutmetric.New(opts...)
	if err != nil {
		return nil, err
	}
	var prOpts []sdkmetric.PeriodicReaderOption
	if o.PeriodicReaderTimeout > 0 {
		prOpts = append(prOpts, sdkmetric.WithTimeout(o.PeriodicReaderTimeout))
	}
	if o.PeriodicReaderInterval > 0 {
		prOpts = append(prOpts, sdkmetric.WithInterval(o.PeriodicReaderInterval))
	}
	return sdkmetric.NewPeriodicReader(exporter, prOpts...), nil
}

func newView(viewOptions []ViewOption) []sdkmetric.View {
	var views []sdkmetric.View
	for _, option := range viewOptions {
		views = append(views, sdkmetric.NewView(option.Criteria, option.Mask))
	}
	return views
}

func grpcMetricReader(ctx context.Context, o *GRPCMetricReader) (sdkmetric.Reader, error) {
	var opts []otlpmetricgrpc.Option
	if stringx.IsNotBlank(o.Endpoint) {
		opts = append(opts, otlpmetricgrpc.WithEndpoint(o.Endpoint))
	}
	if o.Insecure {
		opts = append(opts, otlpmetricgrpc.WithInsecure())
	}
	if o.TLSConfig != nil {
		opts = append(opts, otlpmetricgrpc.WithTLSCredentials(credentials.NewTLS(o.TLSConfig)))
	}
	if len(o.Headers) > 0 {
		opts = append(opts, otlpmetricgrpc.WithHeaders(o.Headers))
	}
	if stringx.IsNotBlank(o.Compressor) {
		opts = append(opts, otlpmetricgrpc.WithCompressor(o.Compressor))
	}
	if len(o.DialOptions) > 0 {
		opts = append(opts, otlpmetricgrpc.WithDialOption(o.DialOptions...))
	}
	if o.GRPCConn != nil {
		opts = append(opts, otlpmetricgrpc.WithGRPCConn(o.GRPCConn))
	}
	if o.ReconnectionPeriod > 0 {
		opts = append(opts, otlpmetricgrpc.WithReconnectionPeriod(o.ReconnectionPeriod))
	}
	if o.Retry != nil {
		opts = append(opts, otlpmetricgrpc.WithRetry(*o.Retry))
	}
	if o.Timeout > 0 {
		opts = append(opts, otlpmetricgrpc.WithTimeout(o.Timeout))
	}
	if stringx.IsNotBlank(o.ServiceConfig) {
		opts = append(opts, otlpmetricgrpc.WithServiceConfig(o.ServiceConfig))
	}
	if o.TemporalitySelector != nil {
		opts = append(opts, otlpmetricgrpc.WithTemporalitySelector(o.TemporalitySelector))
	}
	if o.AggregationSelector != nil {
		opts = append(opts, otlpmetricgrpc.WithAggregationSelector(o.AggregationSelector))
	}
	exporter, err := otlpmetricgrpc.New(ctx, opts...)
	if err != nil {
		return nil, err
	}
	var prOpts []sdkmetric.PeriodicReaderOption
	if o.PeriodicReaderTimeout > 0 {
		prOpts = append(prOpts, sdkmetric.WithTimeout(o.PeriodicReaderTimeout))
	}
	if o.PeriodicReaderInterval > 0 {
		prOpts = append(prOpts, sdkmetric.WithInterval(o.PeriodicReaderInterval))
	}
	return sdkmetric.NewPeriodicReader(exporter, prOpts...), nil
}

func httpMetricReader(ctx context.Context, o *HTTPMetricReader) (sdkmetric.Reader, error) {
	var opts []otlpmetrichttp.Option
	if stringx.IsNotBlank(o.Endpoint) {
		opts = append(opts, otlpmetrichttp.WithEndpoint(o.Endpoint))
	}
	if stringx.IsNotBlank(o.URLPath) {
		opts = append(opts, otlpmetrichttp.WithURLPath(o.URLPath))
	}
	if o.Compression > 0 {
		opts = append(opts, otlpmetrichttp.WithCompression(o.Compression))
	}
	if o.TLSConfig != nil {
		opts = append(opts, otlpmetrichttp.WithTLSClientConfig(o.TLSConfig))
	}
	if o.Insecure {
		opts = append(opts, otlpmetrichttp.WithInsecure())
	}
	if len(o.Headers) > 0 {
		opts = append(opts, otlpmetrichttp.WithHeaders(o.Headers))
	}
	if o.Timeout > 0 {
		opts = append(opts, otlpmetrichttp.WithTimeout(o.Timeout))
	}
	if o.Retry != nil {
		opts = append(opts, otlpmetrichttp.WithRetry(*o.Retry))
	}
	if o.TemporalitySelector != nil {
		opts = append(opts, otlpmetrichttp.WithTemporalitySelector(o.TemporalitySelector))
	}
	if o.AggregationSelector != nil {
		opts = append(opts, otlpmetrichttp.WithAggregationSelector(o.AggregationSelector))
	}
	exporter, err := otlpmetrichttp.New(ctx, opts...)
	if err != nil {
		return nil, err
	}
	var prOpts []sdkmetric.PeriodicReaderOption
	if o.PeriodicReaderTimeout > 0 {
		prOpts = append(prOpts, sdkmetric.WithTimeout(o.PeriodicReaderTimeout))
	}
	if o.PeriodicReaderInterval > 0 {
		prOpts = append(prOpts, sdkmetric.WithInterval(o.PeriodicReaderInterval))
	}
	return sdkmetric.NewPeriodicReader(exporter, prOpts...), nil
}
