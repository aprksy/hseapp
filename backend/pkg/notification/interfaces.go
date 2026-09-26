// Package notification defines interfaces for notification delivery channels
package notification

import (
	"context"
	"github.com/aprksy/hseapp/backend/internal/domain"
)

// DeliveryChannel represents a notification delivery method
type DeliveryChannel string

const (
	ChannelPush   DeliveryChannel = "push"
	ChannelEmail  DeliveryChannel = "email"
	ChannelSMS    DeliveryChannel = "sms"
	ChannelWhatsApp DeliveryChannel = "whatsapp"
	ChannelTelegram DeliveryChannel = "telegram"
)

// NotificationChannel defines the interface for sending notifications
type NotificationChannel interface {
	// Send delivers a notification to the specified recipient
	Send(ctx context.Context, userID string, notification *domain.Notification) error
	// CanHandle returns true if this channel can handle the given notification type/priority
	CanHandle(notification *domain.Notification) bool
	// Priority returns the priority level this channel handles (for fallback chains)
	Priority() int
}

// ChannelConfig holds configuration for a notification channel
type ChannelConfig struct {
	Enabled  bool                   `json:"enabled"`
	Settings map[string]interface{} `json:"settings"`
}

// NotificationService defines the interface for notification orchestration
type NotificationService interface {
	// Send sends a notification through appropriate channels based on type and user preferences
	Send(ctx context.Context, notification *domain.Notification) error
	// SendBatch sends multiple notifications efficiently
	SendBatch(ctx context.Context, notifications []*domain.Notification) error
	// RegisterChannel registers a new notification channel
	RegisterChannel(channel DeliveryChannel, handler NotificationChannel)
	// GetUserPreferences retrieves user's notification preferences
	GetUserPreferences(ctx context.Context, userID string) (*UserPreferences, error)
	// UpdateUserPreferences updates user's notification preferences
	UpdateUserPreferences(ctx context.Context, userID string, prefs *UserPreferences) error
}

// UserPreferences holds user-specific notification settings
type UserPreferences struct {
	UserID            string                        `json:"user_id"`
	EnabledChannels   []DeliveryChannel             `json:"enabled_channels"`
	ModuleFilters     map[string]bool               `json:"module_filters"` // planning, monitoring, etc.
	TypeFilters       map[string]bool               `json:"type_filters"`   // activity, alert
	PriorityOverrides map[DeliveryChannel]string    `json:"priority_overrides"` // channel -> min priority
	QuietHours        *QuietHours                   `json:"quiet_hours,omitempty"`
}

// QuietHours defines when notifications should be suppressed
type QuietHours struct {
	Enabled   bool   `json:"enabled"`
	StartHour int    `json:"start_hour"` // 0-23
	EndHour   int    `json:"end_hour"`   // 0-23
	Timezone  string `json:"timezone"`
}

// AlertThreshold defines when to escalate notifications
type AlertThreshold struct {
	SeverityLevel string `json:"severity_level"`
	TimeoutMinutes int   `json:"timeout_minutes"`
	EscalateTo    string `json:"escalate_to"` // role or user ID
}
