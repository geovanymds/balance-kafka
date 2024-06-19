package consumer

type ConsumerInterface interface {
	Proccess(msg []byte) error
}

type ConsumerManager struct {
	Consumers map[string]ConsumerInterface
}

func NewConsumerManager() *ConsumerManager {
	return &ConsumerManager{
		Consumers: map[string]ConsumerInterface{},
	}
}

func (cm *ConsumerManager) AddConsumer(topicName string, consumer ConsumerInterface) {
	cm.Consumers[topicName] = consumer
}
