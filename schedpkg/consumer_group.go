package schedpkg

// ConsumerGroup 定义消费者组
type ConsumerGroup struct {
	Consumers   []IConsumer
	ConsumerNum int
}

type IConsumer interface {
	Start() error
	Stop() error
}

func NewConsumerGroup(consumerFunc func(num int) IConsumer, consumerNum int) *ConsumerGroup {
	cg := &ConsumerGroup{
		Consumers:   make([]IConsumer, 0),
		ConsumerNum: consumerNum,
	}

	for i := 0; i < consumerNum; i++ {
		cg.Consumers = append(cg.Consumers, consumerFunc(i+1))
	}
	return cg
}

func (cg *ConsumerGroup) Start() error {
	for i := 0; i < cg.ConsumerNum; i++ {
		if err := cg.Consumers[i].Start(); err != nil {
			return err
		}
	}
	return nil
}

func (cg *ConsumerGroup) Stop() error {
	for i := 0; i < cg.ConsumerNum; i++ {
		if err := cg.Consumers[i].Stop(); err != nil {
			return err
		}
	}
	return nil
}
