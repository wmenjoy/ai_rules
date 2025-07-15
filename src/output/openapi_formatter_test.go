/**
 * [AI-ASSISTED]
 * 生成工具: Claude 4 Sonnet
 * 生成日期: 2024-12-19
 * 贡献程度: 完全生成
 * 人工修改: 无
 * 责任人: AI Assistant
 */

package output

import (
	"encoding/json"
	"strings"
	"testing"

	"ai-jar-sdk-generator/src/ai"
	"ai-jar-sdk-generator/src/core"
)

func TestFormatToOpenAPI(t *testing.T) {
	// Create test data
	classInfo := &core.DecompiledClass{
		ClassName:   "UserController",
		PackageName: "com.example.controller",
		Methods: []core.DecompiledMethod{
			{
				Name:       "getUser",
				ReturnType: "User",
				Parameters: []string{"Long id"},
				Modifiers:  []string{"public"},
				Annotations: map[string]string{
					"GetMapping": `value = "/users/{id}"`},
			},
			{
				Name:       "createUser",
				ReturnType: "User",
				Parameters: []string{"User user"},
				Modifiers:  []string{"public"},
				Annotations: map[string]string{
					"PostMapping": `value = "/users"`},
			},
		},
		Fields: []core.DecompiledField{
			{
				Name:      "userService",
				Type:      "UserService",
				Modifiers: []string{"private"},
			},
		},
		Modifiers: []string{"public"},
	}

	methodDocs := map[string]*ai.MethodDocumentation{
		"getUser": {
			MethodName:  "getUser",
			Description: "Retrieves a user by ID",
			Parameters: []ai.ParamDoc{
				{
					Name:        "id",
					Type:        "Long",
					Description: "The user ID",
					Required:    true,
				},
			},
			ReturnValue: "User object",
		},
		"createUser": {
			MethodName:  "createUser",
			Description: "Creates a new user",
			Parameters: []ai.ParamDoc{
				{
					Name:        "user",
					Type:        "User",
					Description: "The user object to create",
					Required:    true,
				},
			},
			ReturnValue: "Created user object",
		},
	}

	config := OpenAPIConfig{
		Title:       "User API",
		Description: "API for user management",
		Version:     "1.0.0",
		ServerURL:   "http://localhost:8080",
	}

	// Test FormatToOpenAPI
	openAPIObj, err := FormatToOpenAPI(classInfo, methodDocs, config)
	if err != nil {
		t.Fatalf("FormatToOpenAPI failed: %v", err)
	}

	// Verify basic structure
	if openAPIObj.OpenAPI != "3.0.0" {
		t.Errorf("Expected OpenAPI version 3.0.0, got %s", openAPIObj.OpenAPI)
	}

	if openAPIObj.Info.Title != "User API" {
		t.Errorf("Expected title 'User API', got %s", openAPIObj.Info.Title)
	}

	if openAPIObj.Info.Version != "1.0.0" {
		t.Errorf("Expected version '1.0.0', got %s", openAPIObj.Info.Version)
	}

	// Verify server
	if len(openAPIObj.Servers) != 1 {
		t.Errorf("Expected 1 server, got %d", len(openAPIObj.Servers))
	} else if openAPIObj.Servers[0].URL != "http://localhost:8080" {
		t.Errorf("Expected server URL 'http://localhost:8080', got %s", openAPIObj.Servers[0].URL)
	}

	// Verify paths
	if len(openAPIObj.Paths) == 0 {
		t.Error("Expected paths to be generated")
	}

	// Check specific paths
	if getUserPath, exists := openAPIObj.Paths["/users/{id}"]; exists {
		if getUserPath.Get == nil {
			t.Error("Expected GET operation for /users/{id}")
		} else {
			if getUserPath.Get.Summary != "Retrieves a user by ID" {
				t.Errorf("Expected summary 'Retrieves a user by ID', got %s", getUserPath.Get.Summary)
			}
		}
	} else {
		t.Error("Expected path /users/{id} to exist")
	}

	if createUserPath, exists := openAPIObj.Paths["/users"]; exists {
		if createUserPath.Post == nil {
			t.Error("Expected POST operation for /users")
		} else {
			if createUserPath.Post.Summary != "Creates a new user" {
				t.Errorf("Expected summary 'Creates a new user', got %s", createUserPath.Post.Summary)
			}
		}
	} else {
		t.Error("Expected path /users to exist")
	}

	// Verify components
	if openAPIObj.Components == nil {
		t.Error("Expected components to be generated")
	} else if openAPIObj.Components.Schemas == nil {
		t.Error("Expected schemas to be generated")
	} else if _, exists := openAPIObj.Components.Schemas["UserController"]; !exists {
		t.Error("Expected UserController schema to exist")
	}
}

func TestFormatToOpenAPIJSON(t *testing.T) {
	// Create a simple OpenAPI object
	openAPIObj := &OpenAPIObject{
		OpenAPI: "3.0.0",
		Info: OpenAPIInfo{
			Title:   "Test API",
			Version: "1.0.0",
		},
		Paths: map[string]PathItem{
			"/test": {
				Get: &Operation{
					Summary: "Test operation",
					Responses: map[string]Response{
						"200": {
							Description: "Success",
						},
					},
				},
			},
		},
	}

	// Test JSON formatting
	jsonStr, err := FormatToOpenAPIJSON(openAPIObj)
	if err != nil {
		t.Fatalf("FormatToOpenAPIJSON failed: %v", err)
	}

	// Verify it's valid JSON
	var result map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		t.Fatalf("Generated JSON is invalid: %v", err)
	}

	// Verify basic structure
	if result["openapi"] != "3.0.0" {
		t.Errorf("Expected openapi version 3.0.0 in JSON")
	}

	if info, ok := result["info"].(map[string]interface{}); ok {
		if info["title"] != "Test API" {
			t.Errorf("Expected title 'Test API' in JSON")
		}
	} else {
		t.Error("Expected info object in JSON")
	}
}

func TestExtractRESTPathInfo(t *testing.T) {
	tests := []struct {
		name           string
		method         core.DecompiledMethod
		expectedPath   string
		expectedMethod string
		expectNil      bool
	}{
		{
			name: "GetMapping annotation",
			method: core.DecompiledMethod{
				Name: "getUser",
				Annotations: map[string]string{
					"GetMapping": `value = "/users/{id}"`},
			},
			expectedPath:   "/users/{id}",
			expectedMethod: "get",
			expectNil:      false,
		},
		{
			name: "PostMapping annotation",
			method: core.DecompiledMethod{
				Name: "createUser",
				Annotations: map[string]string{
					"PostMapping": `"/users"`},
			},
			expectedPath:   "/users",
			expectedMethod: "post",
			expectNil:      false,
		},
		{
			name: "JAX-RS Path with GET",
			method: core.DecompiledMethod{
				Name: "getUser",
				Annotations: map[string]string{
					"Path": `"/users/{id}"`,
					"GET":  "",
				},
			},
			expectedPath:   "/users/{id}",
			expectedMethod: "get",
			expectNil:      false,
		},
		{
			name: "No REST annotations",
			method: core.DecompiledMethod{
				Name:        "regularMethod",
				Annotations: map[string]string{},
			},
			expectNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractRESTPathInfo(tt.method)

			if tt.expectNil {
				if result != nil {
					t.Errorf("Expected nil result, got %+v", result)
				}
				return
			}

			if result == nil {
				t.Errorf("Expected non-nil result")
				return
			}

			if result.Path != tt.expectedPath {
				t.Errorf("Expected path %s, got %s", tt.expectedPath, result.Path)
			}

			if result.HTTPMethod != tt.expectedMethod {
				t.Errorf("Expected method %s, got %s", tt.expectedMethod, result.HTTPMethod)
			}
		})
	}
}

func TestJavaTypeToSchema(t *testing.T) {
	tests := []struct {
		javaType     string
		expectedType string
		expectedFormat string
	}{
		{"String", "string", ""},
		{"java.lang.String", "string", ""},
		{"int", "integer", "int32"},
		{"Integer", "integer", "int32"},
		{"long", "integer", "int64"},
		{"double", "number", "double"},
		{"boolean", "boolean", ""},
		{"Date", "string", "date-time"},
		{"List<String>", "array", ""},
		{"CustomObject", "object", ""},
	}

	for _, tt := range tests {
		t.Run(tt.javaType, func(t *testing.T) {
			schema := javaTypeToSchema(tt.javaType)

			if schema.Type != tt.expectedType {
				t.Errorf("Expected type %s, got %s", tt.expectedType, schema.Type)
			}

			if schema.Format != tt.expectedFormat {
				t.Errorf("Expected format %s, got %s", tt.expectedFormat, schema.Format)
			}
		})
	}
}

func TestFormatToOpenAPIWithNilInput(t *testing.T) {
	config := OpenAPIConfig{
		Title:   "Test API",
		Version: "1.0.0",
	}

	// Test with nil classInfo
	_, err := FormatToOpenAPI(nil, nil, config)
	if err == nil {
		t.Error("Expected error for nil classInfo")
	} else if !strings.Contains(err.Error(), "classInfo cannot be nil") {
		t.Errorf("Expected 'classInfo cannot be nil' error, got: %v", err)
	}
}

func TestOpenAPIConfigDefaults(t *testing.T) {
	classInfo := &core.DecompiledClass{
		ClassName: "TestClass",
		Methods:   []core.DecompiledMethod{},
		Fields:    []core.DecompiledField{},
	}

	// Test with empty config
	config := OpenAPIConfig{}
	openAPIObj, err := FormatToOpenAPI(classInfo, nil, config)
	if err != nil {
		t.Fatalf("FormatToOpenAPI failed: %v", err)
	}

	// Verify defaults
	expectedTitle := "TestClass API"
	if openAPIObj.Info.Title != expectedTitle {
		t.Errorf("Expected default title %s, got %s", expectedTitle, openAPIObj.Info.Title)
	}

	expectedVersion := "1.0.0"
	if openAPIObj.Info.Version != expectedVersion {
		t.Errorf("Expected default version %s, got %s", expectedVersion, openAPIObj.Info.Version)
	}

	// Should have no servers when ServerURL is empty
	if len(openAPIObj.Servers) != 0 {
		t.Errorf("Expected no servers, got %d", len(openAPIObj.Servers))
	}
}