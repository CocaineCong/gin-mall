package kafka

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"

	conf "github.com/CocaineCong/gin-mall/config"
	"github.com/CocaineCong/gin-mall/pkg/utils/log"
)

type Kafka struct {
	Key             string
	DisableConsumer bool
	Debug           bool
	Producer        sarama.SyncProducer
	Consumer        sarama.Consumer
	Client          sarama.Client
}

var globalKafkaClient *sync.Map

func GetClient(key string) (*Kafka, error) {
	val, ok := globalKafkaClient.Load(key)
	if !ok {
		return nil, fmt.Errorf("获取kafka client失败，key：%s", key)
	}
	return val.(*Kafka), nil
}

// SendMessage 发送消息
func SendMessage(ctx context.Context, key, topic, value string) error {
	return SendMessagePartitionPar(ctx, key, topic, value, "")
}

func SendMessagePartitionPar(ctx context.Context, key, topic, value, partitionKey string) error {
	kakfa, err := GetClient(key)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.StringEncoder(value),
		Timestamp: time.Now(),
	}

	if partitionKey != "" {
		msg.Key = sarama.StringEncoder(partitionKey)
	}

	partition, offset, err := kakfa.Producer.SendMessage(msg)

	if err != nil {
		return err
	}

	if kakfa.Debug {
		log.LogrusObj.Infoln("发送kafka消息成功",
			zap.Int32("partition", partition),
			zap.Int64("offset", offset))
	}

	return err
}

func InitKafka() {
	for k, v := range conf.Config.KafKa {
		key := k
		val := v
		scfg := buildConfig(val)
		kafka, err := newKafkaClient(key, val, scfg)
		if err != nil {
			return
		}
		globalKafkaClient.Store(key, kafka)
	}
}

func buildConfig(v *conf.KafkaConfig) *sarama.Config {
	cfg := sarama.NewConfig()
	cfg.Producer.RequiredAcks = sarama.RequiredAcks(v.RequiredAck)
	cfg.Producer.Return.Successes = true

	if v.Partition == 1 {
		cfg.Producer.Partitioner = sarama.NewRandomPartitioner
	}

	if v.Partition == 2 {
		cfg.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	}

	if v.ReadTimeout != 0 {
		cfg.Net.ReadTimeout = time.Duration(v.ReadTimeout) * time.Second
	}

	if v.WriteTimeout != 0 {
		cfg.Net.WriteTimeout = time.Duration(v.WriteTimeout) * time.Second
	}

	if v.MaxOpenRequests != 0 {
		cfg.Net.MaxOpenRequests = v.MaxOpenRequests
	}

	return cfg
}

func newKafkaClient(key string, cfgt interface{}, scfg *sarama.Config) (*Kafka, error) {
	cfg := cfgt.(*conf.KafkaConfig)
	client, err := sarama.NewClient(strings.Split(cfg.Address, ","), scfg)
	if err != nil {
		return nil, err
	}

	syncProducer, err := sarama.NewSyncProducer(strings.Split(cfg.Address, ","), scfg)
	if err != nil {
		return nil, err
	}

	consumer, err := sarama.NewConsumer(strings.Split(cfg.Address, ","), scfg)
	if err != nil {
		return nil, err
	}

	return &Kafka{
		Key:             key,
		DisableConsumer: cfg.DisableConsumer,
		Debug:           cfg.Debug,
		Producer:        syncProducer,
		Consumer:        consumer,
		Client:          client,
	}, nil

}

func ConsumerGroup(ctx context.Context, key, groupId, topics string, msgHandler ConsumerGroupHandler) (err error) {
	kafka, err := GetClient(key)
	if err != nil {
		return
	}

	if isConsumerDisabled(kafka) {
		return
	}

	consumerGroup, err := sarama.NewConsumerGroupFromClient(groupId, kafka.Client)
	if err != nil {
		return
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.LogrusObj.Error("消费kafka发生panic", zap.Any("panic", err))
			}
		}()

		defer func() {
			err := consumerGroup.Close()
			if err != nil {
				log.LogrusObj.Error("close err", zap.Any("panic", err))
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := consumerGroup.Consume(ctx, strings.Split(topics, ","), ConsumerGroupHandler(func(msg *sarama.ConsumerMessage) error {
					return middlewareConsumerHandler(msgHandler)(msg)
				}))
				if err != nil {
					log.LogrusObj.Error("消费kafka失败 err", zap.Any("panic", err))

				}
			}
		}

	}()
	return
}

func Consumer(ctx context.Context, key, topic string, fn func(message *sarama.ConsumerMessage) error) (err error) {
	kafka, err := GetClient(key)
	if err != nil {
		return
	}

	partitions, err := kafka.Consumer.Partitions(topic)
	if err != nil {
		return
	}

	for _, partition := range partitions {
		// 针对每个分区创建一个对应的分区消费者
		offset, errx := kafka.Client.GetOffset(topic, partition, sarama.OffsetNewest)
		if errx != nil {
			log.LogrusObj.Infoln("获取Offset失败:", errx)
			continue
		}

		pc, errx := kafka.Consumer.ConsumePartition(topic, partition, offset)
		if errx != nil {
			log.LogrusObj.Infoln("获取Offset失败:", errx)
			return
		}

		// 从每个分区都消费消息
		go func(consumer sarama.PartitionConsumer) {
			defer func() {
				if err := recover(); err != nil {
					log.LogrusObj.Error("消费kafka信息发生panic,err:%s", err)
				}
			}()

			defer func() {
				err := pc.Close()
				if err != nil {
					log.LogrusObj.Infoln("消费kafka信息发生panic,err:%s", err)
				}
			}()

			for {
				select {
				case msg := <-pc.Messages():
					err := middlewareConsumerHandler(fn)(msg)
					if err != nil {
						return
					}
				case <-ctx.Done():
					return
				}
			}

		}(pc)
	}
	return nil
}

func isConsumerDisabled(in *Kafka) bool {
	if in.DisableConsumer {
		log.LogrusObj.Infoln("kafka consumer disabled,key:%s", in.Key)
	}

	return in.DisableConsumer
}
