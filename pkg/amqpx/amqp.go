package amqpx

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/rabbitmq/amqp091-go"
	"os"
)

type Configs map[string]*Config

type Config struct {
	URL           string          `mapstructure:"url" json:"url" yaml:"url"`
	TLS           *TLSConfig      `mapstructure:"tls" json:"tls" yaml:"tls"`
	ExternalAuth  bool            `mapstructure:"external_auth" json:"external_auth" yaml:"external_auth"`
	Amqp091Config *amqp091.Config `mapstructure:"amqp_091_config" json:"amqp_091_config" yaml:"amqp_091_config"`
}

type TLSConfig struct {
	PEMFile  string `mapstructure:"pem_file" json:"pem_file" yaml:"pem_file"`
	CertFile string `mapstructure:"cert_file" json:"cert_file" yaml:"cert_file"`
	KeyFile  string `mapstructure:"key_file" json:"key_file" yaml:"key_file"`
}

func NewConnections(configs Configs) (map[string]*amqp091.Connection, error) {
	connections := make(map[string]*amqp091.Connection)
	for key, config := range configs {
		connection, err := NewConnection(config)
		if err != nil {
			return nil, err
		}
		connections[key] = connection
	}
	return connections, nil
}

func NewConnection(config *Config) (*amqp091.Connection, error) {
	if len(config.URL) <= 0 {
		return nil, errors.New("url is empty")
	}
	if config.Amqp091Config != nil {
		return amqp091.DialConfig(config.URL, *config.Amqp091Config)
	}
	if config.TLS == nil {
		return amqp091.Dial(config.URL)
	}
	tlsCfg := new(tls.Config)
	tlsCfg.RootCAs = x509.NewCertPool()
	if ca, err := os.ReadFile(config.TLS.PEMFile); err == nil {
		tlsCfg.RootCAs.AppendCertsFromPEM(ca)
	}
	if cert, err := tls.LoadX509KeyPair(config.TLS.CertFile, config.TLS.KeyFile); err == nil {
		tlsCfg.Certificates = append(tlsCfg.Certificates, cert)
	}
	if config.ExternalAuth {
		return amqp091.DialTLS_ExternalAuth(config.URL, tlsCfg)
	}
	return amqp091.DialTLS(config.URL, tlsCfg)
}
