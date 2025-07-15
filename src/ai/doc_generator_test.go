package ai

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"ai-jar-sdk-generator/src/core"
)

// TestValidateConfig tests the configuration validation function.
func TestValidateConfig(t *testing.T) {
	tests := []struct {
		name    string
		config  AIModelConfig
		wantErr bool
	}{
		{
			name: "valid config",
			config: AIModelConfig{
				Endpoint: "https://api.example.com/v1/chat/completions",
				APIKey:   "test-key",
				Model:    "gpt-3.5-turbo",
				Timeout:  30,
			},
			wantErr: false,
		},
		{
			name: "missing endpoint",
			config: AIModelConfig{
				APIKey: "test-key",
				Model:  "gpt-3.5-turbo",
			},
			wantErr: true,
		},
		{
			name: "missing model",
			config: AIModelConfig{
				Endpoint: "https://api.example.com/v1/chat/completions",
				APIKey:   "test-key",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateConfig(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestBuildPrompt tests the prompt building function.
func TestBuildPrompt(t *testing.T) {
	classInfo := &core.DecompiledClass{
		ClassName:   "TestClass",
		PackageName: "com.example",
		Fields: []core.DecompiledField{
			{Name: "name", Type: "String"},
			{Name: "age", Type: "int"},
		},
		Methods: []core.DecompiledMethod{
			{
				Name:       "getName",
				Parameters: []string{},
				ReturnType: "String",
				JavaDoc:    "Returns the name",
			},
			{
				Name:       "setAge",
				Parameters: []string{"int age"},
				ReturnType: "void",
			},
		},
		SourceCode: "public class TestClass { private String name; private int age; }",
	}

	dependencyResult := &core.DependencyAnalysisResult{
		ClassDependencies: []core.ClassDependency{
			{FromClass: "TestClass", ToClass: "String", DepType: "uses_field"},
		},
	}

	prompt := buildPrompt(classInfo, dependencyResult)

	// Check that the prompt contains expected elements
	expectedElements := []string{
		"TestClass",
		"com.example",
		"String name",
		"int age",
		"getName",
		"setAge",
		"Returns the name",
		"uses_field String",
		"Source Code Snippet",
	}

	for _, element := range expectedElements {
		if !strings.Contains(prompt, element) {
			t.Errorf("Expected prompt to contain '%s', but it didn't", element)
		}
	}
}

// TestParseAIResponse tests the AI response parsing function.
func TestParseAIResponse(t *testing.T) {
	response := `# TestClass Documentation

## Purpose
This class represents a person with name and age.

## Usage Example
TestClass person = new TestClass();
person.setAge(25);
String name = person.getName();

## Dependencies
This class depends on String for the name field.`

	result := parseAIResponse(response, "TestClass")

	if result.ClassName != "TestClass" {
		t.Errorf("Expected className to be 'TestClass', got '%s'", result.ClassName)
	}

	if result.Description != response {
		t.Errorf("Expected description to be the full response")
	}

	if !strings.Contains(result.Purpose, "person") {
		t.Errorf("Expected purpose to contain 'person', got '%s'", result.Purpose)
	}

	if !strings.Contains(result.UsageExample, "TestClass person") {
		t.Errorf("Expected usage example to contain 'TestClass person', got '%s'", result.UsageExample)
	}

	if !strings.Contains(result.Dependencies, "String") {
		t.Errorf("Expected dependencies to contain 'String', got '%s'", result.Dependencies)
	}
}

// TestGenerateAIDocumentation tests the main documentation generation function with a mock server.
func TestGenerateAIDocumentation(t *testing.T) {
	// Create a mock AI API server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method and headers
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
		}

		// Mock response
		response := AIResponse{
			Choices: []Choice{
				{
					Message: Message{
						Role:    "assistant",
						Content: "# TestClass\n\nThis is a test class for demonstration purposes.\n\n## Purpose\nUsed for testing.\n\n## Usage Example\n```java\nTestClass obj = new TestClass();\n```",
					},
					FinishReason: "stop",
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer mockServer.Close()

	// Test data
	classInfo := &core.DecompiledClass{
		ClassName:   "TestClass",
		PackageName: "com.example",
		SourceCode:  "public class TestClass {}",
	}

	config := AIModelConfig{
		Endpoint: mockServer.URL,
		APIKey:   "test-key",
		Model:    "test-model",
		Timeout:  10,
	}

	// Test successful generation
	result, err := GenerateAIDocumentation(classInfo, nil, config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.ClassName != "TestClass" {
		t.Errorf("Expected className to be 'TestClass', got '%s'", result.ClassName)
	}

	if !strings.Contains(result.Description, "TestClass") {
		t.Errorf("Expected description to contain 'TestClass'")
	}

	if result.GeneratedAt.IsZero() {
		t.Errorf("Expected GeneratedAt to be set")
	}
}

// TestGenerateAIDocumentationErrors tests error cases.
func TestGenerateAIDocumentationErrors(t *testing.T) {
	tests := []struct {
		name      string
		classInfo *core.DecompiledClass
		config    AIModelConfig
		wantErr   bool
	}{
		{
			name:      "nil class info",
			classInfo: nil,
			config:    AIModelConfig{Endpoint: "http://example.com", Model: "test"},
			wantErr:   true,
		},
		{
			name:      "empty endpoint",
			classInfo: &core.DecompiledClass{ClassName: "Test"},
			config:    AIModelConfig{Model: "test"},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenerateAIDocumentation(tt.classInfo, nil, tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateAIDocumentation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestExtractSection tests the section extraction function.
func TestExtractSection(t *testing.T) {
	response := `# Documentation

## Purpose
This is the purpose section.
It has multiple lines.

## Usage Example
Here's how to use it:
code example

## Dependencies
List of dependencies`

	tests := []struct {
		name     string
		keyword1 string
		keyword2 string
		expected string
	}{
		{
			name:     "extract purpose",
			keyword1: "purpose",
			keyword2: "use case",
			expected: "This is the purpose section.\nIt has multiple lines.",
		},
		{
			name:     "extract usage",
			keyword1: "example",
			keyword2: "usage",
			expected: "Here's how to use it:\ncode example",
		},
		{
			name:     "extract dependencies",
			keyword1: "dependencies",
			keyword2: "depend",
			expected: "List of dependencies",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractSection(response, tt.keyword1, tt.keyword2)
			if result != tt.expected {
				t.Errorf("extractSection() = %q, want %q", result, tt.expected)
			}
		})
	}
}

// TestCallAIModelWithErrorResponse tests error handling in AI API calls.
func TestCallAIModelWithErrorResponse(t *testing.T) {
	// Create a mock server that returns an error
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		response := AIResponse{
			Error: &APIError{
				Message: "Invalid request",
				Type:    "invalid_request_error",
				Code:    "400",
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer mockServer.Close()

	config := AIModelConfig{
		Endpoint: mockServer.URL,
		Model:    "test-model",
		Timeout:  5,
	}

	_, err := callAIModel(config, "test prompt")
	if err == nil {
		t.Error("Expected error from API call, got nil")
	}

	if !strings.Contains(err.Error(), "400") {
		t.Errorf("Expected error to contain status code, got %v", err)
	}
}

// BenchmarkBuildPrompt benchmarks the prompt building function.
func BenchmarkBuildPrompt(b *testing.B) {
	classInfo := &core.DecompiledClass{
		ClassName:   "BenchmarkClass",
		PackageName: "com.benchmark",
		Fields: []core.DecompiledField{
			{Name: "field1", Type: "String"},
			{Name: "field2", Type: "int"},
			{Name: "field3", Type: "List<String>"},
		},
		Methods: []core.DecompiledMethod{
			{Name: "method1", Parameters: []string{"String arg"}, ReturnType: "void"},
			{Name: "method2", Parameters: []string{}, ReturnType: "String"},
		},
		SourceCode: strings.Repeat("public class BenchmarkClass { /* code */ }", 10),
	}

	dependencyResult := &core.DependencyAnalysisResult{
		ClassDependencies: []core.ClassDependency{
			{FromClass: "BenchmarkClass", ToClass: "String", DepType: "uses_field"},
			{FromClass: "BenchmarkClass", ToClass: "List", DepType: "uses_field"},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = buildPrompt(classInfo, dependencyResult)
	}
}