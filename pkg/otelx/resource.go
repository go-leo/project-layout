package otelx

import (
	"context"
	"os"
	"strings"

	"github.com/go-leo/gox/stringx"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

func Attributes() []attribute.KeyValue {
	var attrs []attribute.KeyValue
	if stringx.IsNotBlank(os.Getenv("LEO_SERVICE_NAME")) {
		attrs = append(attrs, attribute.Key("service.name").String(os.Getenv("LEO_SERVICE_NAME")))
	}
	if stringx.IsNotBlank(os.Getenv("LEO_SERVICE_NAMESPACE")) {
		attrs = append(attrs, attribute.Key("service.namespace").String(os.Getenv("LEO_SERVICE_NAMESPACE")))
	}
	if stringx.IsNotBlank(os.Getenv("LEO_SERVICE_ID")) {
		attrs = append(attrs, attribute.Key("service.instance.id").String(os.Getenv("LEO_SERVICE_ID")))
	}
	if stringx.IsNotBlank(os.Getenv("LEO_SERVICE_VERSION")) {
		attrs = append(attrs, attribute.Key("service.version").String(os.Getenv("LEO_SERVICE_VERSION")))
	}
	return attrs
}

func NewResource(ctx context.Context, res string, attrs ...attribute.KeyValue) *resource.Resource {
	opts := []resource.Option{resource.WithAttributes(append(attrs, Attributes()...)...)}
	upperResFlag := strings.ToLower(res)
	if strings.Contains(upperResFlag, "env") {
		opts = append(opts, resource.WithFromEnv())
	}
	if strings.Contains(upperResFlag, "host") {
		opts = append(opts, resource.WithHost(), resource.WithHostID())
	}
	if strings.Contains(upperResFlag, "telemetry_sdk") {
		opts = append(opts, resource.WithTelemetrySDK())
	}
	if strings.Contains(upperResFlag, "os") {
		opts = append(opts, resource.WithOS())
	}
	if strings.Contains(upperResFlag, "process") {
		opts = append(opts, resource.WithProcess())
	}
	if strings.Contains(upperResFlag, "container") {
		opts = append(opts, resource.WithContainer())
	}
	attributes, err := resource.New(ctx, opts...)
	if err != nil {
		return resource.Default()
	}
	return attributes
}
