# AI Providers

Pluggable AI provider implementations for agentic AI capabilities, including OpenFang integration.

## Architecture

All AI providers implement the `AIProvider` interface defined in `backend/pkg/ai/interface.go`:

```go
type AIProvider interface {
    // Name returns the unique identifier for this provider
    Name() string
    
    // DisplayName returns human-readable name
    DisplayName() string
    
    // Capabilities returns supported AI features
    Capabilities() AICapabilities
    
    // Chat sends a chat message and gets response
    Chat(ctx context.Context, messages []Message) (*Response, error)
    
    // Analyze performs analysis on provided data
    Analyze(ctx context.Context, data AnalysisRequest) (*AnalysisResult, error)
    
    // Recommend generates recommendations
    Recommend(ctx context.Context, context RecommendationContext) ([]Recommendation, error)
    
    // StreamChat streams chat responses (for long-running queries)
    StreamChat(ctx context.Context, messages []Message) (<-chan StreamChunk, error)
    
    // Validate validates provider configuration
    Validate(config map[string]interface{}) error
    
    // IsAvailable checks if provider is operational
    IsAvailable(ctx context.Context) bool
    
    // GetUsage returns usage statistics
    GetUsage(ctx context.Context) (*UsageStats, error)
}
```

## Available Providers

### Phase 2 Providers (Future Implementation)

1. **OpenFang Provider** (`openfang.go`) - Primary Agentic AI Backend
   - Type: Agentic AI (autonomous task execution)
   - Capabilities: Chat, Analysis, Recommendations, Tool Use, Multi-step workflows
   - Priority: High
   - Status: Planned for Phase 2

2. **OpenAI Provider** (`openai.go`)
   - Type: LLM API
   - Capabilities: Chat, Analysis, Recommendations
   - Priority: Medium
   - Status: Fallback option

3. **Anthropic Provider** (`anthropic.go`)
   - Type: LLM API
   - Capabilities: Chat, Analysis, Document Processing
   - Priority: Medium
   - Status: Alternative option

4. **Google Vertex AI Provider** (`vertex.go`)
   - Type: LLM + ML Platform
   - Capabilities: Chat, Analysis, Custom Models
   - Priority: Low
   - Status: Enterprise option

5. **Local LLM Provider** (`local.go`)
   - Type: Self-hosted LLM (Ollama, vLLM)
   - Capabilities: Chat, Analysis (limited)
   - Priority: Low
   - Status: Offline/privacy option

## OpenFang Integration

OpenFang serves as the primary agentic AI backend for HSE App, enabling:

- **Autonomous Agents**: AI agents that can execute multi-step tasks
- **Tool Integration**: Access to HSE App APIs and external tools
- **Workflow Automation**: Automated compliance checks, incident analysis
- **Learning & Adaptation**: Improves recommendations over time

### OpenFang Capabilities

```go
type OpenFangCapabilities struct {
    AgenticExecution    bool  // Can execute autonomous tasks
    ToolUse            bool  // Can use external tools/APIs
    MultiStepWorkflows bool  // Can handle complex workflows
    Memory             bool  // Has persistent memory
    Planning           bool  // Can create and execute plans
    SelfCorrection     bool  // Can detect and fix errors
}
```

### Example Use Cases

1. **Incident Analysis Agent**
   ```
   User: "Analyze the recent accident report from Site A"
   Agent: 
     1. Retrieves accident report from database
     2. Cross-references with similar incidents
     3. Identifies root causes using safety frameworks
     4. Generates corrective action recommendations
     5. Creates follow-up tasks automatically
   ```

2. **Compliance Monitoring Agent**
   ```
   Agent (autonomous):
     1. Daily check of compliance scores across all sites
     2. Identifies sites below threshold
     3. Reviews pending corrective actions
     4. Sends alerts to responsible managers
     5. Updates dashboard with findings
   ```

3. **Safety Recommendation Agent**
   ```
   User: "What safety improvements should we prioritize?"
   Agent:
     1. Analyzes incident trends (last 6 months)
     2. Reviews audit findings
     3. Checks regulatory updates
     4. Compares with industry benchmarks
     5. Provides prioritized recommendations with ROI estimates
   ```

## Adding a New AI Provider

### Step 1: Create Provider Implementation

```go
// backend/pkg/ai/providers/openfang.go
package providers

import (
    "context"
    "fmt"
    "net/http"
)

type OpenFangProvider struct {
    apiKey     string
    baseURL    string
    httpClient *http.Client
    agentID    string
}

func NewOpenFangProvider(config map[string]interface{}) (*OpenFangProvider, error) {
    apiKey, ok := config["api_key"].(string)
    if !ok || apiKey == "" {
        return nil, fmt.Errorf("api_key required")
    }
    
    baseURL := "https://api.openfang.ai/v1"
    if url, ok := config["base_url"].(string); ok {
        baseURL = url
    }
    
    agentID := "hse-app-agent"
    if id, ok := config["agent_id"].(string); ok {
        agentID = id
    }
    
    return &OpenFangProvider{
        apiKey:     apiKey,
        baseURL:    baseURL,
        httpClient: &http.Client{},
        agentID:    agentID,
    }, nil
}

func (p *OpenFangProvider) Name() string {
    return "openfang"
}

func (p *OpenFangProvider) DisplayName() string {
    return "OpenFang Agentic AI"
}

func (p *OpenFangProvider) Capabilities() AICapabilities {
    return AICapabilities{
        Chat:                true,
        Analysis:            true,
        Recommendations:     true,
        AgenticExecution:    true,
        ToolUse:             true,
        MultiStepWorkflows:  true,
        Streaming:           true,
        MaxContextTokens:    128000,
    }
}

func (p *OpenFangProvider) Chat(ctx context.Context, messages []Message) (*Response, error) {
    // Prepare OpenFang agent request
    req := &OpenFangRequest{
        AgentID: p.agentID,
        Messages: p.toOpenFangMessages(messages),
        Tools:    p.getAvailableTools(),
    }
    
    // Call OpenFang API
    resp, err := p.callOpenFangAPI(ctx, "/chat", req)
    if err != nil {
        return nil, err
    }
    
    // Handle agentic response (may include tool calls, plans, etc.)
    return p.processAgenticResponse(resp)
}

func (p *OpenFangProvider) Analyze(ctx context.Context, data AnalysisRequest) (*AnalysisResult, error) {
    // OpenFang-specific analysis implementation
    // Can leverage autonomous agents for complex analysis
}

func (p *OpenFangProvider) Recommend(ctx context.Context, context RecommendationContext) ([]Recommendation, error) {
    // OpenFang-specific recommendation engine
    // Uses learned patterns and industry knowledge
}

// Implement other interface methods...
```

### Step 2: Register in Factory

```go
// backend/pkg/ai/factory.go
func init() {
    RegisterProvider("openfang", func(config map[string]interface{}) (AIProvider, error) {
        return NewOpenFangProvider(config)
    })
}
```

### Step 3: Define Available Tools for OpenFang

```go
// backend/pkg/ai/providers/openfang_tools.go
func (p *OpenFangProvider) getAvailableTools() []OpenFangTool {
    return []OpenFangTool{
        {
            Name:        "get_incident_details",
            Description: "Retrieve detailed information about a specific incident",
            Parameters:  map[string]string{"incident_id": "string"},
        },
        {
            Name:        "query_compliance_status",
            Description: "Check compliance status for a site or regulation",
            Parameters:  map[string]string{"site_id": "string", "regulation": "string"},
        },
        {
            Name:        "create_corrective_action",
            Description: "Create a new corrective action task",
            Parameters:  map[string]string{
                "title": "string",
                "assignee_id": "string",
                "due_date": "string",
                "priority": "string",
            },
        },
        {
            Name:        "generate_safety_report",
            Description: "Generate a comprehensive safety report",
            Parameters:  map[string]string{
                "site_id": "string",
                "period": "string",
                "format": "string",
            },
        },
    }
}
```

### Step 4: Add Configuration

```yaml
# config.yaml
ai:
  default_provider: openfang
  providers:
    openfang:
      enabled: true
      api_key: "${OPENFANG_API_KEY}"
      base_url: "https://api.openfang.ai/v1"
      agent_id: "hse-app-safety-agent"
      tools_enabled: true
      max_steps: 10
      timeout_seconds: 60
    openai:
      enabled: false
      api_key: "${OPENAI_API_KEY}"
      model: "gpt-4o"
```

### Step 5: Write Tests

```go
// backend/pkg/ai/providers/openfang_test.go
package providers

import (
    "testing"
    "context"
    "github.com/stretchr/testify/assert"
)

func TestOpenFangProvider_Chat(t *testing.T) {
    provider := NewOpenFangProvider(map[string]interface{}{
        "api_key": "test-key",
        "agent_id": "test-agent",
    })
    
    messages := []Message{
        {Role: "user", Content: "Analyze incident INC-001"},
    }
    
    resp, err := provider.Chat(context.Background(), messages)
    assert.NoError(t, err)
    assert.NotNil(t, resp)
    // Verify agentic behavior (tool calls, plans, etc.)
}
```

## Provider Selection Logic

Providers are selected based on:

1. **Capability Requirements**: Does provider support needed features?
2. **Cost Optimization**: Use cheaper providers for simple tasks
3. **Performance**: Latency and throughput requirements
4. **Availability**: Fallback to alternative providers
5. **Data Residency**: Compliance with data location requirements
6. **User Preference**: Configured in system settings

## BYOK (Bring Your Own Key)

Phase 2 will support BYOK where tenants can provide their own API keys:

```go
type TenantAIConfig struct {
    TenantID       string
    Provider       string
    APIKey         string  // Encrypted
    CustomEndpoint string  // Optional
    QuotaLimit     int     // Monthly token limit
}
```

## Security Considerations

- API keys encrypted at rest (AES-256)
- Keys never logged or exposed in errors
- Rate limiting per tenant/provider
- Audit logging for all AI interactions
- Data anonymization for sensitive information

## Related Documentation

- [Architecture](../../../docs/02-ARCHITECTURE.md)
- [Specification-013: AI Chat & Analysis](../../../specs/00-SPECIFICATIONS.md#spec-013-ai-chat--analysis-byok)
- [Phase 2 Implementation Plan](../../../plans/00-IMPLEMENTATION_PLAN.md#phase-2-enhancement)
