package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"ai-jar-sdk-generator/src/core"
)

// AIModelConfig holds configuration for an AI model.
type AIModelConfig struct {
	Endpoint string `json:"endpoint"` // API endpoint URL
	APIKey   string `json:"api_key"` // API key for authentication
	Model    string `json:"model"`    // Model name, e.g., "qwen3-72b-chat", "deepseek-coder"
	Timeout  int    `json:"timeout"`  // Request timeout in seconds
}

// AIRequest represents the request structure for AI API calls.
type AIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature,omitempty"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
}

// Message represents a single message in the conversation.
type Message struct {
	Role    string `json:"role"`    // "system", "user", or "assistant"
	Content string `json:"content"` // Message content
}

// AIResponse represents the response structure from AI API.
type AIResponse struct {
	Choices []Choice `json:"choices"`
	Error   *APIError `json:"error,omitempty"`
}

// Choice represents a single choice in the AI response.
type Choice struct {
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

// APIError represents an error response from the AI API.
type APIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code"`
}

// DocumentationResult holds the generated documentation for a class.
type DocumentationResult struct {
	ClassName     string `json:"class_name"`
	Description   string `json:"description"`
	Purpose       string `json:"purpose"`
	UsageExample  string `json:"usage_example"`
	Dependencies  string `json:"dependencies"`
	GeneratedAt   time.Time `json:"generated_at"`
}

// GenerateAIDocumentation sends code information to an AI model and gets documentation.
func GenerateAIDocumentation(classInfo *core.DecompiledClass, dependencyResult *core.DependencyAnalysisResult, config AIModelConfig) (*DocumentationResult, error) {
	if classInfo == nil {
		return nil, fmt.Errorf("classInfo cannot be nil")
	}

	if config.Endpoint == "" {
		return nil, fmt.Errorf("AI model endpoint is required")
	}

	// Build the prompt for AI documentation generation
	prompt := buildPrompt(classInfo, dependencyResult)

	// Call the AI model
	response, err := callAIModel(config, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to call AI model: %w", err)
	}

	// Parse the AI response
	documentation := parseAIResponse(response, classInfo.ClassName)

	return documentation, nil
}

// buildPrompt constructs a detailed prompt for AI documentation generation.
func buildPrompt(classInfo *core.DecompiledClass, dependencyResult *core.DependencyAnalysisResult) string {
	var prompt strings.Builder

	prompt.WriteString("Please generate comprehensive documentation for the following Java class:\n\n")

	// Class basic information
	prompt.WriteString(fmt.Sprintf("Class Name: %s\n", classInfo.ClassName))
	if classInfo.PackageName != "" {
		prompt.WriteString(fmt.Sprintf("Package: %s\n", classInfo.PackageName))
	}
	prompt.WriteString("\n")

	// Class structure
	if len(classInfo.Fields) > 0 {
		prompt.WriteString("Fields:\n")
		for _, field := range classInfo.Fields {
			prompt.WriteString(fmt.Sprintf("- %s %s\n", field.Type, field.Name))
		}
		prompt.WriteString("\n")
	}

	if len(classInfo.Methods) > 0 {
		prompt.WriteString("Methods:\n")
		for _, method := range classInfo.Methods {
			prompt.WriteString(fmt.Sprintf("- %s %s(%s) -> %s\n", 
				method.Name, method.Name, strings.Join(method.Parameters, ", "), method.ReturnType))
			if method.JavaDoc != "" {
				prompt.WriteString(fmt.Sprintf("  Comments: %s\n", method.JavaDoc))
			}
		}
		prompt.WriteString("\n")
	}

	// Dependency information
	if dependencyResult != nil {
		prompt.WriteString("Dependencies:\n")
		for _, dep := range dependencyResult.ClassDependencies {
			if dep.FromClass == classInfo.ClassName {
				prompt.WriteString(fmt.Sprintf("- %s %s\n", dep.DepType, dep.ToClass))
			}
		}
		prompt.WriteString("\n")
	}

	// Source code snippet (first 500 characters)
	if classInfo.SourceCode != "" {
		sourceSnippet := classInfo.SourceCode
		if len(sourceSnippet) > 500 {
			sourceSnippet = sourceSnippet[:500] + "..."
		}
		prompt.WriteString(fmt.Sprintf("Source Code Snippet:\n```java\n%s\n```\n\n", sourceSnippet))
	}

	// Documentation requirements
	prompt.WriteString("Please provide:\n")
	prompt.WriteString("1. A clear description of what this class does\n")
	prompt.WriteString("2. The main purpose and use cases\n")
	prompt.WriteString("3. A simple usage example\n")
	prompt.WriteString("4. Key dependencies and their roles\n")
	prompt.WriteString("\nFormat the response as clear, professional documentation.")

	return prompt.String()
}

// callAIModel makes an HTTP request to the AI model API.
func callAIModel(config AIModelConfig, prompt string) (string, error) {
	// Prepare the request
	request := AIRequest{
		Model: config.Model,
		Messages: []Message{
			{
				Role:    "system",
				Content: "You are a technical documentation expert. Generate clear, comprehensive documentation for Java classes.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Temperature: 0.3, // Lower temperature for more consistent documentation
		MaxTokens:   2000,
	}

	// Marshal request to JSON
	requestBody, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", config.Endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	if config.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+config.APIKey)
	}

	// Create HTTP client with timeout
	timeout := time.Duration(config.Timeout) * time.Second
	if timeout == 0 {
		timeout = 30 * time.Second // Default timeout
	}
	client := &http.Client{Timeout: timeout}

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	// Parse response
	var aiResponse AIResponse
	if err := json.Unmarshal(respBody, &aiResponse); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	// Check for API errors
	if aiResponse.Error != nil {
		return "", fmt.Errorf("AI API error: %s", aiResponse.Error.Message)
	}

	// Extract the generated text
	if len(aiResponse.Choices) == 0 {
		return "", fmt.Errorf("no choices in AI response")
	}

	return aiResponse.Choices[0].Message.Content, nil
}

// parseAIResponse parses the AI model response into structured documentation.
func parseAIResponse(response, className string) *DocumentationResult {
	// For now, we'll store the entire response as description
	// In a more sophisticated implementation, we could parse structured responses
	return &DocumentationResult{
		ClassName:    className,
		Description:  response,
		Purpose:      extractSection(response, "purpose", "use case"),
		UsageExample: extractSection(response, "example", "usage"),
		Dependencies: extractSection(response, "dependencies", "depend"),
		GeneratedAt:  time.Now(),
	}
}

// extractSection attempts to extract a specific section from the AI response.
func extractSection(response, keyword1, keyword2 string) string {
	lines := strings.Split(response, "\n")
	var section strings.Builder
	inSection := false

	for _, line := range lines {
		lowerLine := strings.ToLower(strings.TrimSpace(line))
		
		// Check if this line starts a relevant section (header line containing keywords)
		if (strings.Contains(lowerLine, strings.ToLower(keyword1)) || strings.Contains(lowerLine, strings.ToLower(keyword2))) && 
		   (strings.HasPrefix(lowerLine, "#") || strings.HasPrefix(lowerLine, "##")) {
			inSection = true
			continue
		}
		
		// Check if we've moved to a different section (starts with ## or #)
		if inSection && (strings.HasPrefix(strings.TrimSpace(line), "##") || strings.HasPrefix(strings.TrimSpace(line), "#")) {
			break
		}
		
		// Add line to section if we're in the right section
		if inSection {
			trimmedLine := strings.TrimSpace(line)
			if trimmedLine != "" {
				if section.Len() > 0 {
					section.WriteString("\n")
				}
				section.WriteString(trimmedLine)
			}
		}
	}

	return strings.TrimSpace(section.String())
}

// ValidateConfig validates the AI model configuration.
func ValidateConfig(config AIModelConfig) error {
	if config.Endpoint == "" {
		return fmt.Errorf("endpoint is required")
	}
	if config.Model == "" {
		return fmt.Errorf("model is required")
	}
	return nil
}