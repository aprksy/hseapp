# Notification Channels

Pluggable notification channel implementations following the Strategy pattern.

## Architecture

All notification channels implement the `NotificationChannel` interface defined in `backend/pkg/notification/interface.go`:

```go
type NotificationChannel interface {
    // Name returns the unique identifier for this channel
    Name() string
    
    // DisplayName returns human-readable name
    DisplayName() string
    
    // Capabilities returns supported features
    Capabilities() ChannelCapabilities
    
    // Send sends a notification through this channel
    Send(ctx context.Context, recipient Recipient, notification Notification) error
    
    // Validate validates channel configuration
    Validate(config map[string]interface{}) error
    
    // IsAvailable checks if channel is operational
    IsAvailable(ctx context.Context) bool
}
```

## Available Channels

### Built-in Channels

1. **Email Channel** (`email.go`)
   - Provider: AWS SES / SendGrid / SMTP
   - Supports: HTML, attachments, templates
   - Priority: Medium

2. **Push Notification Channel** (`push.go`)
   - Provider: Firebase Cloud Messaging (FCM)
   - Supports: Rich notifications, deep links
   - Priority: High

3. **SMS Channel** (`sms.go`)
   - Provider: Twilio / AWS SNS
   - Supports: Text messages, delivery receipts
   - Priority: Critical (fallback)

### Extension Channels (To be implemented)

4. **Telegram Channel** (`telegram.go`)
   - Provider: Telegram Bot API
   - Supports: Messages, inline keyboards, files
   - Priority: Medium

5. **WhatsApp Channel** (`whatsapp.go`)
   - Provider: Twilio WhatsApp API / Meta Business API
   - Supports: Messages, media, interactive buttons
   - Priority: Medium

6. **Slack Channel** (`slack.go`)
   - Provider: Slack Web API
   - Supports: Messages, blocks, reactions
   - Priority: Low (internal teams)

7. **Webhook Channel** (`webhook.go`)
   - Provider: Custom HTTP endpoints
   - Supports: JSON payloads, signatures
   - Priority: Low (integrations)

## Adding a New Channel

### Step 1: Create Channel Implementation

```go
// backend/pkg/notification/channels/telegram.go
package channels

import (
    "context"
    "fmt"
)

type TelegramChannel struct {
    botToken string
    apiURL   string
}

func NewTelegramChannel(config map[string]interface{}) (*TelegramChannel, error) {
    token, ok := config["bot_token"].(string)
    if !ok || token == "" {
        return nil, fmt.Errorf("bot_token required")
    }
    
    return &TelegramChannel{
        botToken: token,
        apiURL:   "https://api.telegram.org/bot" + token,
    }, nil
}

func (c *TelegramChannel) Name() string {
    return "telegram"
}

func (c *TelegramChannel) DisplayName() string {
    return "Telegram"
}

func (c *TelegramChannel) Capabilities() ChannelCapabilities {
    return ChannelCapabilities{
        SupportsHTML:      false,
        SupportsMarkdown:  true,
        SupportsAttachments: true,
        SupportsButtons:   true,
        MaxLength:         4096,
    }
}

func (c *TelegramChannel) Send(ctx context.Context, recipient Recipient, notification Notification) error {
    // Implement Telegram Bot API call
    chatID := recipient.ExternalID // Telegram chat ID
    
    // Send message via Telegram API
    // ... implementation ...
    
    return nil
}

func (c *TelegramChannel) Validate(config map[string]interface{}) error {
    if _, ok := config["bot_token"]; !ok {
        return fmt.Errorf("bot_token required")
    }
    return nil
}

func (c *TelegramChannel) IsAvailable(ctx context.Context) bool {
    // Check Telegram API health
    // ... implementation ...
    return true
}
```

### Step 2: Register in Factory

```go
// backend/pkg/notification/factory.go
func init() {
    RegisterChannel("telegram", func(config map[string]interface{}) (NotificationChannel, error) {
        return NewTelegramChannel(config)
    })
}
```

### Step 3: Add Configuration Schema

```json
{
  "notification.channels.telegram": {
    "bot_token": "required|string",
    "api_url": "optional|string|default=https://api.telegram.org"
  }
}
```

### Step 4: Write Tests

```go
// backend/pkg/notification/channels/telegram_test.go
package channels

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestTelegramChannel_Send(t *testing.T) {
    channel := NewTelegramChannel(map[string]interface{}{
        "bot_token": "test-token",
    })
    
    // Test implementation
}
```

## Channel Selection Logic

Channels are selected based on:

1. **Notification Type**: Activity vs Alert
2. **User Preferences**: Configured in user profile
3. **Priority**: Critical → SMS + Push + Email
4. **Availability**: Fallback to alternative channels
5. **Time of Day**: Respect quiet hours

## Configuration

Channels are configured via environment variables or config files:

```yaml
# config.yaml
notification:
  channels:
    email:
      enabled: true
      provider: ses
      from_address: "noreply@hseapp.com"
    push:
      enabled: true
      provider: fcm
      service_account: "path/to/service-account.json"
    telegram:
      enabled: false
      bot_token: "${TELEGRAM_BOT_TOKEN}"
```

## Testing

Test channels using mocks:

```go
// backend/internal/service/mocks/notification_channel.go
type MockNotificationChannel struct {
    SendFunc func(ctx context.Context, recipient Recipient, notification Notification) error
}

func (m *MockNotificationChannel) Send(ctx context.Context, recipient Recipient, notification Notification) error {
    return m.SendFunc(ctx, recipient, notification)
}

// Implement other interface methods...
```

## Related Documentation

- [Architecture](../../../docs/02-ARCHITECTURE.md)
- [Specification-012: Notifications](../../../specs/00-SPECIFICATIONS.md#spec-012-notifications--communication)
- [Service Interfaces](../../internal/service/)
