package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

//声明交换函数 接收指向amqp通道的指针
func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable? 是否持久
		false,        // auto-deleted? 是否自动删除
		false,        // internal?  是否内部使用
		false,        // no-wait?
		nil,          // arguements?
	)
}

//声明随机队列
func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name?
		false, // durable?
		false, // delete when unused?
		true,  // exclusive? 独家
		false, // no-wait?
		nil,   // arguments?
	)
}
