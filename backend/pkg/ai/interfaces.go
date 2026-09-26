// Package ai defines interfaces for AI provider integration (OpenFang-ready)
package ai

import (
	"context"
)

// AIProvider defines the interface for AI backend providers
type AIProvider interface {
	// Chat sends a message and receives a response
	Chat(ctx context.Context, request *ChatRequest) (*ChatResponse, error)
	
	// Analyze performs analysis on provided data
	Analyze(ctx context.Context, request *AnalysisRequest) (*AnalysisResponse, error)
	
	// Recommend provides recommendations based on context
	Recommend(ctx context.Context, request *RecommendationRequest) (*RecommendationResponse, error)
	
	// StreamChat opens a streaming chat session
	StreamChat(ctx context.Context, request *ChatRequest) (<-chan StreamChunk, error)
	
	// IsAvailable returns true if the provider is healthy and ready
	IsAvailable(ctx context.Context) bool
}

// ProviderType identifies the AI backend provider
type ProviderType string

const (
	ProviderOpenFang   ProviderType = "openfang"
	ProviderOpenAI     ProviderType = "openai"
	ProviderAnthropic  ProviderType = "anthropic"
	ProviderLocal      ProviderType = "local"
)

// ChatRequest holds parameters for a chat request
type ChatRequest struct {
	Messages    []Message          `json:"messages"`
	SystemPrompt string            `json:"system_prompt,omitempty"`
	Temperature float64            `json:"temperature,omitempty"`
	MaxTokens   int                `json:"max_tokens,omitempty"`
	Context     map[string]interface{} `json:"context,omitempty"` // domain-specific context
}

// Message represents a single chat message
type Message struct {
	Role    string `json:"role"` // system, user, assistant
	Content string `json:"content"`
}

// ChatResponse holds the response from a chat request
type ChatResponse struct {
	Message      Message  `json:"message"`
	Usage        TokenUsage `json:"usage"`
	FinishReason string   `json:"finish_reason"`
}

// AnalysisRequest holds parameters for analysis
type AnalysisRequest struct {
	Data       interface{}        `json:"data"`
	DataType   string             `json:"data_type"` // incident, compliance, performance
	Questions  []string           `json:"questions,omitempty"`
	Context    map[string]interface{} `json:"context,omitempty"`
}

// AnalysisResponse holds analysis results
type AnalysisResponse struct {
	Summary     string                 `json:"summary"`
	Findings    []Finding              `json:"findings"`
	RiskLevel   string                 `json:"risk_level"` // low, medium, high, critical
	Confidence  float64                `json:"confidence"`
	Metadata    map[string]interface{} `json:"metadata"`
}

// Finding represents an analysis finding
type Finding struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
	Evidence    []string `json:"evidence"`
	Recommendation string `json:"recommendation"`
}

// RecommendationRequest holds parameters for recommendations
type RecommendationRequest struct {
	Context     string                 `json:"context"`
	Domain      string                 `json:"domain"` // safety, environmental, compliance
	Constraints map[string]interface{} `json:"constraints,omitempty"`
	Goals       []string               `json:"goals,omitempty"`
}

// RecommendationResponse holds recommendation results
type RecommendationResponse struct {
	Recommendations []Recommendation     `json:"recommendations"`
	Priority        string               `json:"priority"`
	Rationale       string               `json:"rationale"`
	ImplementationSteps []string         `json:"implementation_steps,omitempty"`
}

// Recommendation represents a single recommendation
type Recommendation struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Impact      string `json:"impact"`
	Effort      string `json:"effort"` // low, medium, high
	Category    string `json:"category"`
}

// TokenUsage tracks token consumption
type TokenUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// StreamChunk represents a chunk of streamed response
type StreamChunk struct {
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

// ProviderConfig holds configuration for an AI provider
type ProviderConfig struct {
	Type       ProviderType         `json:"type"`
	Endpoint   string               `json:"endpoint"`
	APIKey     string               `json:"api_key"`
	Model      string               `json:"model"`
	Timeout    int                  `json:"timeout_seconds"`
	Extra      map[string]interface{} `json:"extra,omitempty"`
}

// AIFactory creates AI provider instances
type AIFactory interface {
	CreateProvider(config ProviderConfig) (AIProvider, error)
	GetProvider(providerType ProviderType) (AIProvider, error)
	ListProviders() []ProviderType
}
