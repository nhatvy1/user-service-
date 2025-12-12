package notifier

import "context"

type MessageType string

const (
	MessageTypeOTP     MessageType = "otp"
	MessageTypeWelcome MessageType = "welcome"
)

type NotificationMessage struct {
	Type     MessageType            `json:"type"`
	To       string                 `json:"to"`       // email, phone, user_id
	Template string                 `json:"template"` // optional: template name
	Data     map[string]interface{} `json:"data"`     // flexible data for template
}

type NotificationStrategy interface {
	Send(ctx context.Context, msg NotificationMessage) error
	Name() string // Để log/debug
}
