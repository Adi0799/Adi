

package services

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)


func CreateKafkaWriter(broker, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})
}


func SendMessage(writer *kafka.Writer, message string) error {
	
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
