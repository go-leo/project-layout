package consulx

import (
	"fmt"
	"github.com/go-leo/gox/syncx/lazyloadx"
	"github.com/hashicorp/consul/api"
)

func NewClients(config *Config) *lazyloadx.Group[*api.Client] {
	return &lazyloadx.Group[*api.Client]{
		New: func(key string) (*api.Client, error) {
			configs := config.GetConfigs()
			options, ok := configs[key]
			if !ok {
				return nil, fmt.Errorf("nacos %s not found", key)
			}
			return NewClient(options)
		},
	}
}

func NewClient(options *Options) (*api.Client, error) {
	var httpAuth *api.HttpBasicAuth
	if conf := options.GetHttpAuth(); conf != nil {
		httpAuth = &api.HttpBasicAuth{
			Username: conf.GetUsername().GetValue(),
			Password: conf.GetPassword().GetValue(),
		}
	}
	return api.NewClient(&api.Config{
		Address:    options.GetAddress().GetValue(),
		Scheme:     options.GetScheme().GetValue(),
		PathPrefix: options.GetPathPrefix().GetValue(),
		Datacenter: options.GetDatacenter().GetValue(),
		HttpAuth:   httpAuth,
		WaitTime:   options.GetWaitTime().AsDuration(),
		Token:      options.GetToken().GetValue(),
		TokenFile:  options.GetTokenFile().GetValue(),
		Namespace:  options.GetNamespace().GetValue(),
		Partition:  options.GetPartition().GetValue(),
		TLSConfig: api.TLSConfig{
			Address:            options.GetTlsConfig().GetAddress().GetValue(),
			CAFile:             options.GetTlsConfig().GetCaFile().GetValue(),
			CAPath:             options.GetTlsConfig().GetCaPath().GetValue(),
			CAPem:              options.GetTlsConfig().GetCaPem(),
			CertFile:           options.GetTlsConfig().GetCaFile().GetValue(),
			CertPEM:            options.GetTlsConfig().GetCertPem(),
			KeyFile:            options.GetTlsConfig().GetKeyFile().GetValue(),
			KeyPEM:             options.GetTlsConfig().GetKeyPem(),
			InsecureSkipVerify: options.GetTlsConfig().GetInsecureSkipVerify().GetValue(),
		},
	})
}
