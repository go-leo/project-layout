# project-layout
project-layout是一个go项目的目录结构，结构如下：
- app: 应用程序代码, 具体的代码结构见[app-layout](https://github.com/go-leo/app-layout)
- deployments: 部署相关的配置文件
- docs: 文档相关的配置文件
- githooks: git钩子
- internal: 内部包，不对外暴露
- pkg: 外部包，对外暴露
- scripts: 脚本
- third_party: 第三方依赖
- tools: 工具包

```
.
├── LICENSE
├── Makefile
├── README.md
├── app
│   ├── app-layout
│   └── wire.go
├── deployments
│   └── doc.go
├── docs
│   └── doc.go
├── githooks
│   ├── commit-msg
│   └── pre-commit
├── go.mod
├── go.sum
├── internal
│   └── doc.go
├── pkg
│   ├── actuatorx
│   ├── aliyunx
│   ├── amqpx
│   ├── cachex
│   ├── configx
│   ├── consulx
│   ├── databasex
│   ├── elasticsearchx
│   ├── etcdx
│   ├── gorillax
│   ├── grpcx
│   ├── idx
│   ├── kafkax
│   ├── mongox
│   ├── nacosx
│   ├── otelx
│   ├── redisx
│   ├── registryx
│   └── wire.go
├── scripts
│   ├── app.sh
│   ├── format.sh
│   ├── go_gen.sh
│   ├── lint.sh
│   ├── protoc_gen.sh
│   ├── tools.sh
│   └── wire_gen.sh
├── third_party
│   ├── google
│   ├── leo
│   └── validate
└── tools
    └── tools.go
```