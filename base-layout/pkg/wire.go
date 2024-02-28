package pkg

import (
	"github.com/go-leo/project-layout/ddd-layout/pkg/actuatorx"
	"github.com/go-leo/project-layout/ddd-layout/pkg/amqpx"
	"github.com/go-leo/project-layout/ddd-layout/pkg/configx"
	"github.com/go-leo/project-layout/ddd-layout/pkg/consulx"
	"github.com/go-leo/project-layout/ddd-layout/pkg/databasex"
	"github.com/go-leo/project-layout/ddd-layout/pkg/elasticsearchx"
	"github.com/go-leo/project-layout/ddd-layout/pkg/ginx"
	"github.com/go-leo/project-layout/ddd-layout/pkg/grpcx"
	"github.com/go-leo/project-layout/ddd-layout/pkg/kafkax"
	"github.com/go-leo/project-layout/ddd-layout/pkg/mongox"
	"github.com/go-leo/project-layout/ddd-layout/pkg/nacosx"
	"github.com/go-leo/project-layout/ddd-layout/pkg/otelx"
	"github.com/go-leo/project-layout/ddd-layout/pkg/redisx"
	"github.com/go-leo/project-layout/ddd-layout/pkg/registryx"
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
	mongox.Provider,
	elasticsearchx.Provider,
	consulx.Provider,
)
