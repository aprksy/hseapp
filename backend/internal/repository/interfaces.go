// Package repository defines repository interfaces for data access
package repository

import (
	"context"
	"github.com/aprksy/hseapp/backend/internal/domain"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error
	ListByTenant(ctx context.Context, tenantID string, limit, offset int) ([]*domain.User, error)
}

// TenantRepository defines the interface for tenant data access
type TenantRepository interface {
	GetByID(ctx context.Context, id string) (*domain.Tenant, error)
	GetBySlug(ctx context.Context, slug string) (*domain.Tenant, error)
	Create(ctx context.Context, tenant *domain.Tenant) error
	Update(ctx context.Context, tenant *domain.Tenant) error
	List(ctx context.Context, limit, offset int) ([]*domain.Tenant, error)
}

// SiteRepository defines the interface for site data access
type SiteRepository interface {
	GetByID(ctx context.Context, id string) (*domain.Site, error)
	ListByTenant(ctx context.Context, tenantID string) ([]*domain.Site, error)
	Create(ctx context.Context, site *domain.Site) error
	Update(ctx context.Context, site *domain.Site) error
	Delete(ctx context.Context, id string) error
}

// NotificationRepository defines the interface for notification data access
type NotificationRepository interface {
	GetByID(ctx context.Context, id string) (*domain.Notification, error)
	ListByUser(ctx context.Context, userID string, limit, offset int) ([]*domain.Notification, error)
	Create(ctx context.Context, notif *domain.Notification) error
	Update(ctx context.Context, notif *domain.Notification) error
	MarkAsRead(ctx context.Context, id string) error
	Acknowledge(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}

// TeamRepository defines the interface for team data access
type TeamRepository interface {
	GetByID(ctx context.Context, id string) (*domain.Team, error)
	ListByTenant(ctx context.Context, tenantID string) ([]*domain.Team, error)
	ListByUser(ctx context.Context, userID string) ([]*domain.Team, error)
	Create(ctx context.Context, team *domain.Team) error
	Update(ctx context.Context, team *domain.Team) error
	Delete(ctx context.Context, id string) error
	AddMember(ctx context.Context, teamID, userID string) error
	RemoveMember(ctx context.Context, teamID, userID string) error
}
