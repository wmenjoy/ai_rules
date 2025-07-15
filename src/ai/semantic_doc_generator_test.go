/**
 * [AI-ASSISTED]
 * 生成工具: Claude 4 Sonnet
 * 生成日期: 2024-12-19
 * 贡献程度: 完全生成
 * 人工修改: 无
 * 责任人: AI Assistant
 */

package ai

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"ai-jar-sdk-generator/src/core"
)

func TestGenerateSemanticMethodDoc(t *testing.T) {
	// Create mock AI server
	mockResponse := "## Method Description\nThis method calculates the sum of two integers and returns the result.\n\n## Parameters\n- a (int): The first integer to add\n- b (int): The second integer to add\n\n## Return Value\nReturns an integer representing the sum of the two input parameters.\n\n## Usage Example\n```java\nCalculator calc = new Calculator();\nint result = calc.add(5, 3); // Returns 8\n```\n\n## Exception Handling\nThis method does not throw any exceptions as it performs basic arithmetic.\n\n## Complexity\nTime complexity: O(1) - constant time operation\nSpace complexity: O(1) - no additional space required\n\n## Best Practices\nEnsure input values are within the valid range for integer operations to avoid overflow."

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{
			"choices": []map[string]interface{}{
				{
					"message": map[string]interface{}{
						"content": mockResponse,
					},
				},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Test data
	methodInfo := &core.DecompiledMethod{
		Name:       "add",
		ReturnType: "int",
		Parameters: []string{"int a", "int b"},
		Modifiers:  []string{"public"},
		SourceCode: "public int add(int a, int b) { return a + b; }",
		JavaDoc:    "Adds two integers",
	}

	classContext := &core.DecompiledClass{
		ClassName:   "Calculator",
		PackageName: "com.example",
		SuperClass:  "Object",
		Interfaces:  []string{},
		Methods:     []core.DecompiledMethod{*methodInfo},
	}

	config := AIModelConfig{
		Endpoint: server.URL,
		APIKey:   "test-key",
		Model:    "test-model",
		Timeout:  30,
	}

	// Test successful generation
	doc, err := GenerateSemanticMethodDoc(methodInfo, classContext, nil, config)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if doc == nil {
		t.Fatal("Expected documentation, got nil")
	}

	if doc.MethodName != "add" {
		t.Errorf("Expected method name 'add', got: %s", doc.MethodName)
	}

	if doc.Description == "" {
		t.Error("Expected description to be extracted")
	}

	if doc.ReturnValue == "" {
		t.Error("Expected return value to be extracted")
	}

	if doc.UsageExample == "" {
		t.Error("Expected usage example to be extracted")
	}

	if len(doc.Parameters) == 0 {
		t.Error("Expected parameters to be extracted")
	}

	if doc.GeneratedAt.IsZero() {
		t.Error("Expected GeneratedAt to be set")
	}
}

func TestGenerateSemanticMethodDocWithNilInputs(t *testing.T) {
	config := AIModelConfig{
		Endpoint: "http://localhost:8080",
		APIKey:   "test-key",
		Model:    "test-model",
		Timeout:  30,
	}

	// Test with nil method info
	_, err := GenerateSemanticMethodDoc(nil, &core.DecompiledClass{}, nil, config)
	if err == nil {
		t.Error("Expected error for nil method info")
	}

	// Test with nil class context
	_, err = GenerateSemanticMethodDoc(&core.DecompiledMethod{}, nil, nil, config)
	if err == nil {
		t.Error("Expected error for nil class context")
	}
}

func TestGenerateSemanticClassDoc(t *testing.T) {
	// Create mock AI server
	mockResponse := "## Method Description\nThis is a test method.\n\n## Parameters\n- param1 (String): Test parameter\n\n## Return Value\nReturns void.\n\n## Usage Example\nobj.testMethod(\"test\");\n\n## Exception Handling\nNo exceptions thrown.\n\n## Complexity\nO(1)\n\n## Best Practices\nUse meaningful parameter names."

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{
			"choices": []map[string]interface{}{
				{
					"message": map[string]interface{}{
						"content": mockResponse,
					},
				},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	classInfo := &core.DecompiledClass{
		ClassName:   "TestClass",
		PackageName: "com.example",
		Methods: []core.DecompiledMethod{
			{
				Name:       "testMethod",
				ReturnType: "void",
				Parameters: []string{"String param1"},
			},
			{
				Name:       "anotherMethod",
				ReturnType: "int",
				Parameters: []string{},
			},
		},
	}

	config := AIModelConfig{
		Endpoint: server.URL,
		APIKey:   "test-key",
		Model:    "test-model",
		Timeout:  30,
	}

	result, err := GenerateSemanticClassDoc(classInfo, nil, config)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	if result.ClassName != "TestClass" {
		t.Errorf("Expected class name 'TestClass', got: %s", result.ClassName)
	}

	if result.PackageName != "com.example" {
		t.Errorf("Expected package name 'com.example', got: %s", result.PackageName)
	}

	if result.TotalMethods != 2 {
		t.Errorf("Expected 2 total methods, got: %d", result.TotalMethods)
	}

	if result.DocumentedMethods != 2 {
		t.Errorf("Expected 2 documented methods, got: %d", result.DocumentedMethods)
	}

	if len(result.MethodDocs) != 2 {
		t.Errorf("Expected 2 method docs, got: %d", len(result.MethodDocs))
	}
}

func TestBuildMethodPrompt(t *testing.T) {
	methodInfo := &core.DecompiledMethod{
		Name:       "calculateSum",
		ReturnType: "int",
		Parameters: []string{"int a", "int b"},
		Modifiers:  []string{"public", "static"},
		SourceCode: "public static int calculateSum(int a, int b) { return a + b; }",
		JavaDoc:    "Calculates the sum of two integers",
	}

	classContext := &core.DecompiledClass{
		ClassName:   "MathUtils",
		PackageName: "com.example.utils",
		SuperClass:  "Object",
		Interfaces:  []string{"Serializable"},
		Fields:      []core.DecompiledField{{Name: "PI", Type: "double"}},
		Methods:     []core.DecompiledMethod{*methodInfo},
	}

	dependencyResult := &core.DependencyAnalysisResult{
		ClassDependencies: []core.ClassDependency{
			{
				FromClass: "MathUtils",
				ToClass:   "Math",
				DepType:   "import",
			},
		},
	}

	prompt := buildMethodPrompt(methodInfo, classContext, dependencyResult)

	// Check that prompt contains expected information
	expectedContents := []string{
		"MathUtils",
		"calculateSum",
		"int a, int b",
		"public, static",
		"Calculates the sum of two integers",
		"public static int calculateSum",
		"com.example.utils",
		"Serializable",
		"Math",
		"## Method Description",
		"## Parameters",
		"## Return Value",
		"## Usage Example",
		"## Exception Handling",
		"## Complexity",
		"## Best Practices",
	}

	for _, expected := range expectedContents {
		if !strings.Contains(prompt, expected) {
			t.Errorf("Expected prompt to contain '%s', but it didn't", expected)
		}
	}
}

func TestParseMethodAIResponse(t *testing.T) {
	response := "## Method Description\nThis method performs a complex calculation.\n\n## Parameters\n- input (String): The input value to process\n- options (Map): Configuration options (optional)\n\n## Return Value\nReturns a ProcessedResult object containing the calculation results.\n\n## Usage Example\n```java\nProcessor proc = new Processor();\nMap<String, Object> opts = new HashMap<>();\nProcessedResult result = proc.process(\"data\", opts);\n```\n\n## Exception Handling\nThrows ProcessingException if input is invalid.\nThrows IllegalArgumentException if options are malformed.\n\n## Complexity\nTime complexity: O(n log n) where n is the input size\nSpace complexity: O(n)\n\n## Best Practices\nValidate input parameters before processing.\nUse appropriate error handling for production code."

	methodDoc := parseMethodAIResponse(response, "process")

	if methodDoc.MethodName != "process" {
		t.Errorf("Expected method name 'process', got: %s", methodDoc.MethodName)
	}

	if !strings.Contains(methodDoc.Description, "complex calculation") {
		t.Error("Expected description to be extracted")
	}

	if !strings.Contains(methodDoc.ReturnValue, "ProcessedResult") {
		t.Error("Expected return value to be extracted")
	}

	if !strings.Contains(methodDoc.UsageExample, "Processor") {
		t.Error("Expected usage example to be extracted")
	}

	if !strings.Contains(methodDoc.ExceptionHandling, "ProcessingException") {
		t.Error("Expected exception handling to be extracted")
	}

	if !strings.Contains(methodDoc.Complexity, "O(n log n)") {
		t.Error("Expected complexity to be extracted")
	}

	if !strings.Contains(methodDoc.BestPractices, "Validate input") {
		t.Error("Expected best practices to be extracted")
	}

	if len(methodDoc.Parameters) != 2 {
		t.Errorf("Expected 2 parameters, got: %d", len(methodDoc.Parameters))
	}

	// Check first parameter
	if len(methodDoc.Parameters) > 0 {
		param := methodDoc.Parameters[0]
		if param.Name != "input" {
			t.Errorf("Expected first parameter name 'input', got: %s", param.Name)
		}
		if param.Type != "String" {
			t.Errorf("Expected first parameter type 'String', got: %s", param.Type)
		}
		if !param.Required {
			t.Error("Expected first parameter to be required")
		}
	}

	// Check second parameter
	if len(methodDoc.Parameters) > 1 {
		param := methodDoc.Parameters[1]
		if param.Name != "options" {
			t.Errorf("Expected second parameter name 'options', got: %s", param.Name)
		}
		if param.Type != "Map" {
			t.Errorf("Expected second parameter type 'Map', got: %s", param.Type)
		}
		if param.Required {
			t.Error("Expected second parameter to be optional")
		}
	}
}

func TestExtractMethodSection(t *testing.T) {
	response := "## Method Description\nThis is the description section.\nIt spans multiple lines.\n\n## Parameters\n- param1: First parameter\n- param2: Second parameter\n\n## Return Value\nReturns a string value.\n\n## Usage Example\nSample usage code here."

	tests := []struct {
		name      string
		keyword1  string
		keyword2  string
		expected  string
	}{
		{
			name:     "extract_description",
			keyword1: "description",
			keyword2: "method description",
			expected: "This is the description section.\nIt spans multiple lines.",
		},
		{
			name:     "extract_parameters",
			keyword1: "parameters",
			keyword2: "parameter",
			expected: "- param1: First parameter\n- param2: Second parameter",
		},
		{
			name:     "extract_return",
			keyword1: "return",
			keyword2: "return value",
			expected: "Returns a string value.",
		},
		{
			name:     "extract_usage",
			keyword1: "example",
			keyword2: "usage",
			expected: "Sample usage code here.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractMethodSection(response, tt.keyword1, tt.keyword2)
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestParseParametersFromText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected []ParamDoc
	}{
		{
			name: "standard_format",
			text: "- input (String): The input value\n- count (int): Number of items",
			expected: []ParamDoc{
				{Name: "input", Type: "String", Description: "The input value", Required: true},
				{Name: "count", Type: "int", Description: "Number of items", Required: true},
			},
		},
		{
			name: "optional_parameter",
			text: "- config (Map): Configuration options (optional)",
			expected: []ParamDoc{
				{Name: "config", Type: "Map", Description: "Configuration options (optional)", Required: false},
			},
		},
		{
			name: "type_name_format",
			text: "- String name: The name parameter\n- int age: The age parameter",
			expected: []ParamDoc{
				{Name: "name", Type: "String", Description: "The name parameter", Required: true},
				{Name: "age", Type: "int", Description: "The age parameter", Required: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseParametersFromText(tt.text)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d parameters, got %d", len(tt.expected), len(result))
				return
			}

			for i, expected := range tt.expected {
				if i >= len(result) {
					break
				}
				actual := result[i]
				if actual.Name != expected.Name {
					t.Errorf("Parameter %d: expected name '%s', got '%s'", i, expected.Name, actual.Name)
				}
				if actual.Type != expected.Type {
					t.Errorf("Parameter %d: expected type '%s', got '%s'", i, expected.Type, actual.Type)
				}
				if actual.Required != expected.Required {
					t.Errorf("Parameter %d: expected required %v, got %v", i, expected.Required, actual.Required)
				}
			}
		})
	}
}

func TestExportSemanticDocToJSON(t *testing.T) {
	result := &SemanticDocumentationResult{
		ClassName:         "TestClass",
		PackageName:       "com.example",
		GeneratedAt:       time.Now(),
		TotalMethods:      1,
		DocumentedMethods: 1,
		MethodDocs: []MethodDocumentation{
			{
				MethodName:  "testMethod",
				Description: "Test method",
				Parameters: []ParamDoc{
					{Name: "param1", Type: "String", Description: "Test parameter", Required: true},
				},
				ReturnValue:  "void",
				GeneratedAt:  time.Now(),
			},
		},
	}

	jsonData, err := ExportSemanticDocToJSON(result)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(jsonData) == 0 {
		t.Error("Expected JSON data, got empty")
	}

	// Verify it's valid JSON
	var parsed map[string]interface{}
	err = json.Unmarshal(jsonData, &parsed)
	if err != nil {
		t.Fatalf("Generated JSON is invalid: %v", err)
	}

	// Test with nil input
	_, err = ExportSemanticDocToJSON(nil)
	if err == nil {
		t.Error("Expected error for nil input")
	}
}

func TestValidateMethodDocumentation(t *testing.T) {
	// Test with nil documentation
	issues := ValidateMethodDocumentation(nil)
	if len(issues) == 0 {
		t.Error("Expected validation issues for nil documentation")
	}

	// Test with incomplete documentation
	incompleteDoc := &MethodDocumentation{
		MethodName: "test",
		// Missing other required fields
	}

	issues = ValidateMethodDocumentation(incompleteDoc)
	expectedIssues := []string{
		"missing method description",
		"missing return value description",
		"missing usage example",
		"no parameter documentation found",
	}

	for _, expectedIssue := range expectedIssues {
		found := false
		for _, issue := range issues {
			if issue == expectedIssue {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected validation issue '%s' not found", expectedIssue)
		}
	}

	// Test with complete documentation
	completeDoc := &MethodDocumentation{
		MethodName:        "test",
		Description:       "Test method",
		ReturnValue:       "void",
		UsageExample:      "test();",
		ExceptionHandling: "None",
		Complexity:        "O(1)",
		BestPractices:     "Use wisely",
		Parameters: []ParamDoc{
			{Name: "param1", Description: "Test param", Required: true},
		},
	}

	issues = ValidateMethodDocumentation(completeDoc)
	if len(issues) != 0 {
		t.Errorf("Expected no validation issues for complete documentation, got: %v", issues)
	}
}

// Benchmark for method prompt building
func BenchmarkBuildMethodPrompt(b *testing.B) {
	methodInfo := &core.DecompiledMethod{
		Name:       "complexMethod",
		ReturnType: "List<String>",
		Parameters: []string{"Map<String, Object> config", "List<Integer> values", "boolean flag"},
		Modifiers:  []string{"public", "synchronized"},
		SourceCode: "public synchronized List<String> complexMethod(Map<String, Object> config, List<Integer> values, boolean flag) { /* complex implementation */ }",
		JavaDoc:    "Complex method with multiple parameters and generic types",
	}

	classContext := &core.DecompiledClass{
		ClassName:   "ComplexProcessor",
		PackageName: "com.example.complex",
		SuperClass:  "AbstractProcessor",
		Interfaces:  []string{"Processor", "Configurable", "Serializable"},
		Fields:      make([]core.DecompiledField, 10),
		Methods:     make([]core.DecompiledMethod, 20),
	}

	dependencyResult := &core.DependencyAnalysisResult{
		ClassDependencies: make([]core.ClassDependency, 50),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = buildMethodPrompt(methodInfo, classContext, dependencyResult)
	}
}