// models/subscription.go

package models

type Subscription struct {
    UserID   string              `json:"user_id"`
    Topics   []string            `json:"topics"`
    Channels NotificationChannels `json:"notification_channels"`
}

type NotificationChannels struct {
    Email             string `json:"email"`
    SMS               string `json:"sms"`
    PushNotifications bool   `json:"push_notifications"`
}
