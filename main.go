package main

import (
    "notification-api/config"
    "notification-api/handlers"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    config.ConnectMySQL()

    router := gin.Default()

    // Health check route
    router.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Server is running"})
    })

    // Define main routes
    router.POST("/subscribe", handlers.Subscribe)
    router.POST("/notifications/send", handlers.SendNotification)
    router.POST("/unsubscribe", handlers.Unsubscribe)
    router.GET("/subscriptions/:user_id", handlers.GetUserSubscriptions)

    log.Println("Starting server on port 8081...")
    router.Run(":8081")
}
