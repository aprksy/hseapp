# Backend Development Guidelines

## Overview
This document provides guidelines for developing the backend services of the HSE App using Go, following Domain-Driven Design (DDD) principles.

## Project Structure

```
backend/
├── cmd/                    # Application entry points
│   └── server/             # Main HTTP server
├── internal/               # Private application code
│   ├── domain/             # Domain models and business logic
│   │   ├── user/           # User domain
│   │   ├── organization/   # Organization & Multi-tenancy
│   │   ├── compliance/     # Compliance management
│   │   ├── incident/       # Incident reporting
│   │   ├── notification/   # Notification system
│   │   └── ai/             # AI integration domain
│   ├── repository/         # Data access layer
│   │   ├── postgres/       # PostgreSQL implementations
│   │   ├── supabase/       # Supabase-specific implementations
│   │   └── memory/         # In-memory implementations (testing)
│   ├── service/            # Business logic orchestration
│   │   ├── auth/           # Authentication services
│   │   ├── tenant/         # Multi-tenancy services
│   │   ├── approval/       # Approval routing
│   │   └── notification/   # Notification delivery
│   ├── handler/            # HTTP/gRPC handlers
│   │   ├── http/           # REST API handlers
│   │   └── grpc/           # gRPC service handlers
│   ├── factory/            # Factory patterns for extensibility
│   │   ├── notification/   # Notification channel factories
│   │   └── ai/             # AI provider factories
│   ├── middleware/         # HTTP/gRPC middleware
│   │   ├── auth.go         # JWT authentication
│   │   ├── tenant.go       # Tenant isolation
│   │   └── logging.go      # Request logging
│   └── config/             # Configuration management
├── pkg/                    # Public library code
│   ├── errors/             # Custom error types
│   ├── utils/              # Utility functions
│   └── proto/              # Generated protobuf files
├── configs/                # Configuration files
├── migrations/             # Database migrations
├── scripts/                # Development scripts
├── go.mod                  # Go module definition
├── go.sum                  # Dependency checksums
└── Makefile                # Build automation
```

## Domain-Driven Design Implementation

### Domain Models
Each domain should have:
- Entity definitions with business logic
- Value objects for immutable concepts
- Domain events for cross-domain communication
- Repository interfaces defining data access contracts

Example (`internal/domain/user/user.go`):
```go
package user

type User struct {
    ID        string
    Email     string
    Roles     []string
    TenantID  string
    Sites     []string
    CreatedAt time.Time
}

func (u *User) HasRole(role string) bool {
    for _, r := range u.Roles {
        if r == role {
            return true
        }
    }
    return false
}

func (u *User) CanAccessSite(siteID string) bool {
    for _, s := range u.Sites {
        if s == siteID || s == "*" {
            return true
        }
    }
    return false
}
```

### Repository Pattern
Define interfaces in domain layer, implement in repository layer:

```go
// internal/domain/user/repository.go
package user

type Repository interface {
    FindByID(ctx context.Context, id string) (*User, error)
    FindByEmail(ctx context.Context, email string) (*User, error)
    Save(ctx context.Context, user *User) error
    Delete(ctx context.Context, id string) error
}

// internal/repository/postgres/user_repository.go
package postgres

type userRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) user.Repository {
    return &userRepository{db: db}
}
```

### Service Layer
Orchestrate domain logic and coordinate between repositories:

```go
// internal/service/auth/service.go
package auth

type Service struct {
    userRepo     user.Repository
    tenantRepo   tenant.Repository
    jwtSecret    string
}

func (s *Service) Login(ctx context.Context, email, password string) (*TokenPair, error) {
    // Business logic here
}
```

## API Design

### REST Conventions
- Use resource-oriented URLs: `/api/v1/users`, `/api/v1/tenants/{id}/sites`
- Standard HTTP methods: GET, POST, PUT, DELETE, PATCH
- Consistent response format:
```json
{
  "data": {},
  "meta": {
    "page": 1,
    "limit": 20,
    "total": 100
  },
  "error": null
}
```

### Error Handling
Use custom error types in `pkg/errors`:
```go
type AppError struct {
    Code       string
    Message    string
    StatusCode int
    Details    map[string]interface{}
}

func NewAppError(code, message string, statusCode int) *AppError {
    return &AppError{
        Code:       code,
        Message:    message,
        StatusCode: statusCode,
    }
}
```

## Multi-Tenancy

All queries must include tenant filtering:
```go
func (r *userRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
    tenantID, ok := ctx.Value("tenant_id").(string)
    if !ok {
        return nil, errors.New("tenant_id not found in context")
    }
    
    query := "SELECT * FROM users WHERE id = $1 AND tenant_id = $2"
    // Execute query
}
```

## Testing Strategy

### Unit Tests
Test individual functions in isolation:
```bash
make test-unit
```

### Integration Tests
Test with real database (Supabase local):
```bash
make test-integration
```

### Contract Tests
Verify API contracts between services:
```bash
make test-contract
```

## Code Generation

Generate protobuf clients:
```bash
make proto-generate
```

Generate OpenAPI specs:
```bash
make openapi-generate
```

## Security

- JWT tokens for authentication
- Row-level security via Supabase
- Input validation on all endpoints
- Rate limiting per tenant
- Audit logging for sensitive operations

## Monitoring

- OpenTelemetry for tracing
- Prometheus metrics
- Structured logging with Zap

## Development Workflow

1. Create domain model
2. Define repository interface
3. Implement repository
4. Create service layer
5. Add HTTP/gRPC handlers
6. Write tests
7. Update API documentation

## Adding New Features

Follow this sequence:
1. Update specification in `/specs`
2. Create domain models
3. Implement repository pattern
4. Add service logic
5. Create API handlers
6. Write integration tests
7. Update documentation

## Extensibility

### Adding Notification Channels
1. Implement `NotificationChannel` interface
2. Register in notification factory
3. Add configuration options
4. Write tests

### Adding AI Providers
1. Implement `AIProvider` interface
2. Register in AI factory
3. Configure provider settings
4. Test with OpenFang compatibility

## References

- [Go Best Practices](https://github.com/golang-standards/project-layout)
- [Domain-Driven Design](https://martinfowler.com/bliki/DomainDrivenDesign.html)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
