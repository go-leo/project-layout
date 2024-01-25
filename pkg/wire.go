package pkg

import (
	"github.com/go-leo/ddd-layout/pkg/actuatorx"
	"github.com/go-leo/ddd-layout/pkg/amqpx"
	"github.com/go-leo/ddd-layout/pkg/configx"
	"github.com/go-leo/ddd-layout/pkg/databasex"
	"github.com/go-leo/ddd-layout/pkg/ginx"
	"github.com/go-leo/ddd-layout/pkg/grpcx"
	"github.com/go-leo/ddd-layout/pkg/kafkax"
	"github.com/go-leo/ddd-layout/pkg/nacosx"
	"github.com/go-leo/ddd-layout/pkg/otelx"
	"github.com/go-leo/ddd-layout/pkg/redisx"
	"github.com/go-leo/ddd-layout/pkg/registryx"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	actuatorx.Provider,
	amqpx.Provider,
	configx.Provider,
	databasex.Provider,
	ginx.Provider,
	grpcx.Provider,
	kafkax.Provider,
	nacosx.Provider,
	otelx.Provider,
	redisx.Provider,
	registryx.Provider,
)
