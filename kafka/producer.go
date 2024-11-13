package kafka

import (
	"github.com/IBM/sarama"
)

func NewSyncProducer(address ...string) (producer sarama.SyncProducer, err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	producer, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		return
	}
	return
}

func NewMessage(topic string, key string, value string) (msg *sarama.ProducerMessage) {
	msg = &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}
	return
}

func SendMessage(producer sarama.SyncProducer, msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	partition, offset, err = producer.SendMessage(msg)
	return
}
