// Package domain contains all domain models and business logic
package domain

import "time"

// User represents a system user
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	TenantID  string    `json:"tenant_id"`
	Roles     []string  `json:"roles"`
	Sites     []string  `json:"sites"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Tenant represents an organization
type Tenant struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Status      string    `json:"status"` // active, suspended
	Settings    TenantSettings `json:"settings"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TenantSettings holds tenant-specific configurations
type TenantSettings struct {
	DefaultLanguage string   `json:"default_language"` // id_ID, en_ID
	AllowedSites    []string `json:"allowed_sites"`
	CustomFields    map[string]interface{} `json:"custom_fields"`
}

// Site represents a physical or logical location
type Site struct {
	ID        string                 `json:"id"`
	TenantID  string                 `json:"tenant_id"`
	Name      string                 `json:"name"`
	Code      string                 `json:"code"`
	Address   string                 `json:"address"`
	Metadata  map[string]interface{} `json:"metadata"`
	CreatedAt time.Time              `json:"created_at"`
}

// Notification represents a notification entity
type Notification struct {
	ID           string            `json:"id"`
	UserID       string            `json:"user_id"`
	Type         string            `json:"type"` // activity, alert
	Category     string            `json:"category"`
	Title        string            `json:"title"`
	Message      string            `json:"message"`
	Priority     string            `json:"priority"` // low, medium, high, critical
	Module       string            `json:"module"` // planning, monitoring, evaluation, action
	Status       string            `json:"status"` // pending, read, acknowledged, actioned
	Metadata     map[string]interface{} `json:"metadata"`
	Actions      []NotificationAction `json:"actions"`
	CreatedAt    time.Time         `json:"created_at"`
	ReadAt       *time.Time        `json:"read_at"`
	AcknowledgedAt *time.Time      `json:"acknowledged_at"`
}

// NotificationAction represents an action that can be taken on a notification
type NotificationAction struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	ActionType string `json:"action_type"` // accept, decline, view, escalate, etc.
	Endpoint string `json:"endpoint"`
	Method   string `json:"method"`
	Payload  map[string]interface{} `json:"payload,omitempty"`
}

// Team represents a cross-site or site-based team
type Team struct {
	ID         string   `json:"id"`
	TenantID   string   `json:"tenant_id"`
	Name       string   `json:"name"`
	Type       string   `json:"type"` // site_based, multi_site, functional, project
	SiteIDs    []string `json:"site_ids"`
	MemberIDs  []string `json:"member_ids"`
	LeaderID   string   `json:"leader_id"`
	CreatedAt  time.Time `json:"created_at"`
}
