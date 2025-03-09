package kafkax

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/go-leo/gox/syncx/lazyloadx"
)

func NewConsumerGroups(config *Config) *lazyloadx.Group[sarama.ConsumerGroup] {
	return &lazyloadx.Group[sarama.ConsumerGroup]{
		New: func(key string) (sarama.ConsumerGroup, error) {
			configs := config.GetConfigs()
			options, ok := configs[key]
			if !ok {
				return nil, fmt.Errorf("kafka %s not found", key)
			}

			saramaConfig := sarama.NewConfig()
			if version := options.GetVersion(); version != nil {
				kafkaVersion, err := sarama.ParseKafkaVersion(version.GetValue())
				if err != nil {
					return nil, err
				}
				saramaConfig.Version = kafkaVersion
			}
			if initial := options.GetConsumer().GetOffset().GetInitial(); initial != nil {
				saramaConfig.Consumer.Offsets.Initial = initial.GetValue()
			}
			switch options.GetConsumer().GetGroup().GetRebalance().GetGroupStrategies().GetValue() {
			case "sticky":
				saramaConfig.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
			case "roundrobin":
				saramaConfig.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
			case "range":
				saramaConfig.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
			default:
				saramaConfig.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
			}
			return sarama.NewConsumerGroup(options.GetAddrs(), options.GetGroupId().GetValue(), saramaConfig)
		},
	}
}

func NewSyncProducers(config *Config) *lazyloadx.Group[sarama.SyncProducer] {
	return &lazyloadx.Group[sarama.SyncProducer]{
		New: func(key string) (sarama.SyncProducer, error) {
			configs := config.GetConfigs()
			options, ok := configs[key]
			if !ok {
				return nil, fmt.Errorf("kafka %s not found", key)
			}

			saramaConfig := sarama.NewConfig()
			if version := options.GetVersion(); version != nil {
				kafkaVersion, err := sarama.ParseKafkaVersion(version.GetValue())
				if err != nil {
					return nil, err
				}
				saramaConfig.Version = kafkaVersion
			}
			if producer := options.GetProducer(); producer != nil {
				if requiredAcks := producer.GetRequiredAcks(); requiredAcks != nil {
					saramaConfig.Producer.RequiredAcks = sarama.RequiredAcks(requiredAcks.GetValue())
				}
				if retry := producer.GetRetry(); retry != nil {
					if backoff := retry.GetBackoff(); backoff != nil {
						saramaConfig.Producer.Retry.Backoff = backoff.AsDuration()
					}
					if max := retry.GetMax(); max != nil {
						saramaConfig.Producer.Retry.Max = int(max.GetValue())
					}
				}
				if returns := producer.GetReturn(); returns != nil {
					if errors := returns.GetErrors(); errors != nil {
						saramaConfig.Producer.Return.Errors = errors.GetValue()
					}
					if successes := returns.GetSuccesses(); successes != nil {
						saramaConfig.Producer.Return.Successes = successes.GetValue()
					}
				}
			}
			return sarama.NewSyncProducer(options.Addrs, saramaConfig)
		},
	}
}
