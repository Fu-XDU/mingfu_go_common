package kafka

import (
	"github.com/IBM/sarama"
	"github.com/labstack/gommon/log"
)

type ConsumerGroupHandler struct{}

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		session.MarkMessage(message, "")
	}
	return nil
}

func NewConsumerGroup(groupID string, address ...string) (group sarama.ConsumerGroup, err error) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	group, err = sarama.NewConsumerGroup(address, groupID, config)
	if err != nil {
		log.Errorf("Error creating consumer group: %v", err)
		return
	}
	return
}
