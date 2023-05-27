package kafka

import (
	"github.com/Shopify/sarama"
	"go.uber.org/zap"

	"github.com/CocaineCong/gin-mall/pkg/utils/log"
)

type ConsumerGroupHandler func(message *sarama.ConsumerMessage) error

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if err := h(msg); err != nil {
			log.LogrusObj.Infoln("消息处理失败",
				zap.String("topic", msg.Topic),
				zap.String("value", string(msg.Value)))
			continue
		}
		sess.MarkMessage(msg, "")
	}

	return nil
}

func middlewareConsumerHandler(fn func(message *sarama.ConsumerMessage) error) func(message *sarama.ConsumerMessage) error {
	return func(msg *sarama.ConsumerMessage) error {
		return fn(msg)
	}
}
