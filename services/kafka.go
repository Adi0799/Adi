// services/kafka.go

package services

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

// CreateKafkaWriter initializes a Kafka writer
func CreateKafkaWriter(broker, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})
}

// SendMessage sends a message to the specified Kafka writer with context.TODO()
func SendMessage(writer *kafka.Writer, message string) error {
	// Use context.TODO() instead of nil to avoid passing a nil context
	err := writer.WriteMessages(context.TODO(), kafka.Message{
		Value: []byte(message),
	})

	if err != nil {
		log.Println("Failed to send message:", err)
		return err
	}

	log.Println("Message sent to Kafka successfully")
	return nil
}
