package otelx

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	"runtime"
	"time"

	"github.com/go-leo/gox/stringx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type TraceConfig struct {
	// SamplingRate 采样率
	SamplingRate float64 `mapstructure:"sampling_rate" json:"sampling_rate" yaml:"sampling_rate"`
	// ZipkinOptions
	ZipkinOptions *ZipkinTraceExporter `mapstructure:"zipkin_options" json:"zipkin_options" yaml:"zipkin_options"`
	// WriterOptions
	WriterOptions *WriterTraceExporter `mapstructure:"writer_options" json:"writer_options" yaml:"writer_options"`
	// GRPCOptions
	GRPCOptions *GRPCTraceExporter `mapstructure:"grpc_options" json:"grpc_options" yaml:"grpc_options"`
	// HTTPOptions
	HTTPOptions *HTTPTraceExporter `mapstructure:"http_options" json:"http_options" yaml:"http_options"`
	// ResourceFlag 资源
	ResourceFlag string `mapstructure:"resource_flag" json:"resource_flag" yaml:"resource_flag"`
	// Sampler 自定义Sampler
	Sampler sdktrace.Sampler
	// IDGenerator 自定义id生成器
	IDGenerator sdktrace.IDGenerator
	// SpanProcessor 自定义span处理器
	SpanProcessor sdktrace.SpanProcessor
	// RawSpanLimits
	RawSpanLimits *sdktrace.SpanLimits
	// Attributes trace需要一些额外的信息
	Attributes  []attribute.KeyValue
	Propagators []propagation.TextMapPropagator
}

type GRPCTraceExporter struct {
	Endpoint           string            `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	Insecure           bool              `mapstructure:"insecure" json:"insecure" yaml:"insecure"`
	Headers            map[string]string `mapstructure:"headers" json:"headers" yaml:"headers"`
	Compressor         string            `mapstructure:"compressor" json:"compressor" yaml:"compressor"`
	ReconnectionPeriod time.Duration     `mapstructure:"reconnection_period" json:"reconnection_period" yaml:"reconnection_period"`
	Timeout            time.Duration     `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
	ServiceConfig      string            `mapstructure:"service_config" json:"service_config" yaml:"service_config"`
	TLSConfig          *tls.Config
	DialOptions        []grpc.DialOption
	GRPCConn           *grpc.ClientConn
	Retry              *otlptracegrpc.RetryConfig
}

type HTTPTraceExporter struct {
	Endpoint    string            `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	Insecure    bool              `mapstructure:"insecure" json:"insecure" yaml:"insecure"`
	Headers     map[string]string `mapstructure:"headers" json:"headers" yaml:"headers"`
	Timeout     time.Duration     `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
	URLPath     string            `mapstructure:"url_path" json:"url_path" yaml:"url_path"`
	TLSConfig   *tls.Config
	Compression otlptracehttp.Compression
	Retry       *otlptracehttp.RetryConfig
}

type JaegerTraceExporter struct {
	Endpoint string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

type WriterTraceExporter struct {
	// Writer 标准输入或者文件
	Writer            io.Writer
	PrettyPrint       bool `mapstructure:"pretty_print" json:"pretty_print" yaml:"pretty_print"`
	WithoutTimestamps bool `mapstructure:"without_timestamps" json:"without_timestamps" yaml:"without_timestamps"`
}

type ZipkinTraceExporter struct {
	URL string `mapstructure:"url" json:"url" yaml:"url"`
}

func NewTracerProvider(ctx context.Context, o *TraceConfig) (trace.TracerProvider, error) {
	var exporter sdktrace.SpanExporter
	var err error
	switch {
	case o.ZipkinOptions != nil:
		exporter, err = zipkinTraceExporter(ctx, o.ZipkinOptions)
	case o.HTTPOptions != nil:
		exporter, err = httpTraceExporter(ctx, o.HTTPOptions)
	case o.GRPCOptions != nil:
		exporter, err = grpcTraceExporter(ctx, o.GRPCOptions)
	case o.WriterOptions != nil:
		exporter, err = writerTraceExporter(ctx, o.WriterOptions)
	default:
		return nil, errors.New("not found a trace ExporterProvider")
	}
	if err != nil {
		return nil, err
	}

	var bcOpts []sdktrace.BatchSpanProcessorOption

	tpOpts := []sdktrace.TracerProviderOption{
		sdktrace.WithBatcher(exporter, bcOpts...),
		sdktrace.WithResource(NewResource(ctx, o.ResourceFlag, o.Attributes...)),
	}
	if o.Sampler != nil {
		tpOpts = append(tpOpts, sdktrace.WithSampler(o.Sampler))
	} else {
		tpOpts = append(tpOpts, sdktrace.WithSampler(newSampler(o.SamplingRate)))
	}

	if o.IDGenerator != nil {
		tpOpts = append(tpOpts, sdktrace.WithIDGenerator(o.IDGenerator))
	}
	if o.SpanProcessor != nil {
		tpOpts = append(tpOpts, sdktrace.WithSpanProcessor(o.SpanProcessor))
	}
	if o.RawSpanLimits != nil {
		tpOpts = append(tpOpts, sdktrace.WithRawSpanLimits(*o.RawSpanLimits))
	}
	tracerProvider := sdktrace.NewTracerProvider(tpOpts...)
	otel.SetTracerProvider(tracerProvider)

	if o.Propagators == nil {
		o.Propagators = []propagation.TextMapPropagator{propagation.Baggage{}, propagation.TraceContext{}}
	}
	propagator := propagation.NewCompositeTextMapPropagator(o.Propagators...)
	otel.SetTextMapPropagator(propagator)
	return tracerProvider, nil
}

func zipkinTraceExporter(ctx context.Context, o *ZipkinTraceExporter) (sdktrace.SpanExporter, error) {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConnsPerHost = runtime.GOMAXPROCS(0) + 1
	return zipkin.New(o.URL, zipkin.WithClient(&http.Client{Transport: transport}))
}

func grpcTraceExporter(ctx context.Context, o *GRPCTraceExporter) (sdktrace.SpanExporter, error) {
	var grpcOpts []otlptracegrpc.Option
	if stringx.IsNotBlank(o.Endpoint) {
		grpcOpts = append(grpcOpts, otlptracegrpc.WithEndpoint(o.Endpoint))
	}
	if o.Insecure {
		grpcOpts = append(grpcOpts, otlptracegrpc.WithInsecure())
	}
	if o.TLSConfig != nil {
		grpcOpts = append(grpcOpts, otlptracegrpc.WithTLSCredentials(credentials.NewTLS(o.TLSConfig)))
	}
	if len(o.Headers) > 0 {
		grpcOpts = append(grpcOpts, otlptracegrpc.WithHeaders(o.Headers))
	}
	if stringx.IsNotBlank(o.Compressor) {
		grpcOpts = append(grpcOpts, otlptracegrpc.WithCompressor(o.Compressor))
	}
	if len(o.DialOptions) > 0 {
		grpcOpts = append(grpcOpts, otlptracegrpc.WithDialOption(o.DialOptions...))
	}
	if o.GRPCConn != nil {
		grpcOpts = append(grpcOpts, otlptracegrpc.WithGRPCConn(o.GRPCConn))
	}
	if o.ReconnectionPeriod > 0 {
		grpcOpts = append(grpcOpts, otlptracegrpc.WithReconnectionPeriod(o.ReconnectionPeriod))
	}
	if o.Retry != nil {
		grpcOpts = append(grpcOpts, otlptracegrpc.WithRetry(*o.Retry))
	}
	if o.Timeout > 0 {
		grpcOpts = append(grpcOpts, otlptracegrpc.WithTimeout(o.Timeout))
	}
	if stringx.IsNotBlank(o.ServiceConfig) {
		grpcOpts = append(grpcOpts, otlptracegrpc.WithServiceConfig(o.ServiceConfig))
	}
	return otlptracegrpc.New(ctx, grpcOpts...)
}

func httpTraceExporter(ctx context.Context, o *HTTPTraceExporter) (sdktrace.SpanExporter, error) {
	var opts []otlptracehttp.Option
	if stringx.IsNotBlank(o.Endpoint) {
		opts = append(opts, otlptracehttp.WithEndpoint(o.Endpoint))
	}
	if o.Insecure {
		opts = append(opts, otlptracehttp.WithInsecure())
	}
	if o.TLSConfig != nil {
		opts = append(opts, otlptracehttp.WithTLSClientConfig(o.TLSConfig))
	}
	if len(o.Headers) > 0 {
		opts = append(opts, otlptracehttp.WithHeaders(o.Headers))
	}
	if o.Compression > 0 {
		opts = append(opts, otlptracehttp.WithCompression(o.Compression))
	}
	if o.Retry != nil {
		opts = append(opts, otlptracehttp.WithRetry(*o.Retry))
	}
	if o.Timeout > 0 {
		opts = append(opts, otlptracehttp.WithTimeout(o.Timeout))
	}
	if stringx.IsNotBlank(o.URLPath) {
		opts = append(opts, otlptracehttp.WithURLPath(o.URLPath))
	}
	return otlptracehttp.New(ctx, opts...)
}

func writerTraceExporter(ctx context.Context, o *WriterTraceExporter) (sdktrace.SpanExporter, error) {
	opts := []stdouttrace.Option{
		stdouttrace.WithWriter(o.Writer),
	}
	if o.PrettyPrint {
		opts = append(opts, stdouttrace.WithPrettyPrint())
	}
	if o.WithoutTimestamps {
		opts = append(opts, stdouttrace.WithoutTimestamps())
	}
	return stdouttrace.New(opts...)
}

func newSampler(samplingRate float64) sdktrace.Sampler {
	var sampler sdktrace.Sampler
	switch {
	case samplingRate >= 1:
		sampler = sdktrace.AlwaysSample()
	case samplingRate < 0:
		sampler = sdktrace.NeverSample()
	default:
		sampler = sdktrace.TraceIDRatioBased(samplingRate)
	}
	return sampler
}
