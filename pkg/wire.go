package pkg

import (
	"github.com/go-leo/project-layout/pkg/actuatorx"
	"github.com/go-leo/project-layout/pkg/aliyunx"
	"github.com/go-leo/project-layout/pkg/amqpx"
	"github.com/go-leo/project-layout/pkg/cachex"
	"github.com/go-leo/project-layout/pkg/configx"
	"github.com/go-leo/project-layout/pkg/consulx"
	"github.com/go-leo/project-layout/pkg/databasex"
	"github.com/go-leo/project-layout/pkg/elasticsearchx"
	"github.com/go-leo/project-layout/pkg/ginx"
	"github.com/go-leo/project-layout/pkg/gorillax"
	"github.com/go-leo/project-layout/pkg/grpcx"
	"github.com/go-leo/project-layout/pkg/idx"
	"github.com/go-leo/project-layout/pkg/kafkax"
	"github.com/go-leo/project-layout/pkg/mongox"
	"github.com/go-leo/project-layout/pkg/nacosx"
	"github.com/go-leo/project-layout/pkg/otelx"
	"github.com/go-leo/project-layout/pkg/redisx"
	"github.com/go-leo/project-layout/pkg/registryx"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	actuatorx.Provider,
	aliyunx.Provider,
	amqpx.Provider,
	cachex.Provider,
	configx.Provider,
	consulx.Provider,
	databasex.Provider,
	elasticsearchx.Provider,
	ginx.Provider,
	gorillax.Provider,
	grpcx.Provider,
	idx.Provider,
	kafkax.Provider,
	mongox.Provider,
	nacosx.Provider,
	otelx.Provider,
	redisx.Provider,
	registryx.Provider,
)
