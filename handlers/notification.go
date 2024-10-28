// handlers/notification.go

package handlers

import (
	"log"
	"notification-api/config"
	"notification-api/models"
	"notification-api/services"
	"strings"

	"github.com/gin-gonic/gin"
)

// Subscribe allows users to subscribe to specific topics
func Subscribe(c *gin.Context) {
	var sub models.Subscription
	if err := c.BindJSON(&sub); err != nil {
		log.Printf("JSON binding error: %v", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	topics := strings.Join(sub.Topics, ",")
	query := `INSERT INTO subscriptions (user_id, topics, email, sms, push_notifications) VALUES (?, ?, ?, ?, ?)`
	_, err := config.DB.Exec(query, sub.UserID, topics, sub.Channels.Email, sub.Channels.SMS, sub.Channels.PushNotifications)
	if err != nil {
		log.Printf("Failed to save subscription: %v", err)
		c.JSON(500, gin.H{"error": "Failed to save subscription"})
		return
	}

	c.JSON(200, gin.H{"message": "Subscription successful"})
}

// SendNotification sends a notification to all subscribers of a specific topic
func SendNotification(c *gin.Context) {
	// Define the expected structure for the JSON payload
	var payload struct {
		Topic string `json:"topic"`
		Event struct {
			EventID   string `json:"event_id"`
			Timestamp string `json:"timestamp"`
			Details   struct {
				UserID   string `json:"user_id"`
				Email    string `json:"email"`
				Username string `json:"username"`
			} `json:"details"`
		} `json:"event"`
		Message struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		} `json:"message"`
	}

	// Bind JSON payload to the struct
	if err := c.BindJSON(&payload); err != nil {
		log.Printf("JSON binding error: %v", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Simulate sending message to Kafka (or any notification system)
	writer := services.CreateKafkaWriter("localhost:9092", payload.Topic)
	message := payload.Message.Title + ": " + payload.Message.Body
	if err := services.SendMessage(writer, message); err != nil {
		c.JSON(500, gin.H{"error": "Failed to send notification"})
		return
	}

	c.JSON(200, gin.H{"message": "Notification sent"})
}

// Unsubscribe allows users to unsubscribe from specific topics
func Unsubscribe(c *gin.Context) {
	var payload struct {
		UserID string   `json:"user_id"`
		Topics []string `json:"topics"`
	}

	if err := c.BindJSON(&payload); err != nil {
		log.Printf("JSON binding error: %v", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	topics := strings.Join(payload.Topics, ",")
	query := `DELETE FROM subscriptions WHERE user_id = ? AND topics = ?`
	_, err := config.DB.Exec(query, payload.UserID, topics)
	if err != nil {
		log.Printf("Failed to unsubscribe: %v", err)
		c.JSON(500, gin.H{"error": "Failed to unsubscribe"})
		return
	}

	c.JSON(200, gin.H{"message": "Unsubscribed successfully"})
}

// GetUserSubscriptions fetches all topics a user is subscribed to
func GetUserSubscriptions(c *gin.Context) {
	userID := c.Param("user_id")
	query := `SELECT topics, email, sms, push_notifications FROM subscriptions WHERE user_id = ?`
	rows, err := config.DB.Query(query, userID)
	if err != nil {
		log.Printf("Failed to fetch subscriptions: %v", err)
		c.JSON(500, gin.H{"error": "Failed to fetch subscriptions"})
		return
	}
	defer rows.Close()

	var subscriptions []models.Subscription
	for rows.Next() {
		var sub models.Subscription
		var topics string
		if err := rows.Scan(&topics, &sub.Channels.Email, &sub.Channels.SMS, &sub.Channels.PushNotifications); err != nil {
			log.Printf("Error reading data: %v", err)
			c.JSON(500, gin.H{"error": "Error reading data"})
			return
		}

		sub.Topics = strings.Split(topics, ",")
		subscriptions = append(subscriptions, sub)
	}

	c.JSON(200, gin.H{"subscriptions": subscriptions})
}
