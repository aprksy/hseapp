# 02. ARCHITECTURE & DESIGN PRINCIPLES

## 1. Domain-Driven Design (DDD) Approach

### 1.1 Strategic Design

#### Core Domains
1. **Identity & Access Management (IAM)**
   - Authentication, Authorization, Multi-tenancy
   - Users, Roles, Permissions, Teams
   - OAuth providers, MFA

2. **Organization & Site Management**
   - Tenants, Sites, Departments
   - Cross-site team structures
   - Hierarchical relationships

3. **Compliance Management**
   - Regulations (Kemenaker, KLHK, BSKAP)
   - Checklists, Audits, Inspections
   - Compliance scoring, Indicators

4. **Incident & Risk Management**
   - Incident reporting, Investigation workflows
   - Risk assessments, Hazard identification
   - Corrective & Preventive Actions (CAPA)

5. **Planning & Monitoring**
   - Task assignments, Schedules
   - Progress tracking, Kanban boards
   - Timeline management

6. **Evaluation & Reporting**
   - Performance evaluations
   - Regulatory reports, Analytics
   - Dashboards, KPIs

7. **Notifications & Communication**
   - Activity notifications, Alerts
   - Multi-channel delivery (Email, SMS, Push, Telegram, WhatsApp)
   - Preferences, Templates

8. **Document & Form Management**
   - Form builder, Dynamic forms
   - Document storage, Versioning
   - Import/Export (PDF, Excel, PNG, TXT)

#### Supporting Domains
- **Audit Logging**: All critical actions tracked
- **File Storage**: Attachments, Photos, Certificates
- **Localization**: i18n, Regional formats
- **Integration Gateway**: External systems (ERP, HRIS, Government APIs)

#### Generic Domains
- **Authentication**: Standard auth flows
- **Notification Infrastructure**: Channel-agnostic messaging
- **Reporting Engine**: Chart generation, Export utilities

### 1.2 Tactical Design Patterns

#### Repository Pattern
```go
// Domain entity
type Incident struct {
    ID          uuid.UUID
    TenantID    uuid.UUID
    SiteID      uuid.UUID
    Title       string
    Severity    IncidentSeverity
    Status      IncidentStatus
    ReportedAt  time.Time
    // ... other fields
}

// Repository interface (domain layer)
type IncidentRepository interface {
    Create(ctx context.Context, incident *Incident) error
    GetByID(ctx context.Context, id uuid.UUID) (*Incident, error)
    Update(ctx context.Context, incident *Incident) error
    FindByTenant(ctx context.Context, tenantID uuid.UUID, filters IncidentFilters) ([]*Incident, error)
    FindBySite(ctx context.Context, siteID uuid.UUID, filters IncidentFilters) ([]*Incident, error)
}

// PostgreSQL implementation (infrastructure layer)
type postgresIncidentRepository struct {
    db *sql.DB
}

func (r *postgresIncidentRepository) Create(ctx context.Context, incident *Incident) error {
    // Implementation with RLS enforcement
}
```

#### Service Layer (Business Logic)
```go
// Domain service interface
type IncidentService interface {
    ReportIncident(ctx context.Context, cmd ReportIncidentCommand) (*Incident, error)
    AssignInvestigator(ctx context.Context, cmd AssignInvestigatorCommand) error
    EscalateIncident(ctx context.Context, cmd EscalateIncidentCommand) error
    CloseIncident(ctx context.Context, cmd CloseIncidentCommand) error
}

// Implementation
type incidentService struct {
    repo           IncidentRepository
    notifier       NotificationService
    eventPublisher EventPublisher
}

func (s *incidentService) ReportIncident(ctx context.Context, cmd ReportIncidentCommand) (*Incident, error) {
    // Business logic validation
    // Domain rules enforcement
    // Publish domain events
}
```

#### Factory Pattern (Extensibility)
```go
// Notification channel factory
type NotificationChannelFactory interface {
    GetChannel(channelType string) (NotificationChannel, error)
    RegisterChannel(channelType string, channel NotificationChannel)
}

type notificationChannelFactory struct {
    channels map[string]NotificationChannel
}

func (f *notificationChannelFactory) GetChannel(channelType string) (NotificationChannel, error) {
    channel, exists := f.channels[channelType]
    if !exists {
        return nil, fmt.Errorf("channel type %s not registered", channelType)
    }
    return channel, nil
}
```

## 2. Component Contracts & Interfaces

### 2.1 Backend ↔ Frontend Contracts

#### REST API (OpenAPI 3.0)
- All endpoints documented in `/api/openapi.yaml`
- Versioned: `/api/v1/`, `/api/v2/`
- Request/Response schemas strictly defined
- Error response standardization

```yaml
# Example OpenAPI spec snippet
paths:
  /api/v1/incidents:
    post:
      summary: Report new incident
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/ReportIncidentRequest'
      responses:
        '201':
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Incident'
```

#### WebSocket Events (Real-time)
```typescript
// TypeScript interface for web client
interface ServerEvent {
  type: 'INCIDENT_CREATED' | 'TASK_ASSIGNED' | 'ALERT_TRIGGERED';
  payload: any;
  timestamp: string;
  tenantId: string;
}

// Subscription example
supabase.channel('incidents:tenant_123')
  .on('system', { event: 'INCIDENT_CREATED' }, handleNewIncident)
  .subscribe();
```

#### gRPC (Internal Services)
```protobuf
// proto/notification.proto
syntax = "proto3";
package notification.v1;

service NotificationService {
  rpc Send(SendNotificationRequest) returns (SendNotificationResponse);
  rpc ValidateChannel(ValidateChannelRequest) returns (ValidateChannelResponse);
}

message SendNotificationRequest {
  string channel_type = 1; // email, sms, push, telegram, whatsapp
  string recipient = 2;
  string subject = 3;
  string body = 4;
  NotificationMetadata metadata = 5;
}

message NotificationMetadata {
  string tenant_id = 1;
  string category = 2; // ACTIVITY or ALERT
  string action_url = 3;
  map<string, string> extra = 4;
}
```

### 2.2 Backend Internal Interfaces

#### Notification Channel Interface
```go
// domain/notification/channel.go
package notification

import "context"

// NotificationChannel defines contract for all notification providers
type NotificationChannel interface {
    // Send delivers notification to recipient
    Send(ctx context.Context, req SendRequest) (*SendResponse, error)
    
    // Validate checks if channel configuration is valid
    Validate(ctx context.Context, config ChannelConfig) error
    
    // Name returns channel identifier (email, sms, telegram, etc.)
    Name() string
    
    // Capabilities returns supported features
    Capabilities() ChannelCapabilities
}

type SendRequest struct {
    Recipients []string
    Subject    string
    Body       string
    Category   NotificationCategory // ACTIVITY or ALERT
    Metadata   NotificationMetadata
}

type SendResponse struct {
    MessageID string
    Status    SendStatus
    Error     error
}

type ChannelCapabilities struct {
    SupportsAttachments bool
    SupportsTemplates   bool
    MaxRecipients       int
    DeliveryTracking    bool
}
```

#### Email Channel Implementation
```go
// infrastructure/notification/email_channel.go
type emailChannel struct {
    smtpClient SMTPClient
    config     EmailConfig
}

func (c *emailChannel) Send(ctx context.Context, req SendRequest) (*SendResponse, error) {
    // Implementation
}

func (c *emailChannel) Name() string {
    return "email"
}
```

#### Telegram Channel Implementation (Future)
```go
// infrastructure/notification/telegram_channel.go
type telegramChannel struct {
    botClient  TelegramBotClient
    config     TelegramConfig
}

func (c *telegramChannel) Send(ctx context.Context, req SendRequest) (*SendResponse, error) {
    // Call Telegram Bot API
    // Handle rate limits, retries
}

func (c *telegramChannel) Name() string {
    return "telegram"
}
```

#### AI Provider Interface (OpenFang Ready)
```go
// domain/ai/provider.go
package ai

import "context"

// AIProvider defines contract for AI backends
type AIProvider interface {
    // Chat sends message and receives response
    Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error)
    
    // Analyze processes data and returns insights
    Analyze(ctx context.Context, req AnalysisRequest) (*AnalysisResponse, error)
    
    // Recommend generates recommendations based on context
    Recommend(ctx context.Context, req RecommendationRequest) (*RecommendationResponse, error)
    
    // Name returns provider identifier
    Name() string
    
    // IsAvailable checks if provider is configured and healthy
    IsAvailable(ctx context.Context) bool
}

type ChatRequest struct {
    Messages    []Message
    Context     AIContext
    Temperature float32
    MaxTokens   int
}

type AIContext struct {
    TenantID    string
    Domain      string // compliance, incident, training
    UserRole    string
    Language    string
}
```

#### OpenFang Implementation (Phase 2)
```go
// infrastructure/ai/openfang_provider.go
type openFangProvider struct {
    client       OpenFangClient
    config       OpenFangConfig
    agentManager AgentManager
}

func (p *openFangProvider) Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error) {
    // Route to appropriate OpenFang agent
    // Handle agentic workflows
}

func (p *openFangProvider) Name() string {
    return "openfang"
}
```

#### Storage Provider Interface
```go
// domain/storage/provider.go
package storage

import "context"

type StorageProvider interface {
    Upload(ctx context.Context, req UploadRequest) (*UploadResponse, error)
    Download(ctx context.Context, fileID string) (*DownloadResponse, error)
    Delete(ctx context.Context, fileID string) error
    GetURL(ctx context.Context, fileID string, expires time.Duration) (string, error)
    Name() string
}
```

### 2.3 Mobile ↔ Backend Contracts

#### Sync Protocol
```dart
// Flutter sync request/response
class SyncRequest {
  final String lastSyncTimestamp;
  final List<PendingChange> pendingChanges;
  final String deviceId;
}

class SyncResponse {
  final String newSyncTimestamp;
  final List<ServerChange> serverChanges;
  final List<ConflictResolution> conflicts;
}
```

#### Push Notification Payload
```json
{
  "to": "device_token",
  "notification": {
    "title": "Insiden Baru Dilaporkan",
    "body": "Kecelakaan kerja di Site A requires attention",
    "data": {
      "type": "ALERT",
      "category": "INCIDENT",
      "incidentId": "uuid-here",
      "actionUrl": "/incidents/uuid-here",
      "priority": "high",
      "tenantId": "tenant-uuid"
    }
  },
  "android": {
    "priority": "high",
    "notification": {
      "click_action": "FLUTTER_NOTIFICATION_CLICK"
    }
  },
  "apns": {
    "payload": {
      "aps": {
        "content-available": 1
      }
    }
  }
}
```

## 3. Repository Structure

### 3.1 Monorepo Layout (Recommended)
```
hse-app/
├── Makefile                    # Root makefile for all ops
├── README.md
├── docs/                       # Documentation
│   ├── 01-CONSTITUTION.md
│   ├── 02-ARCHITECTURE.md      # This document
│   └── ...
├── specs/                      # Specifications
│   └── 00-SPECIFICATIONS.md
├── plans/                      # Implementation plans
│   └── 00-IMPLEMENTATION_PLAN.md
├── tasks/                      # Task definitions
│   └── PHASE0-TASKS.md
├── backend/                    # Go backend
│   ├── cmd/
│   │   ├── api/               # Main API server
│   │   └── worker/            # Background jobs
│   ├── internal/
│   │   ├── domain/            # Domain models & interfaces
│   │   │   ├── iam/
│   │   │   ├── incident/
│   │   │   ├── compliance/
│   │   │   └── notification/
│   │   ├── service/           # Business logic implementations
│   │   ├── repository/        # Repository interfaces
│   │   └── handler/           # HTTP/gRPC handlers
│   ├── pkg/                   # Shared packages
│   │   ├── notification/
│   │   │   ├── channel.go     # Interfaces
│   │   │   ├── email.go
│   │   │   ├── sms.go
│   │   │   ├── push.go
│   │   │   └── telegram.go    # Future
│   │   ├── ai/
│   │   │   ├── provider.go    # Interface
│   │   │   └── openfang.go    # Future implementation
│   │   └── storage/
│   │       ├── provider.go
│   │       └── s3.go
│   ├── infrastructure/        # External implementations
│   │   ├── database/
│   │   ├── cache/
│   │   └── messaging/
│   ├── api/
│   │   ├── openapi.yaml
│   │   └── handlers/
│   └── configs/
├── web/                        # SvelteKit frontend
│   ├── src/
│   │   ├── lib/
│   │   │   ├── components/    # Shadcn blocks + custom
│   │   │   ├── stores/        # Svelte stores
│   │   │   ├── services/      # API clients
│   │   │   ├── i18n/          # Translations
│   │   │   └── types/         # TypeScript types (from OpenAPI)
│   │   ├── routes/            # SvelteKit routes
│   │   └── app.html
│   ├── static/
│   └── package.json
├── mobile/                     # Flutter app (separate project)
│   ├── lib/
│   │   ├── core/
│   │   │   ├── models/        # Generated from Protobuf/OpenAPI
│   │   │   ├── services/      # API, Sync, Notifications
│   │   │   └── utils/
│   │   ├── features/
│   │   │   ├── auth/
│   │   │   ├── incidents/
│   │   │   ├── tasks/
│   │   │   └── dashboard/
│   │   ├── shared/
│   │   │   ├── widgets/
│   │   │   └── themes/
│   │   └── main.dart
│   ├── android/
│   ├── ios/
│   └── pubspec.yaml
├── proto/                      # Protobuf definitions
│   ├── notification.proto
│   ├── ai.proto
│   └── sync.proto
├── deployments/                # Kubernetes, Docker Compose
│   ├── docker-compose.yml
│   └── k8s/
└── scripts/                    # Utility scripts
    ├── migrate.sh
    └── seed.sh
```

### 3.2 Alternative: Polyrepo Structure
If teams are large and independent:
- `hse-backend`: Go codebase
- `hse-web`: SvelteKit codebase
- `hse-mobile`: Flutter codebase
- `hse-proto`: Shared Protobuf definitions
- `hse-docs`: Documentation site

## 4. Event-Driven Architecture

### 4.1 Domain Events
```go
// domain/event/events.go
type DomainEvent interface {
    EventType() string
    AggregateID() uuid.UUID
    TenantID() uuid.UUID
    OccurredAt() time.Time
}

// Incident events
type IncidentReportedEvent struct {
    IncidentID uuid.UUID
    TenantID   uuid.UUID
    SiteID     uuid.UUID
    Severity   string
    ReportedBy uuid.UUID
    Timestamp  time.Time
}

type IncidentAssignedEvent struct {
    IncidentID   uuid.UUID
    AssignedTo   uuid.UUID
    AssignedBy   uuid.UUID
    Priority     string
    Timestamp    time.Time
}

// Task events
type TaskAssignedEvent struct {
    TaskID     uuid.UUID
    AssigneeID uuid.UUID
    TeamID     uuid.UUID
    DueDate    time.Time
    Timestamp  time.Time
}

// Alert events
type EmergencyAlertEvent struct {
    AlertID    uuid.UUID
    SiteID     uuid.UUID
    Type       string // accident, environmental_breach, medical_emergency
    Severity   string // critical, high, medium
    RequiresAction bool
    Timestamp  time.Time
}
```

### 4.2 Event Publishing & Subscription
```go
// infrastructure/messaging/event_bus.go
type EventBus interface {
    Publish(ctx context.Context, event DomainEvent) error
    Subscribe(eventType string, handler EventHandler) error
}

// Example: Notification service subscribes to events
func SetupNotificationSubscribers(eventBus EventBus, notifier NotificationService) {
    eventBus.Subscribe("IncidentReportedEvent", func(event DomainEvent) error {
        e := event.(*IncidentReportedEvent)
        // Send alert to HSE Manager
        notifier.SendAlert(...)
        return nil
    })
    
    eventBus.Subscribe("TaskAssignedEvent", func(event DomainEvent) error {
        e := event.(*TaskAssignedEvent)
        // Send activity notification to assignee
        notifier.SendActivity(...)
        return nil
    })
}
```

### 4.3 Event Schema Versioning
```json
{
  "event_type": "IncidentReportedEvent",
  "schema_version": "1.0",
  "data": { ... },
  "metadata": {
    "tenant_id": "...",
    "occurred_at": "2026-09-13T14:36:00Z"
  }
}
```

## 5. Extensibility Guidelines

### 5.1 Adding New Notification Channel (e.g., WhatsApp)

1. **Implement Interface**
```go
// infrastructure/notification/whatsapp_channel.go
type whatsappChannel struct {
    client WhatsAppClient
    config WhatsAppConfig
}

func (c *whatsappChannel) Send(ctx context.Context, req SendRequest) (*SendResponse, error) {
    // WhatsApp Business API integration
}

func (c *whatsappChannel) Name() string {
    return "whatsapp"
}
```

2. **Register in Factory**
```go
// pkg/notification/factory.go
func NewNotificationChannelFactory() NotificationChannelFactory {
    factory := &notificationChannelFactory{
        channels: make(map[string]NotificationChannel),
    }
    
    // Register built-in channels
    factory.RegisterChannel("email", NewEmailChannel(config))
    factory.RegisterChannel("sms", NewSMSChannel(config))
    factory.RegisterChannel("push", NewPushChannel(config))
    
    // Register new channel
    factory.RegisterChannel("whatsapp", NewWhatsAppChannel(config))
    
    return factory
}
```

3. **Update Configuration**
```yaml
# configs/notification.yaml
channels:
  whatsapp:
    enabled: true
    api_key: "${WHATSAPP_API_KEY}"
    phone_number_id: "${WHATSAPP_PHONE_ID}"
    template_mappings:
      incident_alert: "whatsapp_template_001"
      task_assignment: "whatsapp_template_002"
```

4. **Add UI Configuration** (Web Admin)
- Add WhatsApp option to notification preferences
- Template management interface

**Estimated Effort**: <2 days (as per success metrics)

### 5.2 Integrating OpenFang as AI Backend

1. **Implement AI Provider Interface**
```go
// infrastructure/ai/openfang_provider.go
type openFangProvider struct {
    client       *openfang.Client
    config       OpenFangConfig
    agents       map[string]AgentConfig
}

func (p *openFangProvider) Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error) {
    // Select appropriate agent based on context
    agent := p.selectAgent(req.Context.Domain)
    
    // Build agentic workflow
    response, err := p.client.ExecuteAgent(ctx, agent, req.Messages)
    
    return &ChatResponse{
        Message: response.Message,
        Sources: response.Sources,
        Confidence: response.Confidence,
    }, nil
}

func (p *openFangProvider) Analyze(ctx context.Context, req AnalysisRequest) (*AnalysisResponse, error) {
    // Use OpenFang's analysis agents
    // Return structured insights
}

func (p *openFangProvider) Name() string {
    return "openfang"
}
```

2. **Register AI Provider**
```go
// pkg/ai/factory.go
func NewAIProviderFactory() AIProviderFactory {
    factory := &aiProviderFactory{
        providers: make(map[string]AIProvider),
    }
    
    // Phase 2: Add OpenFang
    factory.RegisterProvider("openfang", NewOpenFangProvider(config))
    
    // Future: Add other providers
    // factory.RegisterProvider("openai", NewOpenAIProvider(config))
    
    return factory
}
```

3. **Configure Agents**
```yaml
# configs/ai.yaml
providers:
  openfang:
    enabled: true
    api_endpoint: "${OPENFANG_API_URL}"
    api_key: "${OPENFANG_API_KEY}"
    
agents:
  compliance_advisor:
    model: "compliance-v2"
    system_prompt: "You are an Indonesian HSE compliance expert..."
    tools: ["regulation_search", "checklist_generator"]
    
  incident_analyst:
    model: "investigation-v1"
    system_prompt: "You are an incident investigation specialist..."
    tools: ["root_cause_analysis", "corrective_action_recommendation"]
```

4. **Update Frontend**
- Add chat interface component
- Integrate AI response streaming
- Display sources and confidence scores

## 6. Security & Multi-Tenancy

### 6.1 Row-Level Security (Supabase)
```sql
-- Enable RLS on incidents table
ALTER TABLE incidents ENABLE ROW LEVEL SECURITY;

-- Policy: Users can only see incidents from their tenant
CREATE POLICY tenant_isolation ON incidents
    FOR ALL
    USING (tenant_id = current_setting('app.current_tenant')::uuid);

-- Policy: Site-based access control
CREATE POLICY site_access ON incidents
    FOR SELECT
    USING (
        site_id IN (
            SELECT site_id FROM user_sites
            WHERE user_id = current_setting('app.current_user')::uuid
        )
    );
```

### 6.2 Context Propagation
```go
// middleware/tenant_context.go
func TenantContextMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tenantID := r.Header.Get("X-Tenant-ID")
        userID := r.Context().Value("user_id").(string)
        
        // Set in Postgres session
        ctx := context.WithValue(r.Context(), "tenant_id", tenantID)
        ctx = context.WithValue(ctx, "user_id", userID)
        
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// repository/base.go
func (r *baseRepository) setContext(ctx context.Context, conn *sql.Conn) error {
    tenantID := ctx.Value("tenant_id").(string)
    userID := ctx.Value("user_id").(string)
    
    _, err := conn.ExecContext(ctx, 
        "SET app.current_tenant = $1; SET app.current_user = $2;",
        tenantID, userID)
    
    return err
}
```

## 7. Testing Strategy

### 7.1 Unit Tests (Domain Layer)
```go
// internal/service/incident_service_test.go
func TestIncidentService_ReportIncident(t *testing.T) {
    // Arrange
    mockRepo := new(MockIncidentRepository)
    mockNotifier := new(MockNotificationService)
    service := NewIncidentService(mockRepo, mockNotifier)
    
    cmd := ReportIncidentCommand{
        Title: "Test Incident",
        Severity: "HIGH",
    }
    
    // Act
    incident, err := service.ReportIncident(context.Background(), cmd)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, incident)
    mockRepo.AssertCalled(t, "Create", mock.Anything, mock.Anything)
    mockNotifier.AssertCalled(t, "SendAlert", mock.Anything)
}
```

### 7.2 Integration Tests (Repository Layer)
```go
// internal/repository/incident_repository_test.go
func TestPostgresIncidentRepository_Create(t *testing.T) {
    // Setup test database with testcontainers
    db := setupTestDatabase()
    repo := NewPostgresIncidentRepository(db)
    
    incident := &Incident{
        TenantID: testTenantID,
        Title: "Test",
    }
    
    err := repo.Create(context.Background(), incident)
    
    assert.NoError(t, err)
    assert.NotEmpty(t, incident.ID)
}
```

### 7.3 Contract Tests (API)
```go
// api/handlers/incident_handler_test.go
func TestIncidentHandler_CreateIncident_Contract(t *testing.T) {
    // Use pact.io or similar for contract testing
    // Ensure API matches OpenAPI spec
}
```

### 7.4 E2E Tests (Full Stack)
```typescript
// web/e2e/incident-reporting.spec.ts
test('User can report incident and receive confirmation', async ({ page }) => {
    await page.goto('/incidents/new');
    await page.fill('[name="title"]', 'Test Incident');
    await page.selectOption('[name="severity"]', 'HIGH');
    await page.click('button[type="submit"]');
    
    await expect(page.locator('.toast-success')).toBeVisible();
    await expect(page).toHaveURL(/\/incidents\/[a-f0-9-]+/);
});
```

## 8. Documentation & Code Generation

### 8.1 OpenAPI Client Generation
```makefile
# Makefile
.PHONY: generate-clients

generate-clients:
    # Generate TypeScript client for web
    npx openapi-typescript-codegen \
        --input backend/api/openapi.yaml \
        --output web/src/lib/services/api-client \
        --client fetch
    
    # Generate Dart client for mobile
    docker run --rm \
        -v ${PWD}/backend/api/openapi.yaml:/openapi.yaml \
        -v ${PWD}/mobile/lib/core/services:/output \
        openapitools/openapi-generator-cli generate \
        -i /openapi.yaml \
        -g dart-dio \
        -o /output
```

### 8.2 Protobuf Generation
```makefile
.PHONY: generate-protobuf

generate-protobuf:
    # Generate Go gRPC code
    protoc --go_out=. --go-grpc_out=. \
        proto/*.proto
    
    # Generate TypeScript gRPC-web code
    protoc --grpc-web_out=. \
        --plugin=protoc-gen-grpc-web=$(which protoc-gen-grpc-web) \
        proto/*.proto
```

## 9. Monitoring & Observability

### 9.1 Metrics to Track
- API latency (p50, p95, p99)
- Notification delivery rates by channel
- AI provider response times
- Sync conflict rates (mobile)
- Multi-tenant resource usage

### 9.2 Distributed Tracing
```go
// instrumentation/tracing.go
func TraceRepositoryCall(ctx context.Context, repoName, methodName string) func() {
    span := otel.Tracer("hse-app").Start(ctx, fmt.Sprintf("%s.%s", repoName, methodName))
    span.SetAttributes(attribute.String("tenant_id", GetTenantID(ctx)))
    
    return func() {
        span.End()
    }
}

// Usage in repository
func (r *incidentRepository) Create(ctx context.Context, incident *Incident) error {
    defer TraceRepositoryCall(ctx, "IncidentRepository", "Create")()
    // Implementation
}
```

---

**Document Status**: Draft v1.0  
**Created**: Architecture & Design Principles  
**Next Review**: Before Phase 1 implementation  
**Approval Required**: Yes - Critical for development consistency
