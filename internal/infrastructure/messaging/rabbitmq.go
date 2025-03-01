package messaging

import (
	"e-commerce/pkg/config"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQClient is a wrapper around the RabbitMQ connection and channel
type RabbitMQClient struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewRabbitMQClient creates a new RabbitMQ client
func NewRabbitMQClient(cfg *config.RabbitMQConfig) (*RabbitMQClient, error) {
	// Connect to RabbitMQ
	conn, err := amqp.Dial(cfg.RabbitMQConnectionString())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	log.Println("Connected to RabbitMQ")
	return &RabbitMQClient{
		Connection: conn,
		Channel:    ch,
	}, nil
}

// Close closes the RabbitMQ connection and channel
func (r *RabbitMQClient) Close() error {
	if r.Channel != nil {
		if err := r.Channel.Close(); err != nil {
			log.Printf("Error closing RabbitMQ channel: %v", err)
		}
	}

	if r.Connection != nil {
		if err := r.Connection.Close(); err != nil {
			return fmt.Errorf("error closing RabbitMQ connection: %w", err)
		}
	}

	log.Println("RabbitMQ connection closed")
	return nil
}

// DeclareExchange declares an exchange
func (r *RabbitMQClient) DeclareExchange(name, kind string, durable, autoDelete, internal, noWait bool) error {
	return r.Channel.ExchangeDeclare(
		name,       // name
		kind,       // type
		durable,    // durable
		autoDelete, // auto-deleted
		internal,   // internal
		noWait,     // no-wait
		nil,        // arguments
	)
}

// DeclareQueue declares a queue
func (r *RabbitMQClient) DeclareQueue(name string, durable, autoDelete, exclusive, noWait bool) (amqp.Queue, error) {
	return r.Channel.QueueDeclare(
		name,       // name
		durable,    // durable
		autoDelete, // delete when unused
		exclusive,  // exclusive
		noWait,     // no-wait
		nil,        // arguments
	)
}

// BindQueue binds a queue to an exchange
func (r *RabbitMQClient) BindQueue(queueName, key, exchangeName string, noWait bool) error {
	return r.Channel.QueueBind(
		queueName,    // queue name
		key,          // routing key
		exchangeName, // exchange
		noWait,       // no-wait
		nil,          // arguments
	)
}

// Publish publishes a message to an exchange
func (r *RabbitMQClient) Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	return r.Channel.PublishWithContext(
		nil,       // context
		exchange,  // exchange
		key,       // routing key
		mandatory, // mandatory
		immediate, // immediate
		msg,       // message
	)
}

// Consume consumes messages from a queue
func (r *RabbitMQClient) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool) (<-chan amqp.Delivery, error) {
	return r.Channel.Consume(
		queue,     // queue
		consumer,  // consumer
		autoAck,   // auto-ack
		exclusive, // exclusive
		noLocal,   // no-local
		noWait,    // no-wait
		nil,       // args
	)
}
